package atomx

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"strconv"
)

const (
	DefaultApiUrl = "https://api.atomx.com/v2/"
)

type Client struct {
	ApiUrl string

	client http.Client
}

func New() *Client {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	return &Client{
		ApiUrl: DefaultApiUrl,
		client: http.Client{
			Jar: jar,
		},
	}
}

func (c *Client) Url(path string) string {
	if c.ApiUrl == "" {
		return DefaultApiUrl + path
	} else {
		return c.ApiUrl + path
	}
}

func (c *Client) Login(email, password string) error {
	type logindata struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	data, err := json.Marshal(logindata{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return err
	}

	res, err := c.client.Post(c.Url("login"), "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var response struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
	}

	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&response); err != nil {
		return err
	}

	if !response.Success {
		return &ApiError{Message: response.Error}
	}

	return nil
}

func (c *Client) Site(id int) (*Site, error) {
	path := "site/" + strconv.FormatInt(int64(id), 10)

	res, err := c.client.Get(c.Url(path))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var response struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
		Site    *Site  `json:"site"`
	}

	/*body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	println(string(body))
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}*/

	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&response); err != nil {
		return nil, err
	}

	if !response.Success {
		return nil, &ApiError{Message: response.Error}
	}

	return response.Site, nil
}

func (c *Client) Put(obj Resource) error {
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", c.Url(obj.Path()), bytes.NewReader(data))
	if err != nil {
		return err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var response struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
	}

	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&response); err != nil {
		return err
	}

	if !response.Success {
		return &ApiError{Message: response.Error}
	}

	return nil
}
