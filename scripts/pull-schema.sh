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

modify '.definitions.Spec["x-go-name"] = "Preferences"'
modify 'del(.definitions.Item)' # Old playlist item schema, PlaylistItem is more up to date
modify '.responses.getPlaylistItemsResponse.schema.items["$ref"] = "#/definitions/PlaylistItem"' # Currently pointing to Item (old PlaylistItem model)
modify '.responses.updatePlaylistResponse.schema["$ref"] = "#/definitions/Playlist"' # Currently pointing to Spec (Preferences)
modify '.responses.getPlaylistResponse.schema["$ref"] = "#/definitions/Playlist"' # Currently pointing to Spec (Preferences)
modify '.paths = .paths | walk(if type == "object" and has("operationId") then .operationId |= sub("^Route";"") else . end)' # Remove "Route" prefixes to operation IDs (ex: RouteGetxxx)

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

# Write the schema to a file
echo "${SCHEMA}" > "${SCRIPT_DIR}/schema.json"
