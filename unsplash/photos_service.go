package unsplash

import (
	"context"

	"github.com/eddogola/unsplash-go/unsplash/client"
)

// PhotosServiceClient defines client methods used to get or work
// with photo resources
type PhotosServiceClient interface {
	GetPhotoList(context.Context, client.QueryParams) ([]client.Photo, error)
	GetPhoto(context.Context, string) (*client.Photo, error)
	GetRandomPhoto(context.Context, client.QueryParams) (interface{}, error)
	GetPhotoStats(context.Context, string, client.QueryParams) (*client.PhotoStats, error)
	SearchPhotos(context.Context, client.QueryParams) (*client.PhotoSearchResult, error)
	UpdatePhoto(context.Context, string, map[string]string) (*client.Photo, error)
	LikePhoto(context.Context, string) (*client.LikeResponse, error)
	UnlikePhoto(context.Context, string) error
}

// PhotosService contains an underlying Unsplash client to
//be used for http methods
type PhotosService struct {
	client PhotosServiceClient
}

// All returns a paginated list of all the Photos on Unsplash
func (ps *PhotosService) All(ctx context.Context, queryParams client.QueryParams) ([]client.Photo, error) {
	return ps.client.GetPhotoList(ctx, queryParams)
}

// Get returns a single Photo, requested using the photo's ID
func (ps *PhotosService) Get(ctx context.Context, photoID string) (*client.Photo, error) {
	return ps.client.GetPhoto(ctx, photoID)
}

// GetRandom returns a random Photo.
// Returns a paginated list of Photos if `count` query parameter is provided in the query parameters.
func (ps *PhotosService) GetRandom(ctx context.Context, queryParams client.QueryParams) (interface{}, error) {
	return ps.client.GetRandomPhoto(ctx, queryParams)
}

// GetStats returns the requested Photo's Stats
func (ps *PhotosService) GetStats(ctx context.Context, photoID string, queryParams client.QueryParams) (*client.PhotoStats, error) {
	return ps.client.GetPhotoStats(ctx, photoID, queryParams)
}

// Search takes in a search query in the query parameters and returns a list of Photo search results
func (ps *PhotosService) Search(ctx context.Context, searchQuery string, queryParams client.QueryParams) (*client.PhotoSearchResult, error) {
	if queryParams == nil || queryParams["query"] == "" {
		queryParams = make(client.QueryParams)
		queryParams["query"] = searchQuery
		return ps.client.SearchPhotos(ctx, queryParams)
	}
	queryParams["query"] = searchQuery
	return ps.client.SearchPhotos(ctx, queryParams)
}

// methods requiring private authentication

// Update uses the data provided to update info on the requested Photo
func (ps *PhotosService) Update(ctx context.Context, photoID string, updatedData map[string]string) (*client.Photo, error) {
	return ps.client.UpdatePhoto(ctx, photoID, updatedData)
}

// Like adds a like on the photo whose photo ID is provided on behalf of the authenticated user
func (ps *PhotosService) Like(ctx context.Context, photoID string) (*client.LikeResponse, error) {
	return ps.client.LikePhoto(ctx, photoID)
}

// Unlike removes a like on the photo whose photo ID is provided on behalf of the authenticated user
func (ps *PhotosService) Unlike(ctx context.Context, photoID string) error {
	return ps.client.UnlikePhoto(ctx, photoID)
}
