package custom_models

import (
	"bytes"
	"encoding/json"
)

type JSON map[string]interface{}

// Basic Unmarshal. Do not UseNumber (unmarshal to float64 over json.Number)
func (j *JSON) UnmarshalJSON(data []byte) error {
	type j2 JSON
	dec := json.NewDecoder(bytes.NewReader(data))
	return dec.Decode((*j2)(j))
}
