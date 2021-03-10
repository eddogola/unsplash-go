package unsplash

import (
	"context"
	"fmt"
)

// user functions

// Retrieve public details on a given user.
// https://unsplash.com/documentation#get-a-users-public-profile
func (c *Client) getUserPublicProfile(ctx context.Context, username string) (*User, error) {
	endPoint := baseUserEndpoint + username
	data, err := c.getBodyBytes(ctx, endPoint)
	if err != nil {
		return nil, err
	}
	var user User
	err = parseJSON(data, user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Retrieve a single user’s portfolio link.
// https://unsplash.com/documentation#get-a-users-portfolio-link
func (c *Client) getUserPortfolioLink(ctx context.Context, username string) (string, error) {
	endPoint := baseUserEndpoint + username + "/portfolio"
	data, err := c.getBodyBytes(ctx, endPoint)
	if err != nil {
		return "", err
	}

	// parse json response data using local struct since it has only onee field
	var resp struct {
		Link string `json:"url"`
	}
	err = parseJSON(data, &resp)
	if err != nil {
		return "", err
	}
	return resp.Link, nil
}

// Get a list of photos uploaded by a user.
// https://unsplash.com/documentation#list-a-users-photos
func (c *Client) getUserPhotos(ctx context.Context, username string, queryParams QueryParams) ([]Photo, error) {
	endPoint := baseUserEndpoint + username
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

// Get a list of photos liked by a user.
// https://unsplash.com/documentation#list-a-users-liked-photos
func (c *Client) getUserLikedPhotos(ctx context.Context, username string, queryParams QueryParams) ([]Photo, error) {
	endPoint := baseUserEndpoint + username + "/likes"
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

// Get a list of collections created by the user.
// https://unsplash.com/documentation#list-a-users-collections
func (c *Client) getUserCollections(ctx context.Context, username string, queryParams QueryParams) ([]Collection, error) {
	endPoint := baseUserEndpoint + username + "/collections"
	link, err := buildURL(endPoint, queryParams)
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

// Retrieve the consolidated number of downloads, views and likes of all user’s photos,
// as well as the historical breakdown and average of these stats in a specific timeframe (default is 30 days).
// https://unsplash.com/documentation#get-a-users-statistics
func (c *Client) getUserStats(ctx context.Context, username string, queryParams QueryParams) (*UserStats, error) {
	endPoint := baseUserEndpoint + username + "/statistics"
	link, err := buildURL(endPoint, queryParams)
	if err != nil {
		return nil, err
	}
	data, err := c.getBodyBytes(ctx, link)
	if err != nil {
		return nil, err
	}
	var us UserStats
	err = parseJSON(data, &us)
	if err != nil {
		return nil, err
	}
	return &us, nil
}

// photo functions

// get a single page with a list of all photos
// https://unsplash.com/documentation#list-photos
func (c *Client) getPhotoList(ctx context.Context, queryParams QueryParams) ([]Photo, error) {
	link, err := buildURL(allPhotosEndpoint, queryParams)
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

// get a Photo using photo ID
// https://unsplash.com/documentation#get-a-photo
func (c *Client) getPhoto(ctx context.Context, ID string) (*Photo, error) {
	link := allPhotosEndpoint + ID
	data, err := c.getBodyBytes(ctx, link)
	if err != nil {
		return nil, err
	}
	var pic Photo
	err = parseJSON(data, &pic)
	if err != nil {
		return nil, err
	}
	return &pic, nil
}

// get a random Photo
// return a list of photos if a count query parameter is provided
// https://unsplash.com/documentation#get-a-random-photo
func (c *Client) getRandomPhoto(ctx context.Context, queryParams QueryParams) (interface{}, error) {
	link, err := buildURL(randomPhotoEndpoint, queryParams)
	if err != nil {
		return nil, err
	}
	data, err := c.getBodyBytes(ctx, link)
	if err != nil {
		return nil, err
	}

	/* parse data to Photo object
	or []Photo is count query parameter is present

	From API documentation:
	Note: When supplying a count parameter - and only then -
	the response will be an array of photos, even if the value of count is 1.
	*/
	if _, ok := queryParams["count"]; ok {
		var pics []Photo
		err = parseJSON(data, pics)
		if err != nil {
			return nil, err
		}
		return &pics, nil
	}

	var pic Photo
	err = parseJSON(data, pic)
	if err != nil {
		return nil, err
	}
	return &pic, nil
}

// Retrieve total number of downloads, views and likes of a single photo, as well as the historical
// breakdown of these stats in a specific timeframe (default is 30 days).
// https://unsplash.com/documentation#get-a-photos-statistics
func (c *Client) getPhotoStats(ctx context.Context, ID string, queryParams QueryParams) (*PhotoStats, error) {
	endPoint := allPhotosEndpoint + ID + "/statistics"
	link, err := buildURL(endPoint, queryParams)
	if err != nil {
		return nil, err
	}
	data, err := c.getBodyBytes(ctx, link)
	if err != nil {
		return nil, err
	}
	var ps PhotoStats
	err = parseJSON(data, &ps)
	if err != nil {
		return nil, err
	}
	return &ps, nil
}

// search functions

// get a single page with photo search results
// https://unsplash.com/documentation#search-photos
func (c *Client) searchPhotos(ctx context.Context, queryParams QueryParams) (*PhotoSearchResult, error) {
	link, err := buildURL(searchPhotosEndpoint, queryParams)
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
	link, err := buildURL(searchCollectionsEndpoint, queryParams)
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
	link, err := buildURL(searchUsersEndpoint, queryParams)
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

// collections functions

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

// topics functions

// Get a single page from the list of all topics.
// https://unsplash.com/documentation#list-topics
func (c *Client) getTopicsList(ctx context.Context, queryParams QueryParams) ([]Topic, error) {
	link, err := buildURL(topicsListEndpoint, queryParams)
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

// Retrieve a single topic.
// https://unsplash.com/documentation#get-a-topic
func (c *Client) getTopic(ctx context.Context, IDOrSlug string) (*Topic, error) {
	endPoint := topicsListEndpoint + IDOrSlug
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

// Retrieve a topic’s photos.
// https://unsplash.com/documentation#get-a-topics-photos
func (c *Client) getTopicPhotos(ctx context.Context, IDOrSlug string, queryParams QueryParams) ([]Photo, error) {
	endPoint := topicsListEndpoint + IDOrSlug + "/photos"
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

// stats functions

// Get a list of counts for all of Unsplash.
// https://unsplash.com/documentation#totals
func (c *Client) getStatsTotal(ctx context.Context) (*StatsTotal, error) {
	data, err := c.getBodyBytes(ctx, statsTotalEndpoint)
	if err != nil {
		return nil, err
	}
	var stats StatsTotal
	err = parseJSON(data, stats)
	if err != nil {
		return nil, err
	}
	return &stats, nil
}

// Get the overall Unsplash stats for the past 30 days.
// https://unsplash.com/documentation#month
func (c *Client) getStatsMonth(ctx context.Context) (*StatsMonth, error) {
	data, err := c.getBodyBytes(ctx, statsTotalEndpoint)
	if err != nil {
		return nil, err
	}
	var stats StatsMonth
	err = parseJSON(data, stats)
	if err != nil {
		return nil, err
	}
	return &stats, nil
}
