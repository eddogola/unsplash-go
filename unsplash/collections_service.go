package unsplash

import (
	"context"

	"github.com/eddogola/unsplash-go/unsplash/client"
)

type CollectionsService struct {
	client *client.Client
}

func (cs *CollectionsService) All(ctx context.Context, queryParams client.QueryParams) ([]client.Collection, error) {
	return cs.client.GetCollectionsList(ctx, queryParams)
}

func (cs *CollectionsService) Get(ctx context.Context, id int) (*client.Collection, error) {
	return cs.client.GetCollection(ctx, id)
}

func (cs *CollectionsService) Photos(ctx context.Context, id int, queryParams client.QueryParams) ([]client.Photo, error) {
	return cs.client.GetCollectionPhotos(ctx, id, queryParams)
}

func (cs *CollectionsService) Related(ctx context.Context, id int) ([]client.Collection, error) {
	return cs.client.GetRelatedCollections(ctx, id)
}

func (cs *CollectionsService) Search(ctx context.Context, queryParams client.QueryParams) (*client.CollectionSearchResult, error) {
	return cs.client.SearchCollections(ctx, queryParams)
}

// methods requiring private authentication
func (cs *CollectionsService) Create(ctx context.Context, data map[string]string) (*client.Collection, error) {
	return cs.client.CreateCollection(ctx, data)
}

func (cs *CollectionsService) Update(ctx context.Context, collectionID string, data map[string]string) (*client.Collection, error) {
	return cs.client.UpdateCollection(ctx, collectionID, data)
}

func (cs *CollectionsService) Delete(ctx context.Context, collectionID string) error {
	return cs.client.DeleteCollection(ctx, collectionID)
}

func (cs *CollectionsService) AddPhoto(ctx context.Context, collectionID string, data map[string]string) (*client.CollectionActionResponse, error) {
	return cs.client.AddPhotoToCollection(ctx, collectionID, data)
}

func (cs *CollectionsService) RemovePhoto(ctx context.Context, collectionID string, data map[string]string) (*client.CollectionActionResponse, error) {
	return cs.client.RemovePhotoFromCollection(ctx, collectionID, data)
}