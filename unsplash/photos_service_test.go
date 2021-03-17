package unsplash

import (
	"context"
	"testing"

	"github.com/eddogola/unsplash-go/unsplash/client"
)

var pics = []client.Photo{
	{
		ID: "whackID",
		Width: 1308,
		Height: 768,
		Downloads: 400,
		Likes: 799,
	},
	{
		ID: "noID",
		Width: 450,
		Height: 600,
	},
}

var pic = client.Photo{
	ID: "whackID",
	Width: 1308,
	Height: 768,
	Downloads: 400,
	Likes: 799,
}

type mockPhotoServiceClient struct{}

func (m *mockPhotoServiceClient) GetPhotoList(ctx context.Context, queryParams client.QueryParams) ([]client.Photo, error) {
	return pics, nil
}
func (m *mockPhotoServiceClient) GetPhoto(ctx context.Context, photoID string) (*client.Photo, error) {
	return &pic, nil
}
func (m *mockPhotoServiceClient) GetRandomPhoto(ctx context.Context, queryParams client.QueryParams) (interface{}, error) {
	if _, ok := queryParams["count"]; ok {
		return pics, nil
	}
	return &pic, nil
}
func (m *mockPhotoServiceClient) GetPhotoStats(ctx context.Context,photoID  string, queryParams client.QueryParams) (*client.PhotoStats, error) {
	return &client.PhotoStats{ID: "ladida"}, nil
}
func (m *mockPhotoServiceClient) SearchPhotos(ctx context.Context, queryParams client.QueryParams) (*client.PhotoSearchResult, error){
	return &client.PhotoSearchResult{Results: pics}, nil
}
func (m *mockPhotoServiceClient) UpdatePhoto(ctx context.Context, photoID string, updatedData map[string]string) (*client.Photo, error){
	return &pic, nil
}
func (m *mockPhotoServiceClient) LikePhoto(ctx context.Context, photoID string) (*client.LikeResponse, error){
	return &client.LikeResponse{Photo: pic}, nil
}
func (m *mockPhotoServiceClient) UnlikePhoto(ctx context.Context, photoID string) (error){
	return nil
}

func TestPhotosService(t *testing.T) {
	mockPhotoService := &PhotosService{client: &mockPhotoServiceClient{}}
	unsplash := &Unsplash{
		Photos: mockPhotoService,
	}

	t.Run("random photo when count not passed", func(t *testing.T) {
		res, err := unsplash.Photos.GetRandom(context.Background(), nil)
		randomPhoto := res.(*client.Photo)
		checkErrorIsNil(t, err)
		checkRsNotNil(t, randomPhoto)
	})

	t.Run("random photo when count passed", func(t *testing.T) {
		res, err := unsplash.Photos.GetRandom(context.Background(), client.QueryParams{"count": "1"})
		randomPhotos := res.([]client.Photo)
		checkErrorIsNil(t, err)
		checkRsNotNil(t, randomPhotos)
	})
}

func checkErrorIsNil(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

// raise error if resource is nil
func checkRsNotNil(t *testing.T, rs interface{}) {
	t.Helper()
	if rs == nil {
		t.Errorf("resource gotten is nil: %v", rs)
	}
}
