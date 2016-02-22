package atomx

import (
	"encoding/json"
	"strconv"
)

type CreativeBan struct {
	URL    string            `json:"url"`
	Reason CreativeBanReason `json:"reason"`
}

func (cb CreativeBan) MarshalJSON() ([]byte, error) {
	url, err := json.Marshal(cb.URL)
	if err != nil {
		return nil, err
	}

	return []byte(`{"url":` + string(url) + `,"reason_id":` + strconv.FormatInt(cb.Reason.ID, 10) + `}`), nil
}
