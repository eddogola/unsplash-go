package unsplash

import "context"

// StatsTotal defines fields for an Unsplash Total Stats Resource
type StatsTotal struct {
	Photos             int `json:"photos"`
	Downloads          int `json:"downloads"`
	Views              int `json:"views"`
	Likes              int `json:"likes"`
	Photographers      int `json:"photographers"`
	Pixels             int `json:"pixels"`
	DownloadsPerSecond int `json:"downloads_per_second"`
	ViewPerSecond      int `json:"views_per_second"`
	Developers         int `json:"developers"`
	Applications       int `json:"applications"`
	Requests           int `json:"requests"`
}

// StatsMonth defines fields for an Unsplash 30-day stats
type StatsMonth struct {
	Downloads        int `json:"downloads"`
	Views            int `json:"views"`
	Likes            int `json:"likes"`
	NewPhotos        int `json:"new_photos"`
	NewPhotographers int `json:"new_photographers"`
	NewPixels        int `json:"new_pixels"`
	NewDevelopers    int `json:"new_developers"`
	NewApplications  int `json:"new_applications"`
	NewRequests      int `json:"new_requests"`
}

// Stats defines a blueprint for statistics
type Stats struct {
	Total      int `json:"total"`
	Historical struct {
		Change     int    `json:"change"`
		Resolution string `json:"resolution"`
		Quantity   int    `json:"quantity"`
		Values     []struct {
			Date  string `json:"date"`
			Value int    `json:"value"`
		} `json:"values"`
	} `json:"historical"`
}

// PhotoStats defines specific photo statistics fields
type PhotoStats struct {
	ID        string `json:"id"`
	Downloads Stats  `json:"downloads"`
	Views     Stats  `json:"views"`
	Likes     Stats  `json:"likes"`
}

// UserStats defines specific user statistics fields
type UserStats struct {
	Username  string `json:"username"`
	Downloads Stats  `json:"downloads"`
	Views     Stats  `json:"views"`
}

// Get a list of counts for all of Unsplash.
// https://unsplash.com/documentation#totals
func (c *Client) getStatsTotal(ctx context.Context) (*StatsTotal, error) {
	data, err := c.getBodyBytes(ctx, StatsTotalEndpoint)
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
	data, err := c.getBodyBytes(ctx, StatsTotalEndpoint)
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
