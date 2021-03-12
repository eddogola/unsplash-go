package unsplash

// Permissions scopes
// To write data on behalf of a user or to access their private data,
// you must request additional permission scopes from them.
const (
	publicScope           = "public"            // Default. Read public data.
	readUserScope         = "read_user"         // Access user’s private data.
	writeUserScope        = "write_user"        // Update the user’s profile.
	readPhotosScope       = "read_photos"       // Read private data from the user’s photos.
	writePhotosScope      = "write_photos"      // Update photos on the user’s behalf.
	writeLikesScope       = "write_likes"       // Like or unlike a photo on the user’s behalf.
	writeFollowersScope   = "write_followers"   // Follow or unfollow a user on the user’s behalf.
	readCollectionsScope  = "read_collections"  // View a user’s private collections.
	writeCollectionsScope = "write_collections" // Create and update a user’s collections.
)