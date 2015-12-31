package atomx

import (
	"encoding/json"
)

// Marshal the struct in obj but don't include the ID field.
func marshalWithoutID(obj interface{}) ([]byte, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	var raw map[string]*json.RawMessage

	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	delete(raw, "id")

	return json.Marshal(raw)
}
