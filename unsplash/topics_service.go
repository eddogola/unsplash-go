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
func (ts *TopicsService) All(ctx context.Context, queryParams client.QueryParams) ([]client.Topic, error) {
	return ts.client.GetTopicsList(ctx, queryParams)
}

// Get returns a specific Topic, using the topic's ID or slug
func (ts *TopicsService) Get(ctx context.Context, topicIDOrSlug string) (*client.Topic, error) {
	return ts.client.GetTopic(ctx, topicIDOrSlug)
}

// Photos returns a paginated list of Photos under the Topic requested using the
// topic's ID or slug
func (ts *TopicsService) Photos(ctx context.Context, topicIDOrSlug string, queryParams client.QueryParams) ([]client.Photo, error) {
	return ts.client.GetTopicPhotos(ctx, topicIDOrSlug, queryParams)
}
