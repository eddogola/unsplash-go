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

// user functions

// Retrieve public details on a given user.
// https://unsplash.com/documentation#get-a-users-public-profile
func (c *Client) getUserPublicProfile(ctx context.Context, username string) (*User, error) {
	endPoint := baseUserEndpoint + username
	data, err := c.getBodyBytes(ctx, endPoint)
	if err != nil {
		return nil, err
	}
	var user User
	err = parseJSON(data, user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Retrieve a single user’s portfolio link.
// https://unsplash.com/documentation#get-a-users-portfolio-link
func (c *Client) getUserPortfolioLink(ctx context.Context, username string) (string, error ) {
	endPoint := baseUserEndpoint + username + "/portfolio"
	data, err := c.getBodyBytes(ctx, endPoint)
	if err != nil {
		return "", err
	}

	// parse json response data using local struct since it has only onee field
	var resp struct {
		Link string `json:"url"`
	}
	err = parseJSON(data, &resp)
	if err != nil {
		return "", err
	}
	return resp.Link, nil
}

// Get a list of photos uploaded by a user.
// https://unsplash.com/documentation#list-a-users-photos
func (c *Client) getUserPhotos(ctx context.Context, username string, queryParams QueryParams) ([]Photo, error) {
	endPoint := baseUserEndpoint + username
	link, err := buildURL(endPoint, queryParams)
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

// Get a list of photos liked by a user.
// https://unsplash.com/documentation#list-a-users-liked-photos
func (c *Client) getUserLikedPhotos(ctx context.Context, username string, queryParams QueryParams) ([]Photo, error) {
	endPoint := baseUserEndpoint + username + "/likes"
	link, err := buildURL(endPoint, queryParams)
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

// Get a list of collections created by the user.
// https://unsplash.com/documentation#list-a-users-collections
func (c *Client) getUserCollections(ctx context.Context, username string, queryParams QueryParams) ([]Collection, error) {
	endPoint := baseUserEndpoint + username + "/collections"
	link, err := buildURL(endPoint, queryParams)
	if err != nil {
		return nil, err
	}
	data, err := c.getBodyBytes(ctx, link)
	if err != nil {
		return nil, err
	}
	var collections []Collection
	err = parseJSON(data, &collections)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

// photo functions

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

// collections functions

// gets a single page of a list of collections
// https://unsplash.com/documentation#list-collections
func (c *Client) getCollectionsList(ctx context.Context, queryParams QueryParams) ([]Collection, error) {
	link, err := buildURL(collectionsListEndpoint, queryParams)
	if err != nil {
		return nil, err
	}
	data, err := c.getBodyBytes(ctx, link)
	if err != nil {
		return nil, err
	}
	var collections []Collection
	err = parseJSON(data, collections)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

// get a collection using id
// https://unsplash.com/documentation#get-a-collection
func (c *Client) getCollection(ctx context.Context, id int) (*Collection, error) {
	endPoint := collectionsListEndpoint + fmt.Sprint(id)
	data, err := c.getBodyBytes(ctx, endPoint)
	if err != nil {
		return nil, err
	}
	var collection Collection
	err = parseJSON(data, &collection)
	if err != nil {
		return nil, err
	}
	return &collection, nil
}

// Retrieve a collection's photos
// https://unsplash.com/documentation#get-a-collections-photos
func (c *Client) getCollectionPhotos(ctx context.Context, id int, queryParams QueryParams) ([]Photo, error) {
	endPoint := collectionsListEndpoint + fmt.Sprint(id) + "/photos"
	link, err := buildURL(endPoint, queryParams)
	if err != nil {
		return nil, err
	}
	data, err := c.getBodyBytes(ctx, link)
	if err != nil {
		return nil, err
	}
	var pics []Photo
	err = parseJSON(data, pics)
	if err != nil {
		return nil, err
	}
	return pics, nil
}

// Retrieve a list of collections related to this one.
// https://unsplash.com/documentation#list-a-collections-related-collections
func (c *Client) getRelatedCollections(ctx context.Context, id int) ([]Collection, error) {
	endPoint := collectionsListEndpoint + fmt.Sprint(id) + "/related"
	data, err := c.getBodyBytes(ctx, endPoint)
	if err != nil {
		return nil, err
	}
	var collections []Collection
	err = parseJSON(data, collections)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

// topics functions

// Get a single page from the list of all topics.
// https://unsplash.com/documentation#list-topics
func (c *Client) getTopicsList(ctx context.Context, queryParams QueryParams) ([]Topic, error) {
	link, err := buildURL(topicsListEndpoint, queryParams)
	if err != nil {
		return nil, err
	}
	data, err := c.getBodyBytes(ctx, link)
	if err != nil {
		return nil, err
	}
	var topics []Topic
	err = parseJSON(data, topics)
	if err != nil {
		return nil, err
	}
	return topics, nil
}

// Retrieve a single topic.
// https://unsplash.com/documentation#get-a-topic
func (c *Client) getTopic(ctx context.Context, IDOrSlug string) (*Topic, error) {
	endPoint := topicsListEndpoint + IDOrSlug
	data, err := c.getBodyBytes(ctx, endPoint)
	if err != nil {
		return nil, err
	}
	var topic Topic
	err = parseJSON(data, &topic)
	if err != nil {
		return nil, err
	}
	return &topic, nil
}

// Retrieve a topic’s photos.
// https://unsplash.com/documentation#get-a-topics-photos
func (c *Client) getTopicPhotos(ctx context.Context, IDOrSlug string, queryParams QueryParams) ([]Photo, error) {
	endPoint := topicsListEndpoint + IDOrSlug + "/photos"
	link, err := buildURL(endPoint, queryParams)
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
	return pics, err
}

// utility functions

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
