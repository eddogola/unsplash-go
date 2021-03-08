package unsplash

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	baseEndpoint        = "https://api.unsplash.com/"
	randomPhotoEndpoint = baseEndpoint + "photos/random"
)

// Client defines methods to interact with the Unsplash API
type Client struct {
	ClientID   string
	HTTPClient *http.Client
	Config     *Config
}

// Config sets up configuration details to be used in making requests.
// It contains headers that will be used in all client requests.
type Config struct {
	Headers http.Header
}

// NewClient initializes a new Client.
// if a client is not provided, a default http client is used.
func NewClient(clientID string, client *http.Client, config *Config) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	return &Client{ClientID: clientID, HTTPClient: client, Config: config}
}

// Utility functions to get data from the API using a context

func (c *Client) getHTTP(ctx context.Context, link string, queryParams map[string]string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		return nil, err
	}
	// add query params to request
	if queryParams != nil {
		q := req.URL.Query()
		for key, val := range queryParams {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) getBodyBytes(ctx context.Context, link string, queryParams map[string]string) ([]byte, error) {
	resp, err := c.getHTTP(ctx, link, queryParams)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, errStatusCode{resp.StatusCode, getErrReasons(resp)}
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) getRandomPhoto(ctx context.Context, queryParams map[string]string) (*Photo, error) {
	data, err := c.getBodyBytes(ctx, randomPhotoEndpoint, queryParams)
	if err != nil {
		return &Photo{}, err
	}

	// parse data to Photo object
	var pic Photo
	err = parseJSON(data, pic)
	if err != nil {
		return &Photo{}, err
	}
	return &pic, nil
}

func parseJSON(data []byte, desiredObject interface{}) error {
	if err := json.Unmarshal(data, desiredObject); err != nil {
		return fmt.Errorf("error parsing json to %T: %v", desiredObject, err)
	}
	return nil
}
