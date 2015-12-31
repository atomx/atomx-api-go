package atomx

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

const (
	DefaultApiURL = "https://api.atomx.com/v2/"
)

type Client struct {
	ApiURL string

	client http.Client
}

func New() *Client {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	return &Client{
		ApiURL: DefaultApiURL,
		client: http.Client{
			Jar: jar,
		},
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

	res, err := c.client.Post(c.ApiURL+"login", "application/json", bytes.NewReader(data))
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

func (c *Client) Get(obj Resource, opts *Options) error {
	url := c.ApiURL + obj.path() + "?" + opts.str()

	res, err := c.client.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	response := obj.response()

	if err := json.Unmarshal(body, &response); err != nil {
		return err
	}

	return response.err()
}

func (c *Client) Put(obj Resource) error {
	data, err := marshalWithoutID(obj)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", c.ApiURL+obj.path(), bytes.NewReader(data))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

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

func (c *Client) List(objs Resources, opts *Options) error {
	url := c.ApiURL + objs.path() + "&" + opts.str()

	res, err := c.client.Get(url)
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
