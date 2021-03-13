package client

import "context"

// Topic defines fields in an Unsplash topic
type Topic struct {
	ID          string `json:"id"`
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PublishedAt string `json:"published_at"`
	UpdatedAt   string `json:"updated_at"`
	StartsAt    string `json:"starts_at"`
	EndsAt      string `json:"ends_at"`
	Featured    bool   `json:"featured"`
	TotalPhotos int    `json:"total_photos"`
	Links       struct {
		Self   string `json:"self"`
		HTML   string `json:"html"`
		Photos string `json:"photos"`
	} `json:"links"`
	Status                      string  `json:"status"`
	Owners                      []User  `json:"owners"`
	TopContributors             []User  `json:"top_contributors"`
	CurrentUserContributions    []Photo `json:"current_user_contributions"`
	TotalCurrentUserSubmissions int     `json:"total_current_user_submissions"`
	CoverPhoto                  Photo   `json:"cover_photo"`
	PreviewPhotos               []Photo `json:"preview_photos"`
}

// GetTopicsList takes a context and query parameters, returns a slice of Topic objects.
// Get a single page from the list of all topics.
// https://unsplash.com/documentation#list-topics
func (c *Client) GetTopicsList(ctx context.Context, queryParams QueryParams) ([]Topic, error) {
	link, err := buildURL(TopicsListEndpoint, queryParams)
	if err != nil {
		return nil, err
	}
	data, err := c.getBodyBytes(ctx, link)
	if err != nil {
		return nil, err
	}
	var topics []Topic
	err = parseJSON(data, topics)
	if err != nil {
		return nil, err
	}
	return topics, nil
}

// GetTopic takes a context and a topic's id or slug. Returns a pointer to a Topic,
// if a topic with the given id or slug is found.
// Retrieve a single topic.
// https://unsplash.com/documentation#get-a-topic
func (c *Client) GetTopic(ctx context.Context, IDOrSlug string) (*Topic, error) {
	endPoint := TopicsListEndpoint + IDOrSlug
	data, err := c.getBodyBytes(ctx, endPoint)
	if err != nil {
		return nil, err
	}
	var topic Topic
	err = parseJSON(data, &topic)
	if err != nil {
		return nil, err
	}
	return &topic, nil
}

// GetTopicPhotos takes a context, a topic's id or slug, and query parameters. Returns slice of Photos
// in the topic if topic with the given id or slug is found.
// Retrieve a topic’s photos.
// https://unsplash.com/documentation#get-a-topics-photos
func (c *Client) GetTopicPhotos(ctx context.Context, IDOrSlug string, queryParams QueryParams) ([]Photo, error) {
	endPoint := TopicsListEndpoint + IDOrSlug + "/photos"
	link, err := buildURL(endPoint, queryParams)
	if err != nil {
		return nil, err
	}
	data, err := c.getBodyBytes(ctx, link)
	if err != nil {
		return nil, err
	}
	var pics []Photo
	err = parseJSON(data, &pics)
	if err != nil {
		return nil, err
	}
	return pics, err
}
