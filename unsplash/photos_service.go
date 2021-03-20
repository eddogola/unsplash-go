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
func (ps *PhotosService) All(queryParams client.QueryParams) ([]client.Photo, error) {
	ctx := context.Background()
	return ps.client.GetPhotoList(ctx, queryParams)
}

// Get returns a single Photo, requested using the photo's ID
func (ps *PhotosService) Get(photoID string) (*client.Photo, error) {
	ctx := context.Background()
	return ps.client.GetPhoto(ctx, photoID)
}

// Random returns a random Photo.
// Returns a paginated list of Photos if `count` query parameter is provided in the query parameters.
func (ps *PhotosService) Random(queryParams client.QueryParams) (interface{}, error) {
	ctx := context.Background()
	return ps.client.GetRandomPhoto(ctx, queryParams)
}

// Stats returns the requested Photo's Stats
func (ps *PhotosService) Stats(photoID string, queryParams client.QueryParams) (*client.PhotoStats, error) {
	ctx := context.Background()
	return ps.client.GetPhotoStats(ctx, photoID, queryParams)
}

// Search takes in a search query in the query parameters and returns a list of Photo search results
func (ps *PhotosService) Search(searchQuery string, queryParams client.QueryParams) (*client.PhotoSearchResult, error) {
	ctx := context.Background()
	if queryParams == nil {
		queryParams = make(client.QueryParams)
		queryParams["query"] = searchQuery
		return ps.client.SearchPhotos(ctx, queryParams)
	}
	queryParams["query"] = searchQuery
	return ps.client.SearchPhotos(ctx, queryParams)
}

// methods requiring private authentication

// Update uses the data provided to update info on the requested Photo
func (ps *PhotosService) Update(photoID string, updatedData map[string]string) (*client.Photo, error) {
	ctx := context.Background()
	return ps.client.UpdatePhoto(ctx, photoID, updatedData)
}

// Like adds a like on the photo whose photo ID is provided on behalf of the authenticated user
func (ps *PhotosService) Like(photoID string) (*client.LikeResponse, error) {
	ctx := context.Background()
	return ps.client.LikePhoto(ctx, photoID)
}

// Unlike removes a like on the photo whose photo ID is provided on behalf of the authenticated user
func (ps *PhotosService) Unlike(photoID string) error {
	ctx := context.Background()
	return ps.client.UnlikePhoto(ctx, photoID)
}
