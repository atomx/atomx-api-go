package atomx

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

type Domain struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (d *Domain) path() string {
	return "domain/" + strconv.FormatInt(int64(d.Id), 10)
}

type domainResponse struct {
	Success bool    `json:"success"`
	Error   string  `json:"error"`
	Domain  *Domain `json:"domain"`
}

func (dr *domainResponse) err() error {
	if !dr.Success {
		return &ApiError{Message: dr.Error}
	}

	return nil
}

func (d *Domain) response() response {
	return &domainResponse{
		Domain: d,
	}
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

	res, err := c.client.Do(req)
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
