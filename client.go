package unsplash

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	baseEndpoint               = "https://api.unsplash.com/"
	baseUserEndpoint           = baseEndpoint + "users/"
	privateUserProfileEndpoint = baseEndpoint + "me"
	randomPhotoEndpoint        = baseEndpoint + "photos/random/"
	allPhotosEndpoint          = baseEndpoint + "photos/"
	searchPhotosEndpoint       = baseEndpoint + "search/photos"
	searchCollectionsEndpoint  = baseEndpoint + "search/collections"
	searchUsersEndpoint        = baseEndpoint + "search/users"
	topicsListEndpoint         = baseEndpoint + "topics/"
	collectionsListEndpoint    = baseEndpoint + "collections/"
	statsTotalEndpoint         = baseEndpoint + "stats/total"
	statsMonthEndpoint         = baseEndpoint + "stats/month"
)

// QueryParams defines url link paramaters
type QueryParams map[string]string

// Client defines methods to interact with the Unsplash API
type Client struct {
	ClientID   string
	HTTPClient *http.Client
	Config     *Config
	Private    bool // true if private authentication is required to make requests, default should be false
	AuthScopes AuthScopes
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
	if err != nil {
		return nil, err
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
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
	if err != nil {
		return nil, err
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
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
func (c *Client) getUserPortfolioLink(ctx context.Context, username string) (string, error) {
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

// Retrieve the consolidated number of downloads, views and likes of all user’s photos,
// as well as the historical breakdown and average of these stats in a specific timeframe (default is 30 days).
// https://unsplash.com/documentation#get-a-users-statistics
func (c *Client) getUserStats(ctx context.Context, username string, queryParams QueryParams) (*UserStats, error) {
	endPoint := baseUserEndpoint + username + "/statistics"
	link, err := buildURL(endPoint, queryParams)
	if err != nil {
		return nil, err
	}
	data, err := c.getBodyBytes(ctx, link)
	if err != nil {
		return nil, err
	}
	var us UserStats
	err = parseJSON(data, &us)
	if err != nil {
		return nil, err
	}
	return &us, nil
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
func (c *Client) getPhoto(ctx context.Context, ID string) (*Photo, error) {
	link := allPhotosEndpoint + ID
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

// Retrieve total number of downloads, views and likes of a single photo, as well as the historical
// breakdown of these stats in a specific timeframe (default is 30 days).
// https://unsplash.com/documentation#get-a-photos-statistics
func (c *Client) getPhotoStats(ctx context.Context, ID string, queryParams QueryParams) (*PhotoStats, error) {
	endPoint := allPhotosEndpoint + ID + "/statistics"
	link, err := buildURL(endPoint, queryParams)
	if err != nil {
		return nil, err
	}
	data, err := c.getBodyBytes(ctx, link)
	if err != nil {
		return nil, err
	}
	var ps PhotoStats
	err = parseJSON(data, &ps)
	if err != nil {
		return nil, err
	}
	return &ps, nil
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

// stats functions

// Get a list of counts for all of Unsplash.
// https://unsplash.com/documentation#totals
func (c *Client) getStatsTotal(ctx context.Context) (*StatsTotal, error) {
	data, err := c.getBodyBytes(ctx, statsTotalEndpoint)
	if err != nil {
		return nil, err
	}
	var stats StatsTotal
	err = parseJSON(data, stats)
	if err != nil {
		return nil, err
	}
	return &stats, nil
}

// Get the overall Unsplash stats for the past 30 days.
// https://unsplash.com/documentation#month
func (c *Client) getStatsMonth(ctx context.Context) (*StatsMonth, error) {
	data, err := c.getBodyBytes(ctx, statsTotalEndpoint)
	if err != nil {
		return nil, err
	}
	var stats StatsMonth
	err = parseJSON(data, stats)
	if err != nil {
		return nil, err
	}
	return &stats, nil
}

// private Client methods
// Note: Without a Bearer token (i.e. using a Client-ID token) these requests will return a 401 Unauthorized response.

// Get the user’s private profile
// Note: To access a user’s private data, the user is required to authorize the read_user scope.
func (c *Client) getUserPrivateProfile(ctx context.Context) (*User, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, errClientNotPrivate
	}
	// check if the `read_user` scope is present in the private Client's scopes
	scope := "read_user"
	if ok := c.AuthScopes.Contains(scope); !ok {
		return nil, errRequiredScopeAbsent(scope)
	}

	data, err := c.getBodyBytes(ctx, privateUserProfileEndpoint)
	if err != nil {
		return nil, err
	}

	var usr User
	err = parseJSON(data, usr)
	if err != nil {
		return nil, err
	}
	return &usr, nil
}

// Update the current user’s profile
// https://unsplash.com/documentation#update-the-current-users-profile
// Note: This action requires the write_user scope. Without it, it will return a 403 Forbidden response.
func (c *Client) updateUserProfile(ctx context.Context, updatedData map[string]string) (*User, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, errClientNotPrivate
	}
	// check if the `write_user` scope is present in the private Client's scopes
	scope := "write_user"
	if ok := c.AuthScopes.Contains(scope); !ok {
		return nil, errRequiredScopeAbsent(scope)
	}
	// make PUT request
	// response returns the updated profile
	resp, err := c.putHTTP(ctx, privateUserProfileEndpoint, updatedData)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// parse bytes
	var usr User
	err = parseJSON(data, &usr)
	if err != nil {
		return nil, err
	}
	return &usr, nil
}

// Update a photo on behalf of the logged-in user
// This requires the `write_photos` scope
// https://unsplash.com/documentation#update-a-photo
func (c *Client) updatePhoto(ctx context.Context, ID string, updatedData map[string]string) (*Photo, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, errClientNotPrivate
	}
	// check if the `write_photo` scope is present in the private Client's scopes
	scope := "write_photo"
	if ok := c.AuthScopes.Contains(scope); !ok {
		return nil, errRequiredScopeAbsent(scope)
	}
	// make PUT request
	// response returns the updated profile
	endPoint := allPhotosEndpoint + ID
	resp, err := c.putHTTP(ctx, endPoint, updatedData)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// parse bytes
	var pic Photo
	err = parseJSON(data, &pic)
	if err != nil {
		return nil, err
	}
	return &pic, nil
}

// Like a photo on behalf of the logged-in user
// This requires the `write_likes` scope
// https://unsplash.com/documentation#like-a-photo
func (c *Client) likePhoto(ctx context.Context, ID string) (*LikeResponse, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, errClientNotPrivate
	}
	// check if the `write_likes` scope is present in the private Client's scopes
	scope := "write_likes"
	if ok := c.AuthScopes.Contains(scope); !ok {
		return nil, errRequiredScopeAbsent(scope)
	}
	// make POST request
	// response returns abbreviated versions of the picture and user
	endPoint := allPhotosEndpoint + ID + "/like"
	resp, err := c.postHTTP(ctx, endPoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// parse json response
	var lr LikeResponse
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = parseJSON(data, &lr)
	if err != nil {
		return nil, err
	}
	return &lr, nil
}

// Remove the logged-in user’s like of a photo.
// https://unsplash.com/documentation#unlike-a-photo
func (c *Client) unlikePhoto(ctx context.Context, ID string) error {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return errClientNotPrivate
	}
	// make DELETE request
	// responds with a 204 status code and an empty body
	endPoint := allPhotosEndpoint + ID + "/like"
	resp, err := c.deleteHTTP(ctx, endPoint, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 204 {
		return errStatusCode{resp.StatusCode, getErrReasons(resp)}
	}
	return nil
}

// Create a new collection. This requires the `write_collections` scope.
// https://unsplash.com/documentation#create-a-new-collection
func (c *Client) createCollection(ctx context.Context, ID string, data map[string]string) (*Collection, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, errClientNotPrivate
	}
	// check if the `write_collections` scope is present in the private Client's scopes
	scope := "write_collections"
	if ok := c.AuthScopes.Contains(scope); !ok {
		return nil, errRequiredScopeAbsent(scope)
	}
	// make POST request
	// responds with the new collection
	resp, err := c.postHTTP(ctx, collectionsListEndpoint, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// parse json response
	var collection Collection
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = parseJSON(bs, &collection)
	if err != nil {
		return nil, err
	}
	return &collection, nil
}

// Update an existing collection belonging to the logged-in user. This requires the `write_collections` scope.
// https://unsplash.com/documentation#update-an-existing-collection
// check if client is private to do private requests
func (c *Client) updateCollection(ctx context.Context, ID string, data map[string]string) (*Collection, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, errClientNotPrivate
	}
	// check if the `write_collections` scope is present in the private Client's scopes
	scope := "write_collections"
	if ok := c.AuthScopes.Contains(scope); !ok {
		return nil, errRequiredScopeAbsent(scope)
	}
	// make PUT request
	// responds with the updated collection
	endPoint := collectionsListEndpoint + ID
	resp, err := c.putHTTP(ctx, endPoint, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// parse bytes
	var collection Collection
	err = parseJSON(bs, &collection)
	if err != nil {
		return nil, err
	}
	return &collection, nil
}

// Delete a collection belonging to the logged-in user. This requires the `write_collections` scope.
// https://unsplash.com/documentation#delete-a-collection
func (c *Client) deleteCollection(ctx context.Context, ID string) error {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return errClientNotPrivate
	}
	// check if the `write_collections` scope is present in the private Client's scopes
	scope := "write_collections"
	if ok := c.AuthScopes.Contains(scope); !ok {
		return errRequiredScopeAbsent(scope)
	}
	// make DELETE request
	// responds with a 204 status code and an empty body
	endPoint := collectionsListEndpoint + ID
	resp, err := c.deleteHTTP(ctx, endPoint, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return errStatusCode{resp.StatusCode, getErrReasons(resp)}
	}
	return nil
}

// Add a photo to one of the logged-in user’s collections. Requires the `write_collections` scope.
// https://unsplash.com/documentation#add-a-photo-to-a-collection
func (c *Client) addPhotoToCollection(ctx context.Context, ID string, data map[string]string) (*CollectionActionResponse, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, errClientNotPrivate
	}
	// check if the `write_collections` scope is present in the private Client's scopes
	scope := "write_collections"
	if ok := c.AuthScopes.Contains(scope); !ok {
		return nil, errRequiredScopeAbsent(scope)
	}
	// make POST request
	// response returns abbreviated versions of the picture and user
	endPoint := collectionsListEndpoint + ID + "/add"
	resp, err := c.postHTTP(ctx, endPoint, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// parse json response
	var car CollectionActionResponse
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = parseJSON(bs, &car)
	if err != nil {
		return nil, err
	}
	return &car, nil
}

// Remove a photo from one of the logged-in user’s collections. Requires the `write_collections` scope.
// https://unsplash.com/documentation#remove-a-photo-from-a-collection
func (c *Client) removePhotoFromCollection(ctx context.Context, ID string, data map[string]string) (*CollectionActionResponse, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, errClientNotPrivate
	}
	// check if the `write_collections` scope is present in the private Client's scopes
	scope := "write_collections"
	if ok := c.AuthScopes.Contains(scope); !ok {
		return nil, errRequiredScopeAbsent(scope)
	}
	// make DELETE request
	// responds with a 204 status code and an empty body
	endPoint := collectionsListEndpoint + ID + "/remove"
	resp, err := c.deleteHTTP(ctx, endPoint, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 204 {
		return nil, errStatusCode{resp.StatusCode, getErrReasons(resp)}
	}

	// parse json response
	var car CollectionActionResponse
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = parseJSON(bs, &car)
	if err != nil {
		return nil, err
	}
	return &car, nil
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

func isClientPrivate(c *Client) bool {
	return c.Private
}
