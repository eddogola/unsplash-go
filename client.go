package unsplash

import (
	"net/http"
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
	LowContentSafety bool // if true, the `content_safety` parameter is set to `low`
	// otherwise, it is `high`
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
