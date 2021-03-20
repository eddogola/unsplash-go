package unsplash

import (
	"context"

	"github.com/eddogola/unsplash-go/unsplash/client"
)

// TopicsServiceClient defines client methods used to get topic resources
type TopicsServiceClient interface {
	GetTopicsList(context.Context, client.QueryParams) ([]client.Topic, error)
	GetTopic(context.Context, string) (*client.Topic, error)
	GetTopicPhotos(context.Context, string, client.QueryParams) ([]client.Photo, error)
}

// TopicsService contains an underlying Unsplash client to
//be used for http methods
type TopicsService struct {
	client TopicsServiceClient
}

// All returns a paginated list of all Topics on unsplash
func (ts *TopicsService) All(queryParams client.QueryParams) ([]client.Topic, error) {
	ctx := context.Background()
	return ts.client.GetTopicsList(ctx, queryParams)
}

// Get returns a specific Topic, using the topic's ID or slug
func (ts *TopicsService) Get(topicIDOrSlug string) (*client.Topic, error) {
	ctx := context.Background()
	return ts.client.GetTopic(ctx, topicIDOrSlug)
}

// Photos returns a paginated list of Photos under the Topic requested using the
// topic's ID or slug
func (ts *TopicsService) Photos(topicIDOrSlug string, queryParams client.QueryParams) ([]client.Photo, error) {
	ctx := context.Background()
	return ts.client.GetTopicPhotos(ctx, topicIDOrSlug, queryParams)
}
