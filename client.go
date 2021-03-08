package unsplash

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	baseEndpoint              = "https://api.unsplash.com/"
	baseUserEndpoint          = baseEndpoint + "users/"
	randomPhotoEndpoint       = baseEndpoint + "photos/random/"
	allPhotosEndpoint         = baseEndpoint + "photos/"
	searchPhotosEndpoint      = baseEndpoint + "search/photos"
	searchCollectionsEndpoint = baseEndpoint + "search/collections"
	searchUsersEndpoint       = baseEndpoint + "search/users"
	topicsListEndpoint        = baseEndpoint + "topics/"
	collectionsListEndpoint   = baseEndpoint + "collections/"
)

// QueryParams defines url link paramaters
type QueryParams map[string]string

// Client defines methods to interact with the Unsplash API
type Client struct {
	ClientID   string
	HTTPClient *http.Client
	Config     *Config
}

// Config sets up configuration details to be used in making requests.
// It contains headers that will be used in all client requests.
type Config struct {
	Headers      http.Header
	AuthInHeader bool // if true, Client-ID YOUR_ACCESS_KEY is added to the request Authentication header
	// if false, client_id is passed as a query parameter to the url being requested
}

// NewClient initializes a new Client.
// if a client is not provided, a default http client is used.
func NewClient(clientID string, client *http.Client, config *Config) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	config.Headers.Add("Accept-Version", "v1") // Add api version
	// Unsplash strongly encourages a specific request of the api version

	return &Client{ClientID: clientID, HTTPClient: client, Config: config}
}

// Utility functions to get data from the API using a context

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
	}
	return resp, nil
}

func (c *Client) getBodyBytes(ctx context.Context, link string) ([]byte, error) {
	resp, err := c.getHTTP(ctx, link)
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

// get a single page with a list of all photos
// https://unsplash.com/documentation#list-photos
func (c *Client) getPhotoList(ctx context.Context, queryParams QueryParams) ([]Photo, error) {
	link, err := buildURL(allPhotosEndpoint, queryParams)
	if err != nil {
		return nil, err
	}
	data, err := c.getBodyBytes(ctx, link)
	if err != nil {
		return nil, err
	}
	var pics []Photo
	err = parseJSON(data, &pics)
	if err != nil {
		return nil, err
	}
	return pics, nil
}

// get a Photo using photo ID
// https://unsplash.com/documentation#get-a-photo
func (c *Client) getPhoto(ctx context.Context, id string) (*Photo, error) {
	link := allPhotosEndpoint + id
	data, err := c.getBodyBytes(ctx, link)
	if err != nil {
		return nil, err
	}
	var pic Photo
	err = parseJSON(data, &pic)
	if err != nil {
		return nil, err
	}
	return &pic, nil
}

// get a random Photo
// return a list of photos if a count query parameter is provided
// https://unsplash.com/documentation#get-a-random-photo
func (c *Client) getRandomPhoto(ctx context.Context, queryParams QueryParams) (interface{}, error) {
	link, err := buildURL(randomPhotoEndpoint, queryParams)
	if err != nil {
		return nil, err
	}
	data, err := c.getBodyBytes(ctx, link)
	if err != nil {
		return nil, err
	}

	/* parse data to Photo object
	or []Photo is count query parameter is present

	From API documentation:
	Note: When supplying a count parameter - and only then -
	the response will be an array of photos, even if the value of count is 1.
	*/
	if _, ok := queryParams["count"]; ok {
		var pics []Photo
		err = parseJSON(data, pics)
		if err != nil {
			return nil, err
		}
		return &pics, nil
	}

	var pic Photo
	err = parseJSON(data, pic)
	if err != nil {
		return nil, err
	}
	return &pic, nil
}


// search functions

// get a single page with photo search results
// https://unsplash.com/documentation#search-photos
func (c *Client) searchPhotos(ctx context.Context, queryParams QueryParams) (*PhotoSearchResult, error) {
	link, err := buildURL(searchPhotosEndpoint, queryParams)
	if err != nil {
		return nil, err
	}
	// throw an error if search query parameter not in URL
	if _, ok := queryParams["query"]; !ok {
		return nil, errQueryNotInURL(link)
	}
	data, err := c.getBodyBytes(ctx, link)
	if err != nil {
		return nil, err
	}
	var res PhotoSearchResult
	err = parseJSON(data, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// get a single page with collection search results
// https://unsplash.com/documentation#search-collections
func (c *Client) searchCollections(ctx context.Context, queryParams QueryParams) (*CollectionSearchResult, error) {
	link, err := buildURL(searchCollectionsEndpoint, queryParams)
	if err != nil {
		return nil, err
	}
	// throw an error if search query parameter not in URL
	if _, ok := queryParams["query"]; !ok {
		return nil, errQueryNotInURL(link)
	}
	data, err := c.getBodyBytes(ctx, link)
	if err != nil {
		return nil, err
	}
	var res CollectionSearchResult
	err = parseJSON(data, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// get a single page with users search results
// https://unsplash.com/documentation#search-users
func (c *Client) searchUsers(ctx context.Context, queryParams QueryParams) (*UserSearchResult, error) {
	link, err := buildURL(searchUsersEndpoint, queryParams)
	if err != nil {
		return nil, err
	}
	// throw an error if search query parameter not in URL
	if _, ok := queryParams["query"]; !ok {
		return nil, errQueryNotInURL(link)
	}
	data, err := c.getBodyBytes(ctx, link)
	if err != nil {
		return nil, err
	}
	var res UserSearchResult
	err = parseJSON(data, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

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
