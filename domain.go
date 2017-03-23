package atomx

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Domain struct {
	ID         int64            `json:"id"`
	Name       string           `json:"name,omitempty"`
	CategoryID int64            `json:"category,omitempty"`
	Attributes DomainAttributes `json:"attributes,omitempty"`
}

func (c *Client) PostDomains(body string) ([]Domain, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.ApiURL+"domains", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var resp struct {
		Success bool     `json:"success"`
		Error   string   `json:"error"`
		Domains []Domain `json:"domains"`
	}

	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&resp); err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, &ApiError{Message: resp.Error}
	}

	return resp.Domains, nil
}
