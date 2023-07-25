package client

import (
	"io"
	"net/http"
)

type requestConfig struct {
	Method string
	Body   io.Reader
	Header map[string]string
}

const baseUrl = "https://api.squarecloud.app/v2/"

func WithMethod(method string) func(*requestConfig) {
	return func(c *requestConfig) {
		c.Method = method
	}
}

func WithHeader(key, value string) func(*requestConfig) {
	return func(c *requestConfig) {
		c.Header[key] = value
	}
}

func WithBody(body io.Reader) func(*requestConfig) {
	return func(c *requestConfig) {
		c.Body = body
	}
}

func (c *APIClient) request(url string, opts ...func(*requestConfig)) (*http.Response, error) {
	config := &requestConfig{
		Method: "GET",
		Body:   nil,
	}

	req, err := http.NewRequest(config.Method, url, config.Body)
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		opt(config)
	}

	req.Header.Set("Authorization", c.Config.Identity.IdentityToken)
	for key, value := range config.Header {
		req.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
