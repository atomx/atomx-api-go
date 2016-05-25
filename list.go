package atomx

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type List struct {
	// These attributes are used as input and are set in the output.
	Offset int64  `json:"offset,omitempty"`
	Limit  int64  `json:"limit,omitempty"`
	Sort   string `json:"sort,omitempty"`

	// These attributes are only set in the output.
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Count   int64  `json:"count"`
}

func (l List) str() string {
	r := "offset=" + strconv.FormatInt(l.Offset, 10)

	if l.Limit > 0 {
		r += "&limit=" + strconv.FormatInt(l.Limit, 10)
	}

	if l.Sort != "" {
		r += "&sort=" + l.Sort
	}

	return r
}

func (l List) err() error {
	if !l.Success {
		return &ApiError{Message: l.Error}
	}

	return nil
}

func (c *Client) List(objs Resources, opts *Options) error {
	url := c.ApiURL + objs.path() + "&" + opts.str()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	res, err := c.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, objs); err != nil {
		return err
	}

	return objs.err()
}
