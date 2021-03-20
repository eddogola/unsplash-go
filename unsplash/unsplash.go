package unsplash

import (
	"github.com/eddogola/unsplash-go/unsplash/client"
)

// Unsplash wraps the whole API
type Unsplash struct {
	Users       *UsersService
	Photos      *PhotosService
	Collections *CollectionsService
	Topics      *TopicsService
	client      *client.Client
}

// New constructs a new Unsplash object
func New(c *client.Client) *Unsplash {
	unsplash := &Unsplash{client: c}
	unsplash.Users = &UsersService{client: unsplash.client}
	unsplash.Photos = &PhotosService{client: unsplash.client}
	unsplash.Collections = &CollectionsService{client: unsplash.client}

	return unsplash
}
