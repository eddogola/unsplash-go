package unsplash

type Unsplash struct {
	Users       *UsersService
	Photos      *PhotosService
	Search      *SearchService
	Collections *CollectionsService
	client      *Client
}

func NewUnsplash(c *Client) *Unsplash {
	unsplash := &Unsplash{client: c}
	unsplash.Users = &UsersService{client: unsplash.client}
	unsplash.Photos = &PhotosService{client: unsplash.client}
	unsplash.Search = &SearchService{client: unsplash.client}
	unsplash.Collections = &CollectionsService{client: unsplash.client}
	
	return unsplash
}
