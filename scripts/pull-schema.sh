#!/usr/bin/env bash

# https://stackoverflow.com/a/246128
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# Pull the schema
SCHEMA="$(curl -s -L https://raw.githubusercontent.com/grafana/grafana/${GRAFANA_TARGET_VERSION}/public/api-merged.json)"

# Custom extensions: https://goswagger.io/use/models/schemas.html#custom-extensions
# These may have to be updated for future versions of Grafana
modify() {
    SCHEMA="$(echo "${SCHEMA}" | jq "${1}")"
}

# Playlist models are all messed up
# TODO: Upstream fix
modify 'del(.definitions.Item)' # Old playlist item schema, PlaylistItem is more up to date
modify '.responses.getPlaylistItemsResponse.schema.items["$ref"] = "#/definitions/PlaylistItem"' # Currently pointing to Item (old PlaylistItem model)
modify '.responses.updatePlaylistResponse.schema["$ref"] = "#/definitions/Playlist"' # Currently pointing to Spec (Preferences)
modify '.responses.getPlaylistResponse.schema["$ref"] = "#/definitions/Playlist"' # Currently pointing to Spec (Preferences)

# Fixed in the grafana repo here: https://github.com/grafana/grafana/pull/77605
modify '.definitions.ItemDTO["x-go-name"] = "Annotation"'

# Currently string, should be integer
# Fixed in the grafana repo here: https://github.com/grafana/grafana/pull/76749
modify '.responses.postDashboardResponse.schema.properties.id.type = "integer"'
modify '.responses.postDashboardResponse.schema.properties.id.format = "int64"'

# Currently set as a single TokenDTO, not an array
# Fixed in the grafana repo here: https://github.com/grafana/grafana/pull/78155
modify '.responses.listTokensResponse.schema = {
    "type": "array", 
    "items": {
        "$ref": "#/definitions/TokenDTO" 
    }
}'

# This endpoint should return an array of LibraryElementDTOs, not a single LibraryElementDTO
# Fixed in the grafana repo here: https://github.com/grafana/grafana/pull/78221
modify '.paths["/library-elements/name/{library_element_name}"].get.responses["200"]["$ref"] = "#/responses/getLibraryElementArrayResponse"'
modify '.definitions.LibraryElementArrayResponse = {
    "type": "object",
    "title": "LibraryElementArrayResponse is a response struct for an array of LibraryElementDTO.",
    "properties": {
        "result": {
            "type": "array",
            "items": {
                "$ref": "#/definitions/LibraryElementDTO"
            }
        }
    }
}'
modify '.responses.getLibraryElementArrayResponse = {
    "description": "(empty)",
    "schema": {
        "$ref": "#/definitions/LibraryElementArrayResponse"
    }
}'

# Fixed in the grafana repo here: https://github.com/grafana/grafana/pull/78226
modify '.definitions.Spec["x-go-name"] = "Preferences"'

# Fixed in the grafana repo here: https://github.com/grafana/grafana/pull/78491
modify '.definitions.AddCommand["x-go-name"] = "AddAPIKeyCommand"'

# Any endpoint that starts with /api/ should be trimmed because it's redundant (API path is configured on the client), ex: /api/dashboards/ -> /dashboards/
# TODO: Upstream fix
# Move /api/ map keys to a new key without /api/ prefix
modify '.paths = (.paths | with_entries(.key |= sub("^/api"; "")))'

# Remove "Route" prefixes to operation IDs (ex: RouteGetxxx)
# TODO: Upstream fix
modify '.paths = .paths | walk(if type == "object" and has("operationId") then .operationId |= sub("^Route";"") else . end)' 

# The "for" property returned by the API is a string (can't be unmarshaled to time.Duration)
# TODO: Upstream fix
modify '.definitions.ProvisionedAlertRule.properties.for = {
    "type": "string",
    "format": "duration"
}'

# The X-Disable-Provenance header is only on the alert-rules endpoint in the swagger spec, but it also works for entire rule groups
# Fixed in the grafana repo here: https://github.com/grafana/grafana/pull/79278
disable_provenance_header='[{
    "type": "string",
    "name": "X-Disable-Provenance",
    "in": "header"
}]'
modify '.paths["/v1/provisioning/alert-rules/{UID}"].delete.parameters += '"${disable_provenance_header}"
modify '.paths["/v1/provisioning/contact-points"].post.parameters += '"${disable_provenance_header}"
modify '.paths["/v1/provisioning/contact-points/{UID}"].put.parameters += '"${disable_provenance_header}"
modify '.paths["/v1/provisioning/folder/{FolderUID}/rule-groups/{Group}"].put.parameters += '"${disable_provenance_header}"
modify '.paths["/v1/provisioning/policies"].put.parameters += '"${disable_provenance_header}"
modify '.paths["/v1/provisioning/mute-timings"].post.parameters += '"${disable_provenance_header}"
modify '.paths["/v1/provisioning/mute-timings/{name}"].put.parameters += '"${disable_provenance_header}"
modify '.paths["/v1/provisioning/templates/{name}"].put.parameters += '"${disable_provenance_header}"

# The global property is not in the RoleDTO model, it is added in the MarshalJSON method
# As a result, it is not found by go-swagger, but it's still useful
# TODO: Upstream fix
modify '.definitions.RoleDTO.properties.global = {
    "type": "boolean",
    "description": "Whether the role is global or not."
}'

# StartDate and EndDate of reports must be nullable
modify '.definitions.ScheduleDTO.properties.startDate["x-nullable"] = true'
modify '.definitions.ScheduleDTO.properties.endDate["x-nullable"] = true'

# Write the schema to a file
echo "${SCHEMA}" > "${SCRIPT_DIR}/schema.json"
