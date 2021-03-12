package unsplash

// Permissions scopes
// To write data on behalf of a user or to access their private data,
// you must request additional permission scopes from them.
const (
	// Default. Read public data.
	PublicScope = "public"
	// Access user’s private data.
	ReadUserScope = "read_user"
	// Update the user’s profile.
	WriteUserScope = "write_user"
	// Read private data from the user’s photos.
	ReadPhotosScope = "read_photos"
	// Update photos on the user’s behalf.
	WritePhotosScope = "write_photos"
	// Like or unlike a photo on the user’s behalf.
	WriteLikesScope = "write_likes"
	// Follow or unfollow a user on the user’s behalf.
	WriteFollowersScope = "write_followers"
	// View a user’s private collections.
	ReadCollectionsScope = "read_collections"
	// Create and update a user’s collections.
	WriteCollectionsScope = "write_collections"
)