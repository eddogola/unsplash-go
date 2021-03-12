package unsplash

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

// Private authorization endpoints
const (
	authCodeEndpoint  = "https://unsplash.com/oauth/authorize"
	authTokenEndpoint = "https://unsplash.com/oauth/token"
)