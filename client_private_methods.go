package unsplash

import (
	"context"
	"io/ioutil"
)

// private Client methods
// Note: Without a Bearer token (i.e. using a Client-ID token) these requests will return a 401 Unauthorized response.

// Get the user’s private profile
// Note: To access a user’s private data, the user is required to authorize the read_user scope.
func (c *Client) getUserPrivateProfile(ctx context.Context) (*User, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, errClientNotPrivate
	}
	// check if the `read_user` scope is present in the private Client's scopes
	scope := "read_user"
	if ok := c.AuthScopes.Contains(scope); !ok {
		return nil, errRequiredScopeAbsent(scope)
	}

	data, err := c.getBodyBytes(ctx, privateUserProfileEndpoint)
	if err != nil {
		return nil, err
	}

	var usr User
	err = parseJSON(data, usr)
	if err != nil {
		return nil, err
	}
	return &usr, nil
}

// Update the current user’s profile
// https://unsplash.com/documentation#update-the-current-users-profile
// Note: This action requires the write_user scope. Without it, it will return a 403 Forbidden response.
func (c *Client) updateUserProfile(ctx context.Context, updatedData map[string]string) (*User, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, errClientNotPrivate
	}
	// check if the `write_user` scope is present in the private Client's scopes
	scope := "write_user"
	if ok := c.AuthScopes.Contains(scope); !ok {
		return nil, errRequiredScopeAbsent(scope)
	}
	// make PUT request
	// response returns the updated profile
	resp, err := c.putHTTP(ctx, privateUserProfileEndpoint, updatedData)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// parse bytes
	var usr User
	err = parseJSON(data, &usr)
	if err != nil {
		return nil, err
	}
	return &usr, nil
}

// Update a photo on behalf of the logged-in user
// This requires the `write_photos` scope
// https://unsplash.com/documentation#update-a-photo
func (c *Client) updatePhoto(ctx context.Context, ID string, updatedData map[string]string) (*Photo, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, errClientNotPrivate
	}
	// check if the `write_photo` scope is present in the private Client's scopes
	scope := "write_photo"
	if ok := c.AuthScopes.Contains(scope); !ok {
		return nil, errRequiredScopeAbsent(scope)
	}
	// make PUT request
	// response returns the updated profile
	endPoint := allPhotosEndpoint + ID
	resp, err := c.putHTTP(ctx, endPoint, updatedData)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// parse bytes
	var pic Photo
	err = parseJSON(data, &pic)
	if err != nil {
		return nil, err
	}
	return &pic, nil
}

// Like a photo on behalf of the logged-in user
// This requires the `write_likes` scope
// https://unsplash.com/documentation#like-a-photo
func (c *Client) likePhoto(ctx context.Context, ID string) (*LikeResponse, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, errClientNotPrivate
	}
	// check if the `write_likes` scope is present in the private Client's scopes
	scope := "write_likes"
	if ok := c.AuthScopes.Contains(scope); !ok {
		return nil, errRequiredScopeAbsent(scope)
	}
	// make POST request
	// response returns abbreviated versions of the picture and user
	endPoint := allPhotosEndpoint + ID + "/like"
	resp, err := c.postHTTP(ctx, endPoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// parse json response
	var lr LikeResponse
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = parseJSON(data, &lr)
	if err != nil {
		return nil, err
	}
	return &lr, nil
}

// Remove the logged-in user’s like of a photo.
// https://unsplash.com/documentation#unlike-a-photo
func (c *Client) unlikePhoto(ctx context.Context, ID string) error {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return errClientNotPrivate
	}
	// make DELETE request
	// responds with a 204 status code and an empty body
	endPoint := allPhotosEndpoint + ID + "/like"
	resp, err := c.deleteHTTP(ctx, endPoint, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 204 {
		return errStatusCode{resp.StatusCode, getErrReasons(resp)}
	}
	return nil
}

// Create a new collection. This requires the `write_collections` scope.
// https://unsplash.com/documentation#create-a-new-collection
func (c *Client) createCollection(ctx context.Context, ID string, data map[string]string) (*Collection, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, errClientNotPrivate
	}
	// check if the `write_collections` scope is present in the private Client's scopes
	scope := "write_collections"
	if ok := c.AuthScopes.Contains(scope); !ok {
		return nil, errRequiredScopeAbsent(scope)
	}
	// make POST request
	// responds with the new collection
	resp, err := c.postHTTP(ctx, collectionsListEndpoint, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// parse json response
	var collection Collection
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = parseJSON(bs, &collection)
	if err != nil {
		return nil, err
	}
	return &collection, nil
}

// Update an existing collection belonging to the logged-in user. This requires the `write_collections` scope.
// https://unsplash.com/documentation#update-an-existing-collection
// check if client is private to do private requests
func (c *Client) updateCollection(ctx context.Context, ID string, data map[string]string) (*Collection, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, errClientNotPrivate
	}
	// check if the `write_collections` scope is present in the private Client's scopes
	scope := "write_collections"
	if ok := c.AuthScopes.Contains(scope); !ok {
		return nil, errRequiredScopeAbsent(scope)
	}
	// make PUT request
	// responds with the updated collection
	endPoint := collectionsListEndpoint + ID
	resp, err := c.putHTTP(ctx, endPoint, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// parse bytes
	var collection Collection
	err = parseJSON(bs, &collection)
	if err != nil {
		return nil, err
	}
	return &collection, nil
}

// Delete a collection belonging to the logged-in user. This requires the `write_collections` scope.
// https://unsplash.com/documentation#delete-a-collection
func (c *Client) deleteCollection(ctx context.Context, ID string) error {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return errClientNotPrivate
	}
	// check if the `write_collections` scope is present in the private Client's scopes
	scope := "write_collections"
	if ok := c.AuthScopes.Contains(scope); !ok {
		return errRequiredScopeAbsent(scope)
	}
	// make DELETE request
	// responds with a 204 status code and an empty body
	endPoint := collectionsListEndpoint + ID
	resp, err := c.deleteHTTP(ctx, endPoint, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return errStatusCode{resp.StatusCode, getErrReasons(resp)}
	}
	return nil
}

// Add a photo to one of the logged-in user’s collections. Requires the `write_collections` scope.
// https://unsplash.com/documentation#add-a-photo-to-a-collection
func (c *Client) addPhotoToCollection(ctx context.Context, ID string, data map[string]string) (*CollectionActionResponse, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, errClientNotPrivate
	}
	// check if the `write_collections` scope is present in the private Client's scopes
	scope := "write_collections"
	if ok := c.AuthScopes.Contains(scope); !ok {
		return nil, errRequiredScopeAbsent(scope)
	}
	// make POST request
	// response returns abbreviated versions of the picture and user
	endPoint := collectionsListEndpoint + ID + "/add"
	resp, err := c.postHTTP(ctx, endPoint, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// parse json response
	var car CollectionActionResponse
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = parseJSON(bs, &car)
	if err != nil {
		return nil, err
	}
	return &car, nil
}

// Remove a photo from one of the logged-in user’s collections. Requires the `write_collections` scope.
// https://unsplash.com/documentation#remove-a-photo-from-a-collection
func (c *Client) removePhotoFromCollection(ctx context.Context, ID string, data map[string]string) (*CollectionActionResponse, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, errClientNotPrivate
	}
	// check if the `write_collections` scope is present in the private Client's scopes
	scope := "write_collections"
	if ok := c.AuthScopes.Contains(scope); !ok {
		return nil, errRequiredScopeAbsent(scope)
	}
	// make DELETE request
	// responds with a 204 status code and an empty body
	endPoint := collectionsListEndpoint + ID + "/remove"
	resp, err := c.deleteHTTP(ctx, endPoint, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 204 {
		return nil, errStatusCode{resp.StatusCode, getErrReasons(resp)}
	}

	// parse json response
	var car CollectionActionResponse
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = parseJSON(bs, &car)
	if err != nil {
		return nil, err
	}
	return &car, nil
}
