package unsplash

import (
	"context"
	"reflect"
	"testing"

	"github.com/eddogola/unsplash-go/unsplash/client"
)

var collections = []client.Collection{
	{
		ID: 1,
		Title: "J Cole Photoshoot",
		TotalPhotos: 60,
		CoverPhoto: pic, // `pic` from photo_service_test.go
	},
	{
		ID: 11,
		Title: "Kendrick Photoshoot",
		TotalPhotos: 89,
		CoverPhoto: pic, // `pic` from photo_service_test.go
	},
}

var collection = collections[0]

type mockCollectionsServiceClient struct{}

func (m *mockCollectionsServiceClient) GetCollectionsList(ctx context.Context, queryParams client.QueryParams) ([]client.Collection, error) {
	return collections, nil
}

func (m *mockCollectionsServiceClient) GetCollection(ctx context.Context, collectionID int) (*client.Collection, error) {
	collection.ID = collectionID
	return &collection, nil
}

func (m *mockCollectionsServiceClient) GetCollectionPhotos(ctx context.Context, collectionID int, queryParams client.QueryParams) ([]client.Photo, error) {
	return pics, nil // get `pics` from photo_service_test.go
}

func (m *mockCollectionsServiceClient) GetRelatedCollections(ctx context.Context, collectionID int) ([]client.Collection, error) {
	return collections, nil
}

func (m *mockCollectionsServiceClient) SearchCollections(ctx context.Context, queryParams client.QueryParams) (*client.CollectionSearchResult, error) {
	return &client.CollectionSearchResult{Results: collections}, nil
}

func (m *mockCollectionsServiceClient) CreateCollection(ctx context.Context, data map[string]string) (*client.Collection, error) {
	return &collection, nil
}

func (m *mockCollectionsServiceClient) UpdateCollection(ctx context.Context, collectionID string, data map[string]string) (*client.Collection, error) {
	return &collection, nil
}

func (m *mockCollectionsServiceClient) DeleteCollection(ctx context.Context, collectionID string) error {
	return nil
}

func (m *mockCollectionsServiceClient) AddPhotoToCollection(ctx context.Context, collectionID string, data map[string]string) (*client.CollectionActionResponse, error) {
	return &client.CollectionActionResponse{Photo: pic, Collection: collection}, nil // `pic` from photo_service_test.go
}

func (m *mockCollectionsServiceClient) RemovePhotoFromCollection(ctx context.Context, collectionID string, data map[string]string) (*client.CollectionActionResponse, error) {
	return &client.CollectionActionResponse{Photo: pic, Collection: collection}, nil // `pic` from photo_service_test.go
}

func TestCollectionsService(t *testing.T) {
	mockUnsplash := &Unsplash{
		Collections: &CollectionsService{client: &mockCollectionsServiceClient{}},
	}

	t.Run("all collections", func(t *testing.T) {
		res, err := mockUnsplash.Collections.All(context.Background(), nil)
		checkErrorIsNil(t, err)
		checkRsNotNil(t, res)

		if !reflect.DeepEqual(res, collections) {
			t.Errorf("expected %v but got %v", collections, res)
		}
	})
}
