package unsplash

import "context"

// User defines public & private fields Unsplash provides on a user
type User struct {
	ID                string `json:"id"`
	UpdatedAt         string `json:"updated_at"`
	Username          string `json:"username"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	InstagramUsername string `json:"instagram_username"`
	TwitterUsername   string `json:"twitter_username"`
	PortfolioURL      string `json:"portfolio_url"`
	Bio               string `json:"bio"`
	Location          string `json:"location"`
	TotalLikes        int    `json:"total_likes"`
	TotalPhotos       int    `json:"total_photos"`
	TotalCollections  int    `json:"total_collections"`
	FollowedByUser    bool   `json:"followed_by_user"`
	FollowersCount    int    `json:"followers_count"`
	FollowingCount    int    `json:"following_count"`
	Downloads         int    `json:"downloads"`
	UploadsRemaining  int    `json:"uploads_remaining"`
	AcceptedTos       bool   `json:"accepted_tos"`
	ProfileImage      struct {
		Small  string `json:"small"`
		Medium string `json:"medium"`
		Large  string `json:"large"`
	} `json:"profile_image"`
	Badge struct {
		Title   string `json:"title"`
		Primary bool   `json:"primary"`
		Slug    string `json:"slug"`
		Link    string `json:"link"`
	} `json:"badge"`
	Links struct {
		Self      string `json:"self"`
		HTML      string `json:"html"`
		Photos    string `json:"photos"`
		Likes     string `json:"likes"`
		Portfolio string `json:"portfolio"`
		Following string `json:"following"`
		Followers string `json:"followers"`
	} `json:"links"`
}

// LikeResponse defines the struct returned on liking and unliking photos
// returns abbreviated versions of the picture and User
type LikeResponse struct {
	Photo Photo `json:"photo"`
	User  User  `json:"user"`
}

// GetUserPublicProfile takes in a context and username. Returns a pointer to a User,
// if user with provided username is found.
// Retrieves public details on a given user.
// https://unsplash.com/documentation#get-a-users-public-profile
func (c *Client) GetUserPublicProfile(ctx context.Context, username string) (*User, error) {
	endPoint := BaseUserEndpoint + username
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

// GetUserPortfolioLink takes in a context and username, returns the user's portfolio link.
// Retrieves a single user’s portfolio link.
// https://unsplash.com/documentation#get-a-users-portfolio-link
func (c *Client) GetUserPortfolioLink(ctx context.Context, username string) (string, error) {
	endPoint := BaseUserEndpoint + username + "/portfolio"
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

// GetUserPhotos takes a context, user's username, and query parameters. Returns a slice of
// Photo objects uploaded by the user, if user with provided username is found.
// Gets a list of photos uploaded by a user.
// https://unsplash.com/documentation#list-a-users-photos
func (c *Client) GetUserPhotos(ctx context.Context, username string, queryParams QueryParams) ([]Photo, error) {
	endPoint := BaseUserEndpoint + username
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

// GetUserLikedPhotos takes a context, user's username, and query parameters. Returns a slice of
// Photo objects liked by the user, if user with provided username is found.
// Gets a list of photos liked by a user.
// https://unsplash.com/documentation#list-a-users-liked-photos
func (c *Client) GetUserLikedPhotos(ctx context.Context, username string, queryParams QueryParams) ([]Photo, error) {
	endPoint := BaseUserEndpoint + username + "/likes"
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

// GetUserCollections takes a context, user's username, and query parameters. Returns a slice of
// Collection objects if user with provided username is found.
// Gets a list of collections created by the user.
// https://unsplash.com/documentation#list-a-users-collections
func (c *Client) GetUserCollections(ctx context.Context, username string, queryParams QueryParams) ([]Collection, error) {
	endPoint := BaseUserEndpoint + username + "/collections"
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

// GetUserStats takes a context, user's username, and query parameters. Returns user's stats if user with provided
// username is found.
// Retrieves the consolidated number of downloads, views and likes of all user’s photos,
// as well as the historical breakdown and average of these stats in a specific timeframe (default is 30 days).
// https://unsplash.com/documentation#get-a-users-statistics
func (c *Client) GetUserStats(ctx context.Context, username string, queryParams QueryParams) (*UserStats, error) {
	endPoint := BaseUserEndpoint + username + "/statistics"
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
