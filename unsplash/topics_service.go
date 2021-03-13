package unsplash

import (
	"context"

	"github.com/eddogola/unsplash-go/unsplash/client"
)

type TopicsService struct {
	client *client.Client
}

func (ts *TopicsService) All(ctx context.Context, queryParams client.QueryParams) ([]client.Topic, error) {
	return ts.client.GetTopicsList(ctx, queryParams)
}

func (ts *TopicsService) Get(ctx context.Context, IDOrSlug string) (*client.Topic, error) {
	return ts.client.GetTopic(ctx, IDOrSlug)
}

func (ts *TopicsService) Photos(ctx context.Context, IDOrSlug string, queryParams client.QueryParams) ([]client.Photo, error) {
	return ts.client.GetTopicPhotos(ctx, IDOrSlug, queryParams)
}