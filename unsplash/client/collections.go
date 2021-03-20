package client

import (
	"context"
	"fmt"
)

// Collection defines fields in a collection resource
type Collection struct {
	ID              string `json:"id"`
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

// GetCollectionsList takes in a context and query parameters to build the required response
// and return a slice of Collection objects.
// Gets a single page of a list of collections
// https://unsplash.com/documentation#list-collections
func (c *Client) GetCollectionsList(ctx context.Context, queryParams QueryParams) ([]Collection, error) {
	link, err := buildURL(CollectionsListEndpoint, queryParams)
	if err != nil {
		return nil, err
	}
	data, err := c.getBodyBytes(ctx, link)
	if err != nil {
		return nil, err
	}
	var collections []Collection
	err = parseJSON(data, &collections)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

// GetCollection takes in a contect and a collection id to return a single *Collection
// if a collection of the given id is found.
// Get a collection using id
// https://unsplash.com/documentation#get-a-collection
func (c *Client) GetCollection(ctx context.Context, id string) (*Collection, error) {
	endPoint := CollectionsListEndpoint + fmt.Sprint(id)
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

// GetCollectionPhotos takes in a context, collection id, and query parameters.
// If a collection of the given id is found, a slice of Photo objects in the collection
// is returned.
// Retrieve a collection's photos
// https://unsplash.com/documentation#get-a-collections-photos
func (c *Client) GetCollectionPhotos(ctx context.Context, id string, queryParams QueryParams) ([]Photo, error) {
	endPoint := CollectionsListEndpoint + fmt.Sprint(id) + "/photos"
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
	return pics, nil
}

// GetRelatedCollections takes in a context and a collection id to return a slice of Collection
// objects if the collection of the given id is found.
// Retrieve a list of collections related to this one.
// https://unsplash.com/documentation#list-a-collections-related-collections
func (c *Client) GetRelatedCollections(ctx context.Context, id string) ([]Collection, error) {
	endPoint := CollectionsListEndpoint + fmt.Sprint(id) + "/related"
	data, err := c.getBodyBytes(ctx, endPoint)
	if err != nil {
		return nil, err
	}
	var collections []Collection
	err = parseJSON(data, &collections)
	if err != nil {
		return nil, err
	}
	return collections, nil
}
