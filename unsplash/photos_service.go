package unsplash

import (
	"context"

	"github.com/eddogola/unsplash-go/unsplash/client"
)

type PhotosService struct {
	client *client.Client
}

func (ps *PhotosService) GetAll(ctx context.Context, queryParams client.QueryParams) ([]client.Photo, error) {
	return ps.client.GetPhotoList(ctx, queryParams)
}

func (ps *PhotosService) Get(ctx context.Context, ID string) (*client.Photo, error) {
	return ps.client.GetPhoto(ctx, ID)
}

func (ps *PhotosService) GetRandom(ctx context.Context, queryParams client.QueryParams) (interface{}, error) {
	return ps.client.GetRandomPhoto(ctx, queryParams)
}

func (ps *PhotosService) GetStats(ctx context.Context, ID string, queryParams client.QueryParams) (*client.PhotoStats, error) {
	return ps.client.GetPhotoStats(ctx, ID, queryParams)
}

func (ps *PhotosService) Search(ctx context.Context, queryParams client.QueryParams) (*client.PhotoSearchResult, error) {
	return ps.client.SearchPhotos(ctx, queryParams)
}

// methods requiring private authentication
func (ps *PhotosService) Update(ctx context.Context, photoID string, updatedData map[string]string) (*client.Photo, error) {
	return ps.client.UpdatePhoto(ctx, photoID, updatedData)
}

func (ps *PhotosService) Like(ctx context.Context, photoID string) (*client.LikeResponse, error) {
	return ps.client.LikePhoto(ctx, photoID)
}

func (ps *PhotosService) Unlike(ctx context.Context, photoID string) error {
	return ps.client.UnlikePhoto(ctx, photoID)
}
