package unsplash

import (
	"context"

	"github.com/eddogola/unsplash-go/unsplash/client"
)

// CollectionsServiceClient defines client methods used to get collection resources
type CollectionsServiceClient interface {
	GetCollectionsList(context.Context, client.QueryParams) ([]client.Collection, error)
	GetCollection(context.Context, string) (*client.Collection, error)
	GetCollectionPhotos(context.Context, string, client.QueryParams) ([]client.Photo, error)
	GetRelatedCollections(context.Context, string) ([]client.Collection, error)
	SearchCollections(context.Context, client.QueryParams) (*client.CollectionSearchResult, error)
	CreateCollection(context.Context, map[string]string) (*client.Collection, error)
	UpdateCollection(context.Context, string, map[string]string) (*client.Collection, error)
	DeleteCollection(context.Context, string) error
	AddPhotoToCollection(context.Context, string, map[string]string) (*client.CollectionActionResponse, error)
	RemovePhotoFromCollection(context.Context, string, map[string]string) (*client.CollectionActionResponse, error)
}

// CollectionsService contains an underlying Unsplash client to
//be used for http methods
type CollectionsService struct {
	client CollectionsServiceClient
}

// All returns a paginated list of all Collections on Unsplash
func (cs *CollectionsService) All(queryParams client.QueryParams) ([]client.Collection, error) {
	ctx := context.Background()
	return cs.client.GetCollectionsList(ctx, queryParams)
}

// Get returns a specific Collection, given its id
func (cs *CollectionsService) Get(collectionID string) (*client.Collection, error) {
	ctx := context.Background()
	return cs.client.GetCollection(ctx, collectionID)
}

// Photos returns a paginated list of Photos under the given collection
func (cs *CollectionsService) Photos(collectionID string, queryParams client.QueryParams) ([]client.Photo, error) {
	ctx := context.Background()
	return cs.client.GetCollectionPhotos(ctx, collectionID, queryParams)
}

// Related returns a paginated list of collections related to the given collection
func (cs *CollectionsService) Related(collectionID string) ([]client.Collection, error) {
	ctx := context.Background()
	return cs.client.GetRelatedCollections(ctx, collectionID)
}

// Search takes in a search query under the given query parameters to return a list of Collection search results
func (cs *CollectionsService) Search(searchQuery string, queryParams client.QueryParams) (*client.CollectionSearchResult, error) {
	ctx := context.Background()
	if queryParams == nil || queryParams["query"] == "" {
		queryParams = make(client.QueryParams)
		queryParams["query"] = searchQuery
		return cs.client.SearchCollections(ctx, queryParams)
	}
	queryParams["query"] = searchQuery
	return cs.client.SearchCollections(ctx, queryParams)
}

// methods requiring private authentication

// Create creates a new collection using the data provided in the map, returning it if the creation process is successful
func (cs *CollectionsService) Create(data map[string]string) (*client.Collection, error) {
	ctx := context.Background()
	return cs.client.CreateCollection(ctx, data)
}

// Update uses the data provided in the map to update the given collection
// returning the updated collection
func (cs *CollectionsService) Update(collectionID string, data map[string]string) (*client.Collection, error) {
	ctx := context.Background()
	return cs.client.UpdateCollection(ctx, collectionID, data)
}

// Delete removes the given collection
func (cs *CollectionsService) Delete(collectionID string) error {
	ctx := context.Background()
	return cs.client.DeleteCollection(ctx, collectionID)
}

// AddPhoto takes in a `photo_id` in the data map, to add the Photo to the given collection
func (cs *CollectionsService) AddPhoto(collectionID string, data map[string]string) (*client.CollectionActionResponse, error) {
	ctx := context.Background()
	return cs.client.AddPhotoToCollection(ctx, collectionID, data)
}

// RemovePhoto takes in a `photo_id` in the data map, to remove the Photo from the given collection
func (cs *CollectionsService) RemovePhoto(collectionID string, data map[string]string) (*client.CollectionActionResponse, error) {
	ctx := context.Background()
	return cs.client.RemovePhotoFromCollection(ctx, collectionID, data)
}
