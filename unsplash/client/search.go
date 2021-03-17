package client

import (
	"context"
)

// PhotoSearchResult defines the structure of the response gotten
// after searching for a picture
type PhotoSearchResult struct {
	Total      int     `json:"total"`
	TotalPages int     `json:"total_pages"`
	Results    []Photo `json:"results"`
}

// CollectionSearchResult defines the structure of the response gotten
// after searching for a collection
type CollectionSearchResult struct {
	Total      int          `json:"total"`
	TotalPages int          `json:"total_pages"`
	Results    []Collection `json:"results"`
}

// UserSearchResult defines the structure of the response gotten
// after searching for a user
type UserSearchResult struct {
	Total      int    `json:"total"`
	TotalPages int    `json:"total_pages"`
	Results    []User `json:"results"`
}

// SearchPhotos takes in a context and query parameters, returns the search results,
// in a pointer to a PhotoSearchResult object.
// Get a single page with photo search results
// https://unsplash.com/documentation#search-photos
func (c *Client) SearchPhotos(ctx context.Context, queryParams QueryParams) (*PhotoSearchResult, error) {
	link, err := buildURL(SearchPhotosEndpoint, queryParams)
	if err != nil {
		return nil, err
	}
	// throw an error if search query parameter not in URL
	if _, ok := queryParams["query"]; !ok {
		return nil, errQueryNotInURL(link)
	}
	data, err := c.getBodyBytes(ctx, link)
	if err != nil {
		return nil, err
	}
	var res PhotoSearchResult
	err = parseJSON(data, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// SearchCollections takes in a context and query parameters, returns the search results,
// in a pointer to a CollectionSearchResult object.
// Get a single page with collection search results
// https://unsplash.com/documentation#search-collections
func (c *Client) SearchCollections(ctx context.Context, queryParams QueryParams) (*CollectionSearchResult, error) {
	link, err := buildURL(SearchCollectionsEndpoint, queryParams)
	if err != nil {
		return nil, err
	}
	// throw an error if search query parameter not in URL
	if _, ok := queryParams["query"]; !ok {
		return nil, errQueryNotInURL(link)
	}
	data, err := c.getBodyBytes(ctx, link)
	if err != nil {
		return nil, err
	}
	var res CollectionSearchResult
	err = parseJSON(data, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// SearchUsers takes in a context and query parameters, returns the search results,
// in a pointer to a UserSearchResult object.
// Get a single page with users search results
// https://unsplash.com/documentation#search-users
func (c *Client) SearchUsers(ctx context.Context, queryParams QueryParams) (*UserSearchResult, error) {
	link, err := buildURL(SearchUsersEndpoint, queryParams)
	if err != nil {
		return nil, err
	}
	// throw an error if search query parameter not in URL
	if _, ok := queryParams["query"]; !ok {
		return nil, errQueryNotInURL(link)
	}
	data, err := c.getBodyBytes(ctx, link)
	if err != nil {
		return nil, err
	}
	var res UserSearchResult
	err = parseJSON(data, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
