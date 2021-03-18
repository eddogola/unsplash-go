package unsplash

import (
	"context"
	"net/url"
	"testing"

	"github.com/eddogola/unsplash-go/unsplash/client"
)

var users = []client.User{
	{
		Username: "cole",
		PortfolioURL: "https://www.ogola.me",
		InstagramUsername: "lionkingonice",
		TotalPhotos: 30,
		Bio: "Imma do it so big they won't know what to call it",
	},
	{
		Username: "kendrick",
		PortfolioURL: "https://www.ogola.me",
		InstagramUsername: "damn",
		TotalPhotos: 39,
		Bio: "I got loyalty and royalty in my DNA",
	},
}

var user = users[0]

type mockUsersServiceClient struct{}

func (m *mockUsersServiceClient) GetUserPublicProfile(ctx context.Context, username string) (*client.User, error) {
	return &user, nil
}

func (m *mockUsersServiceClient) GetUserPortfolioLink(ctx context.Context, username string) (*url.URL, error) {
	link := "https://www.ogola.me"
	return url.Parse(link)
}

func (m *mockUsersServiceClient) GetUserPhotos(ctx context.Context, username string, queryParams client.QueryParams) ([]client.Photo, error) {
	return pics, nil
}

func (m *mockUsersServiceClient) GetUserLikedPhotos(ctx context.Context, username string, queryParams client.QueryParams) ([]client.Photo, error) {
	return pics, nil
}

func (m *mockUsersServiceClient) GetUserCollections(ctx context.Context, username string, queryParams client.QueryParams) ([]client.Collection, error) {
	return collections, nil
}

func (m *mockUsersServiceClient) GetUserStats(ctx context.Context, username string, queryParams client.QueryParams) (*client.UserStats, error) {
	return &client.UserStats{Username: username}, nil
}

func (m *mockUsersServiceClient) SearchUsers(ctx context.Context, queryParams client.QueryParams) (*client.UserSearchResult, error) {
	return &client.UserSearchResult{Results: users}, nil
}

func (m *mockUsersServiceClient) GetUserPrivateProfile(ctx context.Context) (*client.User, error) {
	return &user, nil
}

func (m *mockUsersServiceClient) UpdateUserProfile(ctx context.Context, data map[string]string) (*client.User, error) {
	return &user, nil
}

func TestUsersService(t *testing.T) {
	mockUnsplash := &Unsplash{
		Users: &UsersService{client: &mockUsersServiceClient{}},
	}

	t.Run("public profile", func(t *testing.T) {
		res, err := mockUnsplash.Users.PublicProfile(context.Background(), "cole")
		checkErrorIsNil(t, err)
		checkRsNotNil(t, res)

		if res != &user {
			t.Errorf("expected %v but got %v", &user, res)
		}
	})
}
