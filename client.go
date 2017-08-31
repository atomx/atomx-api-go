package atomx

//go:generate go run gen.go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	DefaultApiURL    = "https://api.atomx.com/v3/"
	DefaultUserAgent = "atomx-api-go"
)

type Client struct {
	ApiURL    string
	UserAgent string
	User      User
	AuthToken string

	client http.Client
}

func New() *Client {
	return &Client{
		ApiURL:    DefaultApiURL,
		UserAgent: DefaultUserAgent,
		client:    http.Client{},
	}
}

func (c *Client) Login(email, password, totp string) error {
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
		Success      bool   `json:"success"`
		Error        string `json:"error"`
		AuthToken    string `json:"auth_token"`
		User         User   `json:"user"`
		TOTPRequired bool   `json:"totp_required"`
	}

	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&response); err != nil {
		return err
	}

	if !response.Success {
		return &ApiError{Message: response.Error}
	}

	c.AuthToken = response.AuthToken

	if response.TOTPRequired {
		return c.totp(totp)
	}

	c.User = response.User

	return nil
}

func (c *Client) totp(totp string) error {
	url := c.ApiURL + "totp"

	type logindata struct {
		TOTP string `json:"totp"`
	}

	data, err := json.Marshal(logindata{
		TOTP: totp,
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

	res, err := c.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var response struct {
		Success   bool   `json:"success"`
		Error     string `json:"error"`
		AuthToken string `json:"auth_token"`
		User      User   `json:"user"`
	}

	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&response); err != nil {
		return err
	}

	if !response.Success {
		return &ApiError{Message: response.Error}
	}

	c.AuthToken = response.AuthToken
	c.User = response.User

	return nil
}

func (c *Client) Do(req *http.Request) (resp *http.Response, err error) {
	req.Header.Add("User-Agent", c.UserAgent)

	if c.AuthToken != "" {
		req.Header.Add("Authorization", "Bearer "+c.AuthToken)
	}

	return c.client.Do(req)
}

func (c *Client) Get(obj Resource, opts *Options) error {
	url := c.ApiURL + obj.path() + "?" + opts.str()

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

	response := obj.response()

	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("%v: %q", err, body)
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

	res, err := c.Do(req)
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
		return fmt.Errorf("%v: %q", err, body)
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

	res, err := c.Do(req)
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
		return fmt.Errorf("%v: %q", err, body)
	}

	if err := response.err(); err != nil {
		return err
	}

	return nil
}
