package atomx

//go:generate go run gen.go

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

const (
	DefaultApiURL    = "https://api.atomx.com/v2/"
	DefaultUserAgent = "atomx-api-go"
)

type Client struct {
	ApiURL    string
	UserAgent string
	User      User

	client http.Client
}

func New() *Client {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	return &Client{
		ApiURL:    DefaultApiURL,
		UserAgent: DefaultUserAgent,
		client: http.Client{
			Jar: jar,
		},
	}
}

func (c *Client) Login(email, password string) error {
	url := c.ApiURL + "login"

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

	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", c.UserAgent)

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var response struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
		User    User   `json:"user"`
	}

	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&response); err != nil {
		return err
	}

	if !response.Success {
		return &ApiError{Message: response.Error}
	}

	c.User = response.User

	return nil
}

func (c *Client) Get(obj Resource, opts *Options) error {
	url := c.ApiURL + obj.path() + "?" + opts.str()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("User-Agent", c.UserAgent)

	res, err := c.client.Do(req)
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

func (c *Client) Put(obj Resource, opts *Options) error {
	url := c.ApiURL + obj.path() + "?" + opts.str()

	data, err := marshalWithoutID(obj)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewReader(data))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", c.UserAgent)

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	response := obj.response()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return err
	}

	if err := response.err(); err != nil {
		return err
	}

	return nil
}

func (c *Client) Post(obj Resource, opts *Options) error {
	url := c.ApiURL + obj.path() + "?" + opts.str()

	data, err := marshalWithoutID(obj)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", c.UserAgent)

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	response := obj.response()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return err
	}

	if err := response.err(); err != nil {
		return err
	}

	return nil
}
