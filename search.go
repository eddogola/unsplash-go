package unsplash

import "context"

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

// get a single page with photo search results
// https://unsplash.com/documentation#search-photos
func (c *Client) searchPhotos(ctx context.Context, queryParams QueryParams) (*PhotoSearchResult, error) {
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

// get a single page with collection search results
// https://unsplash.com/documentation#search-collections
func (c *Client) searchCollections(ctx context.Context, queryParams QueryParams) (*CollectionSearchResult, error) {
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

// get a single page with users search results
// https://unsplash.com/documentation#search-users
func (c *Client) searchUsers(ctx context.Context, queryParams QueryParams) (*UserSearchResult, error) {
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
