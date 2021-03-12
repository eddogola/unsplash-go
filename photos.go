package unsplash

import "context"

// Photo defines fields in a photo resource
type Photo struct {
	ID             string `json:"id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	PromotedAt     string `json:"promoted_at"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	Color          string `json:"color"`
	Downloads      int    `json:"downloads"`
	BlurHash       string `json:"blur_hash"`
	Likes          int    `json:"likes"`
	LikedByUser    bool   `json:"liked_by_user"`
	Description    string `json:"description"`
	AltDescription string `json:"alt_description"`
	Exif           struct {
		Make         string `json:"make"`
		Model        string `json:"model"`
		ExposureTime string `json:"exposure_time"`
		Aperture     string `json:"aperture"`
		FocalLength  string `json:"focal_length"`
		ISO          int    `json:"iso"`
	} `json:"exif"`
	Location struct {
		City     string `json:"city"`
		Country  string `json:"country"`
		Position struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"position"`
	} `json:"location"`
	Tags                   []Tag        `json:"tags"`
	CurrentUserCollections []Collection `json:"current_user_collections"`
	URLs                   struct {
		Raw     string `json:"raw"`
		Full    string `json:"full"`
		Regular string `json:"regular"`
		Small   string `json:"small"`
		Thumb   string `json:"thumb"`
	} `json:"urls"`
	Links struct {
		Self             string `json:"self"`
		HTML             string `json:"html"`
		Download         string `json:"download"`
		DownloadLocation string `json:"download_location"`
	} `json:"links"`
	User       User `json:"user"`
	Statistics struct {
		Downloads Stats `json:"downloads"`
		Views     Stats `json:"views"`
		Likes     Stats `json:"likes"`
	} `json:"statistics"`
}

// Tag defines fields in a photo's tag
type Tag struct {
	Title string `json:"title"`
}

// GetPhotoList takes in a context and query parameters to return a list(given page, default first) of all
// Unsplash photos.
// Get a single page with a list of all photos
// https://unsplash.com/documentation#list-photos
func (c *Client) GetPhotoList(ctx context.Context, queryParams QueryParams) ([]Photo, error) {
	link, err := buildURL(AllPhotosEndpoint, queryParams)
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

// GetPhotoList takes in a context and photo id. If a photo with the given id exists,
// the Photo is returned.
// Get a Photo using photo ID
// https://unsplash.com/documentation#get-a-photo
func (c *Client) GetPhoto(ctx context.Context, ID string) (*Photo, error) {
	link := AllPhotosEndpoint + ID
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

// GetRandomPhoto takes in a context and query parameters. If a `count` query parameter is provided,
// a list of photos is returned, otherwise, a single photo is returned.
// Get a random Photo
// return a list of photos if a count query parameter is provided
// https://unsplash.com/documentation#get-a-random-photo
func (c *Client) GetRandomPhoto(ctx context.Context, queryParams QueryParams) (interface{}, error) {
	link, err := buildURL(RandomPhotoEndpoint, queryParams)
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

// GetPhotoStats takes in a context, photo id and query parameters. If photo with given id is found,
// the photo stats are returned in a pointer to a PhotoStats object.
// Retrieve total number of downloads, views and likes of a single photo, as well as the historical
// breakdown of these stats in a specific timeframe (default is 30 days).
// https://unsplash.com/documentation#get-a-photos-statistics
func (c *Client) GetPhotoStats(ctx context.Context, ID string, queryParams QueryParams) (*PhotoStats, error) {
	endPoint := AllPhotosEndpoint + ID + "/statistics"
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
