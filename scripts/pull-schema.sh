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

modify '.definitions.ItemDTO["x-go-name"] = "Annotation"'
modify '.definitions.Spec["x-go-name"] = "Preferences"'
modify 'del(.definitions.Item)' # Old playlist item schema, PlaylistItem is more up to date
modify '.responses.getPlaylistItemsResponse.schema.items["$ref"] = "#/definitions/PlaylistItem"' # Currently pointing to Item (old PlaylistItem model)
modify '.responses.updatePlaylistResponse.schema["$ref"] = "#/definitions/Playlist"' # Currently pointing to Spec (Preferences)
modify '.responses.getPlaylistResponse.schema["$ref"] = "#/definitions/Playlist"' # Currently pointing to Spec (Preferences)
modify '.paths = .paths | walk(if type == "object" and has("operationId") then .operationId |= sub("^Route";"") else . end)' # Remove "Route" prefixes to operation IDs (ex: RouteGetxxx)
modify '.responses.postDashboardResponse.schema.properties.id.type = "integer"' # Currently string, should be integer
modify '.responses.postDashboardResponse.schema.properties.id.format = "int64"'
modify '.responses.listTokensResponse.schema = {"type": "array", "items": {"$ref": "#/definitions/TokenDTO" }}' # Currently set as a single TokenDTO, not an array

# Write the schema to a file
echo "${SCHEMA}" > "${SCRIPT_DIR}/schema.json"
