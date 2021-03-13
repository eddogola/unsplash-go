package unsplash

import (
	"github.com/eddogola/unsplash-go/unsplash/client"
)

type Unsplash struct {
	Users       *UsersService
	Photos      *PhotosService
	Collections *CollectionsService
	Topics 		*TopicsService
	client      *client.Client
}

func NewUnsplash(c *client.Client) *Unsplash {
	unsplash := &Unsplash{client: c}
	unsplash.Users = &UsersService{client: unsplash.client}
	unsplash.Photos = &PhotosService{client: unsplash.client}
	unsplash.Collections = &CollectionsService{client: unsplash.client}
	
	return unsplash
}
