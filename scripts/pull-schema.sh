#!/usr/bin/env bash

# https://stackoverflow.com/a/246128
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# Pull the schema (stable commit)
SCHEMA="$(curl -s -L https://raw.githubusercontent.com/grafana/grafana/2e5b41cbcb403d7dee6ecdfacb94a6bc52cd6e90/public/api-merged.json)"

# Custom extensions: https://goswagger.io/use/models/schemas.html#custom-extensions
# These may have to be updated for future versions of Grafana
modify() {
    SCHEMA="$(echo "${SCHEMA}" | jq "${1}")"
}

# Playlist models are all messed up
# TODO: Upstream fix
modify '.responses.getPlaylistItemsResponse.schema.items["$ref"] = "#/definitions/PlaylistItem"' # Currently pointing to Item (old PlaylistItem model)
modify '.responses.updatePlaylistResponse.schema["$ref"] = "#/definitions/Playlist"' # Currently pointing to Spec (Preferences)
modify '.responses.getPlaylistResponse.schema["$ref"] = "#/definitions/Playlist"' # Currently pointing to Spec (Preferences)


# Remove "Route" prefixes to operation IDs (ex: RouteGetxxx)
# TODO: Upstream fix
modify '.paths = .paths | walk(if type == "object" and has("operationId") then .operationId |= sub("^Route";"") else . end)'

# The "for" property returned by the API is a string (can't be unmarshaled to time.Duration)
# https://github.com/grafana/grafana/pull/90841
modify '.definitions.ProvisionedAlertRule.properties.for = {
    "type": "string",
    "format": "duration"
}'

# StartDate and EndDate of reports must be nullable
# https://github.com/grafana/grafana/pull/90845
modify '.definitions.ReportSchedule.properties.startDate["x-nullable"] = true'
modify '.definitions.ReportSchedule.properties.endDate["x-nullable"] = true'

# Alerting validation error is wrong. Message doesn't show up
# https://github.com/grafana/grafana/pull/90846
modify '.definitions.ValidationError.properties.message = .definitions.ValidationError.properties.msg'
modify 'del(.definitions.ValidationError.properties.msg)'

# Remap field time_intervals of MuteTimeInterval and TimeInterval from TimeInterval (collision) to an equivalent model TimeIntervalItem.
modify '.definitions.TimeInterval.properties.time_intervals.items["$ref"] = "#/definitions/TimeIntervalItem"'
modify '.definitions.MuteTimeInterval.properties.time_intervals.items["$ref"] = "#/definitions/TimeIntervalItem"'

# Write the schema to a file
echo "${SCHEMA}" > "${SCRIPT_DIR}/schema.json"
