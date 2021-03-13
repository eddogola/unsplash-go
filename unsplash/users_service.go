package unsplash

import (
	"context"
	"net/url"

	"github.com/eddogola/unsplash-go/unsplash/client"
)

type UsersService struct {
	client *client.Client
}

func (us *UsersService) PublicProfile(ctx context.Context, username string) (*client.User, error) {
	return us.client.GetUserPublicProfile(ctx, username)
}

func (us *UsersService) PortfolioURL(ctx context.Context, username string) (*url.URL, error) {
	return us.client.GetUserPortfolioLink(ctx, username)
}

func (us *UsersService) Photos(ctx context.Context, username string, queryParams client.QueryParams) ([]client.Photo, error) {
	return us.client.GetUserPhotos(ctx, username, queryParams)
}

func (us *UsersService) LikedPhotos(ctx context.Context, username string, queryParams client.QueryParams) ([]client.Photo, error) {
	return us.client.GetUserLikedPhotos(ctx, username, queryParams)
}

func (us *UsersService) Collections(ctx context.Context, username string, queryParams client.QueryParams) ([]client.Collection, error) {
	return us.client.GetUserCollections(ctx, username, queryParams)
}

func (us *UsersService) Stats(ctx context.Context, username string, queryParams client.QueryParams) (*client.UserStats, error) {
	return us.client.GetUserStats(ctx, username, queryParams)
}

func (us *UsersService) Search(ctx context.Context, queryParams client.QueryParams) (*client.UserSearchResult, error) {
	return us.client.SearchUsers(ctx, queryParams)
}

// methods requiring private authentication
func (us *UsersService) PrivateProfile(ctx context.Context) (*client.User, error) {
	return us.client.GetUserPrivateProfile(ctx)
}

func (us *UsersService) UpdateProfile(ctx context.Context, updatedData map[string]string) (*client.User, error) {
	return us.client.UpdateUserProfile(ctx, updatedData)
}
