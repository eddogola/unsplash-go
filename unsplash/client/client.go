package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// QueryParams defines url link parameters
type QueryParams map[string]string

// Client defines methods to interact with the Unsplash API
type Client struct {
	ClientID   string
	HTTPClient *http.Client
	Config     *Config
	Private    bool // true if private authentication is required to make requests, default should be false
	AuthScopes *AuthScopes
}

// Config sets up configuration details to be used in making requests.
// It contains headers that will be used in all client requests.
type Config struct {
	Headers http.Header
}

// NewConfig constructs an empty Config object
func NewConfig() *Config {
	headers := make(http.Header)
	headers.Add("Content-Type", "application/json")
	headers.Add("Accept-Version", "v1") // Add api version
	// Unsplash strongly encourages a specific request of the api version
	// do sth to get access token
	return &Config{headers}
}

// New initializes a new Client.
// if a client is not provided, a default http client is used.
func New(clientID string, client *http.Client, config *Config) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	config.Headers.Add("Authorization", fmt.Sprintf("Client-ID %s", clientID))

	return &Client{ClientID: clientID, HTTPClient: client, Config: config, Private: false}
}

// Client http methods to get data from the API using a context

func (c *Client) getHTTP(ctx context.Context, link string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	// set request headers specified in Client.Config
	req.Header = c.Config.Headers
	if err != nil {
		return nil, err
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, ErrStatusCode{resp.StatusCode, getErrReasons(resp)}
	}
	return resp, nil
}

// get a response, by a post request
func (c *Client) postHTTP(ctx context.Context, link string, postData map[string]string) (*http.Response, error) {
	data, err := json.Marshal(postData)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, link, bytes.NewBuffer(data))
	req.Header = c.Config.Headers
	if err != nil {
		return nil, err
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusCreated {
		return nil, ErrStatusCode{resp.StatusCode, getErrReasons(resp)}
	}
	return resp, nil
}

// update resource using PUT
func (c *Client) putHTTP(ctx context.Context, link string, putData map[string]string) (*http.Response, error) {
	data, err := json.Marshal(putData)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, link, bytes.NewBuffer(data))
	req.Header = c.Config.Headers
	if err != nil {
		return nil, err
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	} else if (resp.StatusCode != http.StatusOK) || (resp.StatusCode != http.StatusCreated) {
		return nil, ErrStatusCode{resp.StatusCode, getErrReasons(resp)}
	}
	return resp, nil
}

// deletes resource using DELETE
func (c *Client) deleteHTTP(ctx context.Context, link string, dt map[string]string) (*http.Response, error) {
	data, err := json.Marshal(dt)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, link, bytes.NewBuffer(data))
	req.Header = c.Config.Headers
	if err != nil {
		return nil, err
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	} else if (resp.StatusCode != http.StatusOK) || (resp.StatusCode != http.StatusNoContent)  {
		return nil, ErrStatusCode{resp.StatusCode, getErrReasons(resp)}
	}
	return resp, nil
}

func (c *Client) getBodyBytes(ctx context.Context, link string) ([]byte, error) {
	resp, err := c.getHTTP(ctx, link)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
