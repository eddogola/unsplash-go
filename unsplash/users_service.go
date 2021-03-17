package unsplash

import (
	"context"
	"net/url"

	"github.com/eddogola/unsplash-go/unsplash/client"
)

// UsersServiceClient defines client methods used to get user resources
type UsersServiceClient interface {
	GetUserPublicProfile(context.Context, string) (*client.User, error)
	GetUserPortfolioLink(context.Context, string) (*url.URL, error)
	GetUserPhotos(context.Context, string, client.QueryParams) ([]client.Photo, error)
	GetUserLikedPhotos(context.Context, string, client.QueryParams) ([]client.Photo, error)
	GetUserCollections(context.Context, string, client.QueryParams) ([]client.Collection, error)
	GetUserStats(context.Context, string, client.QueryParams) (*client.UserStats, error)
	SearchUsers(context.Context, client.QueryParams) (*client.UserSearchResult, error)
	GetUserPrivateProfile(context.Context) (*client.User, error)
	UpdateUserProfile(context.Context, map[string]string) (*client.User, error)
}

// UsersService contains an underlying Unsplash client to
//be used for http methods
type UsersService struct {
	client UsersServiceClient
}

// PublicProfile returns the public profile of the user
func (us *UsersService) PublicProfile(ctx context.Context, username string) (*client.User, error) {
	return us.client.GetUserPublicProfile(ctx, username)
}

// PortfolioURL returns a parsed URL of the user
func (us *UsersService) PortfolioURL(ctx context.Context, username string) (*url.URL, error) {
	return us.client.GetUserPortfolioLink(ctx, username)
}

// Photos returns a paginated list of Photos uploaded by the user
func (us *UsersService) Photos(ctx context.Context, username string, queryParams client.QueryParams) ([]client.Photo, error) {
	return us.client.GetUserPhotos(ctx, username, queryParams)
}

// LikedPhotos returns a paginated list of photos liked by the user
func (us *UsersService) LikedPhotos(ctx context.Context, username string, queryParams client.QueryParams) ([]client.Photo, error) {
	return us.client.GetUserLikedPhotos(ctx, username, queryParams)
}

// Collections returns a paginated list of collections created by the user
func (us *UsersService) Collections(ctx context.Context, username string, queryParams client.QueryParams) ([]client.Collection, error) {
	return us.client.GetUserCollections(ctx, username, queryParams)
}

// Stats returns the user's stats
func (us *UsersService) Stats(ctx context.Context, username string, queryParams client.QueryParams) (*client.UserStats, error) {
	return us.client.GetUserStats(ctx, username, queryParams)
}

// Search takes in a search query under the given query parameters to return a list of User search results
func (us *UsersService) Search(ctx context.Context, searchQuery string, queryParams client.QueryParams) (*client.UserSearchResult, error) {
	if queryParams == nil || queryParams["query"] == "" {
		queryParams = make(client.QueryParams)
		queryParams["query"] = searchQuery
		return us.client.SearchUsers(ctx, queryParams)
	}
	queryParams["query"] = searchQuery
	return us.client.SearchUsers(ctx, queryParams)
}

// methods requiring private authentication

// PrivateProfile returns the authenticated user's private profile
func (us *UsersService) PrivateProfile(ctx context.Context) (*client.User, error) {
	return us.client.GetUserPrivateProfile(ctx)
}

// UpdateProfile updates the authenticated user's profile using the data map provided
func (us *UsersService) UpdateProfile(ctx context.Context, updatedData map[string]string) (*client.User, error) {
	return us.client.UpdateUserProfile(ctx, updatedData)
}
