package unsplash

import (
	"context"
	"fmt"
)

// Collection defines fields in a collection resource
type Collection struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	PublishedAt     string `json:"published_at"`
	LastCollectedAt string `json:"last_collected_at"`
	UpdatedAt       string `json:"updated_at"`
	Featured        bool   `json:"featured"`
	TotalPhotos     int    `json:"total_photos"`
	Private         bool   `json:"private"`
	ShareKey        string `json:"share_key"`
	CoverPhoto      Photo  `json:"cover_photo"`
	User            User   `json:"user"`
	Links           struct {
		Self   string `json:"self"`
		HTML   string `json:"html"`
		Photos string `json:"photos"`
	} `json:"links"`
}

// CollectionActionResponse defines the fields returned on adding a photo to a collection
// or deleting a photo from a collection.
type CollectionActionResponse struct {
	Photo      Photo      `json:"photo"`
	Collection Collection `json:"collection"`
	User       User       `json:"user"`
	CreatedAt  string     `json:"created_at"`
}

// gets a single page of a list of collections
// https://unsplash.com/documentation#list-collections
func (c *Client) getCollectionsList(ctx context.Context, queryParams QueryParams) ([]Collection, error) {
	link, err := buildURL(collectionsListEndpoint, queryParams)
	if err != nil {
		return nil, err
	}
	data, err := c.getBodyBytes(ctx, link)
	if err != nil {
		return nil, err
	}
	var collections []Collection
	err = parseJSON(data, collections)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

// get a collection using id
// https://unsplash.com/documentation#get-a-collection
func (c *Client) getCollection(ctx context.Context, id int) (*Collection, error) {
	endPoint := collectionsListEndpoint + fmt.Sprint(id)
	data, err := c.getBodyBytes(ctx, endPoint)
	if err != nil {
		return nil, err
	}
	var collection Collection
	err = parseJSON(data, &collection)
	if err != nil {
		return nil, err
	}
	return &collection, nil
}

// Retrieve a collection's photos
// https://unsplash.com/documentation#get-a-collections-photos
func (c *Client) getCollectionPhotos(ctx context.Context, id int, queryParams QueryParams) ([]Photo, error) {
	endPoint := collectionsListEndpoint + fmt.Sprint(id) + "/photos"
	link, err := buildURL(endPoint, queryParams)
	if err != nil {
		return nil, err
	}
	data, err := c.getBodyBytes(ctx, link)
	if err != nil {
		return nil, err
	}
	var pics []Photo
	err = parseJSON(data, pics)
	if err != nil {
		return nil, err
	}
	return pics, nil
}

// Retrieve a list of collections related to this one.
// https://unsplash.com/documentation#list-a-collections-related-collections
func (c *Client) getRelatedCollections(ctx context.Context, id int) ([]Collection, error) {
	endPoint := collectionsListEndpoint + fmt.Sprint(id) + "/related"
	data, err := c.getBodyBytes(ctx, endPoint)
	if err != nil {
		return nil, err
	}
	var collections []Collection
	err = parseJSON(data, collections)
	if err != nil {
		return nil, err
	}
	return collections, nil
}
