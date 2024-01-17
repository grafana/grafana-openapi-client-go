#!/usr/bin/env bash

# https://stackoverflow.com/a/246128
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# Pull the schema (main branch)
SCHEMA="$(curl -s -L https://raw.githubusercontent.com/grafana/grafana/dd7259b77e3b185eeb4b06b57b3cdae54c17da33/public/api-merged.json)"

# Custom extensions: https://goswagger.io/use/models/schemas.html#custom-extensions
# These may have to be updated for future versions of Grafana
modify() {
    SCHEMA="$(echo "${SCHEMA}" | jq "${1}")"
}

# TODO: Remove on next Terraform provider version
# This is a deprecated attribute for newer Grafana versions
modify '.definitions.Preferences.properties.homeDashboardId = {
  "description": "ID for the home dashboard. This is deprecated and will be removed in a future version. Use homeDashboardUid instead.",
  "type": "integer"
}'

# Playlist models are all messed up
# TODO: Upstream fix
modify 'del(.definitions.Item)' # Old playlist item schema, PlaylistItem is more up to date
modify '.responses.getPlaylistItemsResponse.schema.items["$ref"] = "#/definitions/PlaylistItem"' # Currently pointing to Item (old PlaylistItem model)
modify '.responses.updatePlaylistResponse.schema["$ref"] = "#/definitions/Playlist"' # Currently pointing to Spec (Preferences)
modify '.responses.getPlaylistResponse.schema["$ref"] = "#/definitions/Playlist"' # Currently pointing to Spec (Preferences)

# Any endpoint that starts with /api/ should be trimmed because it's redundant (API path is configured on the client), ex: /api/dashboards/ -> /dashboards/
# Move /api/ map keys to a new key without /api/ prefix
# Fixed here: https://github.com/grafana/grafana/pull/79025
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

# StartDate and EndDate of reports must be nullable
# TODO: Upstream fix
modify '.definitions.ReportSchedule.properties.startDate["x-nullable"] = true'
modify '.definitions.ReportSchedule.properties.endDate["x-nullable"] = true'

# Fixed here: https://github.com/grafana/grafana/pull/79477
modify '.definitions += {
    "ObjectMatcher": {
      "type": "array",
      "title": "ObjectMatcher is a matcher that can be used to filter alerts.",
      "items": {
        "type": "string"
      }
    },
    "ObjectMatchers": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/ObjectMatcher"
      }
    }
}'

# Mutetiming TimeInterval is wrong
modify '.paths["/v1/provisioning/mute-timings/{name}"].put.responses["202"] = .paths["/v1/provisioning/mute-timings/{name}"].put.responses["200"]'
modify 'del(.paths["/v1/provisioning/mute-timings/{name}"].put.responses["200"])'
modify '.definitions.TimeIntervalRange = {
    "type": "object",
    "properties": {
      "start_time": {
        "type": "string",
      },
      "end_time": {
        "type": "string",
      }
    }
}'
modify '.definitions.TimeInterval.properties.times.items["$ref"] = "#/definitions/TimeIntervalRange"'

# Write the schema to a file
echo "${SCHEMA}" > "${SCRIPT_DIR}/schema.json"
