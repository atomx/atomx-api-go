package atomx

import (
	"encoding/json"
)

type CreativeCheck struct {
	URL string `json:"url"`
}

func (cc CreativeCheck) MarshalJSON() ([]byte, error) {
	url, err := json.Marshal(cc.URL)
	if err != nil {
		return nil, err
	}

	return []byte(`{"url":` + string(url) + `}`), nil
}
