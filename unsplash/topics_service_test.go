package unsplash

import (
	"context"
	"reflect"
	"testing"

	"github.com/eddogola/unsplash-go/unsplash/client"
)

var topics = []client.Topic{
	{
		ID: "cole",
		Title: "king Cole",
		CoverPhoto: pic,
	},
	{
		ID: "kendrick",
		Title: "king Kunta",
		CoverPhoto: pic,
	},
}

var topic = topics[0]

type mockTopicsServiceClient struct{}

func (m *mockTopicsServiceClient) GetTopicsList(ctx context.Context, queryParams client.QueryParams) ([]client.Topic, error) {
	return topics, nil
}

func (m *mockTopicsServiceClient) GetTopic(ctx context.Context, topicIDOrSlug string) (*client.Topic, error) {
	topic.ID = topicIDOrSlug
	topic.Slug = topicIDOrSlug
	return &topic, nil
}

func (m *mockTopicsServiceClient) GetTopicPhotos(ctx context.Context, topicIDOrSlug string, queryParams client.QueryParams) ([]client.Photo, error) {
	return pics, nil
}

func TestTopicsService(t *testing.T) {
	mockUnsplash := &Unsplash{
		Topics: &TopicsService{client: &mockTopicsServiceClient{}},
	}

	t.Run("all topics", func(t *testing.T) {
		res, err := mockUnsplash.Topics.All(context.Background(), nil)
		checkErrorIsNil(t, err)
		checkRsNotNil(t, res)

		if !reflect.DeepEqual(res, topics) {
			t.Errorf("expected %v but got %v", topics, res)
		}
	})

	t.Run("get specific topic", func(t *testing.T) {
		res, err := mockUnsplash.Topics.Get(context.Background(), "kingCole")
		checkErrorIsNil(t, err)
		checkRsNotNil(t, res)

		if res != &topic {
			t.Errorf("expected %v but got %v", topic, res)
		}
	})
}