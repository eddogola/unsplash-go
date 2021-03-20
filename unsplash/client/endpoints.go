package client

const (
	// BaseEndpoint defines the base api address
	BaseEndpoint = "https://api.unsplash.com/"
	// BaseUserEndpoint defines the base api address for user resources
	BaseUserEndpoint = BaseEndpoint + "users/"
	// PrivateUserProfileEndpoint defines the api address for accessing a user's private profile
	PrivateUserProfileEndpoint = BaseEndpoint + "me"
	// RandomPhotoEndpoint defines the api address for getting a random photo
	RandomPhotoEndpoint = BaseEndpoint + "photos/random/"
	// AllPhotosEndpoint defines the api address for photos resources
	AllPhotosEndpoint = BaseEndpoint + "photos/"
	// SearchPhotosEndpoint defines the api address for searching photo resources
	SearchPhotosEndpoint = BaseEndpoint + "search/photos"
	// SearchCollectionsEndpoint defines the api address for searching collection resources
	SearchCollectionsEndpoint = BaseEndpoint + "search/collections"
	// SearchUsersEndpoint defines the api address for searching users resources
	SearchUsersEndpoint = BaseEndpoint + "search/users"
	// TopicsListEndpoint defines the api address for topic resources
	TopicsListEndpoint = BaseEndpoint + "topics/"
	// CollectionsListEndpoint defines the api address for collections resources
	CollectionsListEndpoint = BaseEndpoint + "collections/"
	// StatsTotalEndpoint defines the api address for total Unsplash stats
	StatsTotalEndpoint = BaseEndpoint + "stats/total"
	// StatsMonthEndpoint defines the api address for monthly Unsplash stats
	StatsMonthEndpoint = BaseEndpoint + "stats/month"
)

// Private authorization endpoints
const (
	AuthCodeEndpoint  = "https://unsplash.com/oauth/authorize"
	AuthTokenEndpoint = "https://unsplash.com/oauth/token"
)
