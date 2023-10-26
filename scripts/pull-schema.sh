#!/usr/bin/env bash

# https://stackoverflow.com/a/246128
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# Pull the schema
SCHEMA="$(curl -s -L https://raw.githubusercontent.com/grafana/grafana/${GRAFANA_TARGET_VERSION}/public/api-merged.json)"

# Custom extensions: https://goswagger.io/use/models/schemas.html#custom-extensions
# These may have to be updated for future versions of Grafana
SCHEMA="$(echo "${SCHEMA}" | jq '.definitions.ItemDTO["x-go-name"] = "Annotation"')"
SCHEMA="$(echo "${SCHEMA}" | jq '.definitions.Spec["x-go-name"] = "Preferences"')"
SCHEMA="$(echo "${SCHEMA}" | jq 'del(.definitions.Item)')" # Old playlist item schema, PlaylistItem is more up to date
SCHEMA="$(echo "${SCHEMA}" | jq '.responses.getPlaylistItemsResponse.schema.items["$ref"] = "#/definitions/PlaylistItem"')" # Currently pointing to Item (old PlaylistItem model)
SCHEMA="$(echo "${SCHEMA}" | jq '.responses.updatePlaylistResponse.schema["$ref"] = "#/definitions/Playlist"')" # Currently pointing to Spec (Preferences)
SCHEMA="$(echo "${SCHEMA}" | jq '.responses.getPlaylistResponse.schema["$ref"] = "#/definitions/Playlist"')" # Currently pointing to Spec (Preferences)


# Write the schema to a file
echo "${SCHEMA}" > "${SCRIPT_DIR}/schema.json"
