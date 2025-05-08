package http

import (
	"bytes"
	"encoding/json"
	"net/http"

	logger "client_server/pkg/logger"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
	log        logger.Logger
}

func NewClient(baseURL string, log logger.Logger) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{},
		log:        log,
	}
}

func (c *Client) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.baseURL+url, nil)
	if err != nil {
		c.log.Error(err.Error())
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.log.Error(err.Error())
		return nil, err
	}

	return resp, nil
}

func (c *Client) Post(url string, body interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		c.log.Error(err.Error())
		return nil, err
	}

	req, err := http.NewRequest("POST", c.baseURL+url, bytes.NewBuffer(jsonBody))
	if err != nil {
		c.log.Error(err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.log.Error(err.Error())
		return nil, err
	}

	return resp, nil
}

func (c *Client) Put(url string, body interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		c.log.Error(err.Error())
		return nil, err
	}

	req, err := http.NewRequest("PUT", c.baseURL+url, bytes.NewBuffer(jsonBody))
	if err != nil {
		c.log.Error(err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.log.Error(err.Error())
		return nil, err
	}

	return resp, nil
}

func (c *Client) Delete(url string, body interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		c.log.Error(err.Error())
		return nil, err
	}

	req, err := http.NewRequest("DELETE", c.baseURL+url, bytes.NewBuffer(jsonBody))
	if err != nil {
		c.log.Error(err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.log.Error(err.Error())
		return nil, err
	}

	return resp, nil
}

func (c *Client) Patch(url string, body interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		c.log.Error(err.Error())
		return nil, err
	}

	req, err := http.NewRequest("PATCH", c.baseURL+url, bytes.NewBuffer(jsonBody))
	if err != nil {
		c.log.Error(err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.log.Error(err.Error())
		return nil, err
	}

	return resp, nil
}
