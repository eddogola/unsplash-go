package client

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func parseJSON(data []byte, desiredObject interface{}) error {
	if err := json.Unmarshal(data, desiredObject); err != nil {
		return fmt.Errorf("error parsing json to %T: %v", desiredObject, err)
	}
	return nil
}

func buildURL(link string, queryParams QueryParams) (string, error) {
	// add query params to request
	if queryParams != nil {
		URL, err := url.Parse(link)
		if err != nil {
			return "", err
		}
		q := URL.Query()
		for key, val := range queryParams {
			q.Add(key, val)
		}
		URL.RawQuery = q.Encode()
		return URL.String(), nil
	}
	return link, nil
}

func isClientPrivate(c *Client) bool {
	return c.Private
}
