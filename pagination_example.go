package unsplash

import (
	"context"
	"fmt"
	"net/http"
)

// an example,
// getting photos in the subsequent pages
func getPics() {
	clientID := "-jLuawEhNTrJByNkD-scww7cz0u-fC4W8DjMOXyKAEY"
	qParams := QueryParams(map[string]string{
		"order_by": "latest",
	})
	c := NewClient(clientID, http.DefaultClient, &Config{AuthInHeader: true, LowContentSafety: false})
	if c.Config.AuthInHeader {
		c.Config.Headers.Add("Authorization", fmt.Sprintf("Client-ID %s", clientID))
	} else {
		qParams["client_id"] = clientID
	}
	// how I'd want it to be.
	var photos []Photo
	// get the first five pages
	for i := 0; i <= 5; i++ {
		qParams["page"] = fmt.Sprint(i)
		pics, _ := c.getPhotoList(context.Background(), qParams)
		photos = append(photos, pics...)
	}
}

// check if response final page is equal to first page
// photos, search results, collections, topics
