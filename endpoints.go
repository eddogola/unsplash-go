package unsplash

const (
	BaseEndpoint               = "https://api.unsplash.com/"
	BaseUserEndpoint           = BaseEndpoint + "users/"
	PrivateUserProfileEndpoint = BaseEndpoint + "me"
	RandomPhotoEndpoint        = BaseEndpoint + "photos/random/"
	AllPhotosEndpoint          = BaseEndpoint + "photos/"
	SearchPhotosEndpoint       = BaseEndpoint + "search/photos"
	SearchCollectionsEndpoint  = BaseEndpoint + "search/collections"
	SearchUsersEndpoint        = BaseEndpoint + "search/users"
	TopicsListEndpoint         = BaseEndpoint + "topics/"
	CollectionsListEndpoint    = BaseEndpoint + "collections/"
	StatsTotalEndpoint         = BaseEndpoint + "stats/total"
	StatsMonthEndpoint         = BaseEndpoint + "stats/month"
)

// Private authorization endpoints
const (
	AuthCodeEndpoint  = "https://unsplash.com/oauth/authorize"
	AuthTokenEndpoint = "https://unsplash.com/oauth/token"
)