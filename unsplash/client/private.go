package client

import (
	"context"
	"io/ioutil"
)

/*
	Private Client methods
	Note: Without a Bearer token (i.e. using a Client-ID token) these requests will return a 401 Unauthorized response.
*/

// GetUserPrivateProfile takes a context and returns the private profile of the authenticated user
// Note: To access a user’s private data, the user is required to authorize the read_user scope.
func (c *Client) GetUserPrivateProfile(ctx context.Context) (*User, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, ErrClientNotPrivate
	}
	// check if the `read_user` scope is present in the private Client's scopes
	if ok := c.AuthScopes.Contains(ReadUserScope); !ok {
		return nil, ErrRequiredScopeAbsent(ReadUserScope)
	}

	data, err := c.getBodyBytes(ctx, PrivateUserProfileEndpoint)
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

// UpdateUserProfile takes a context and a map with the data to update a user. Returns the updated profile.
// Updates the current user’s profile.
// https://unsplash.com/documentation#update-the-current-users-profile
// Note: This action requires the write_user scope. Without it, it will return a 403 Forbidden response.
func (c *Client) UpdateUserProfile(ctx context.Context, updatedData map[string]string) (*User, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, ErrClientNotPrivate
	}
	// check if the `write_user` scope is present in the private Client's scopes
	if ok := c.AuthScopes.Contains(WriteUserScope); !ok {
		return nil, ErrRequiredScopeAbsent(WriteUserScope)
	}
	// make PUT request
	resp, err := c.putHTTP(ctx, PrivateUserProfileEndpoint, updatedData)
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

// UpdatePhoto takes in a context, photo id and data with which to update photo. Returns the updated Photo.
// Updates a photo on behalf of the logged-in user. This requires the `write_photos` scope.
// https://unsplash.com/documentation#update-a-photo
func (c *Client) UpdatePhoto(ctx context.Context, photoID string, updatedData map[string]string) (*Photo, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, ErrClientNotPrivate
	}
	// check if the `write_photo` scope is present in the private Client's scopes
	if ok := c.AuthScopes.Contains(WritePhotosScope); !ok {
		return nil, ErrRequiredScopeAbsent(WritePhotosScope)
	}
	// make PUT request
	endPoint := AllPhotosEndpoint + photoID
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

// LikePhoto takes in a context and photo id, returns abbreviated versions of the picture and user.
// Likes a photo on behalf of the logged-in user. This requires the `write_likes` scope.
// https://unsplash.com/documentation#like-a-photo
func (c *Client) LikePhoto(ctx context.Context, photoID string) (*LikeResponse, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, ErrClientNotPrivate
	}
	// check if the `write_likes` scope is present in the private Client's scopes
	if ok := c.AuthScopes.Contains(WriteLikesScope); !ok {
		return nil, ErrRequiredScopeAbsent(WriteLikesScope)
	}
	// make POST request
	endPoint := AllPhotosEndpoint + photoID + "/like"
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

// UnlikePhoto takes in a context and photo id. Returns an error if failed
// deleting like from Photo, nil otherwise.
// Removes the logged-in user’s like of a photo.
// https://unsplash.com/documentation#unlike-a-photo
func (c *Client) UnlikePhoto(ctx context.Context, photoID string) error {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return ErrClientNotPrivate
	}
	// make DELETE request
	// responds with a 204 status code and an empty body
	endPoint := AllPhotosEndpoint + photoID + "/like"
	resp, err := c.deleteHTTP(ctx, endPoint, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 204 {
		return ErrStatusCode{resp.StatusCode, getErrReasons(resp)}
	}
	return nil
}

// CreateCollection takes in a context and the new collection's data. Returns a pointer to
// the created collection.
// Creates a new collection. This requires the `write_collections` scope.
// https://unsplash.com/documentation#create-a-new-collection
func (c *Client) CreateCollection(ctx context.Context, data map[string]string) (*Collection, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, ErrClientNotPrivate
	}
	// check if the `write_collections` scope is present in the private Client's scopes
	if ok := c.AuthScopes.Contains(WriteCollectionsScope); !ok {
		return nil, ErrRequiredScopeAbsent(WriteCollectionsScope)
	}
	// make POST request
	// responds with the new collection
	resp, err := c.postHTTP(ctx, CollectionsListEndpoint, data)
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

// UpdateCollection takes in a context, collection's id and data with which to update the collection.
// Returns a pointer to the updated collection if collection with the provided collection is found.
// Updates an existing collection belonging to the logged-in user. This requires the `write_collections` scope.
// https://unsplash.com/documentation#update-an-existing-collection
// check if client is private to do private requests
func (c *Client) UpdateCollection(ctx context.Context, collectionID string, data map[string]string) (*Collection, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, ErrClientNotPrivate
	}
	// check if the `write_collections` scope is present in the private Client's scopes
	if ok := c.AuthScopes.Contains(WriteCollectionsScope); !ok {
		return nil, ErrRequiredScopeAbsent(WriteCollectionsScope)
	}
	// make PUT request
	// responds with the updated collection
	endPoint := CollectionsListEndpoint + collectionID
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

// DeleteCollection takes in a context and collection's id. Returns an error if delete has failed,
// nil otherwise.
// Deletes a collection belonging to the logged-in user. This requires the `write_collections` scope.
// https://unsplash.com/documentation#delete-a-collection
func (c *Client) DeleteCollection(ctx context.Context, collectionID string) error {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return ErrClientNotPrivate
	}
	// check if the `write_collections` scope is present in the private Client's scopes
	if ok := c.AuthScopes.Contains(WriteCollectionsScope); !ok {
		return ErrRequiredScopeAbsent(WriteCollectionsScope)
	}
	// make DELETE request
	// responds with a 204 status code and an empty body
	endPoint := CollectionsListEndpoint + collectionID
	resp, err := c.deleteHTTP(ctx, endPoint, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return ErrStatusCode{resp.StatusCode, getErrReasons(resp)}
	}
	return nil
}

// AddPhotoToCollection takes in a context, collection id, and data with photo details. `data` must have `photo_id`.
// Returns abbreviated versions of the added picture and user.
// Adds a photo to one of the logged-in user’s collections. Requires the `write_collections` scope.
// https://unsplash.com/documentation#add-a-photo-to-a-collection
func (c *Client) AddPhotoToCollection(ctx context.Context, collectionID string, data map[string]string) (*CollectionActionResponse, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, ErrClientNotPrivate
	}
	// check if the `write_collections` scope is present in the private Client's scopes
	if ok := c.AuthScopes.Contains(WriteCollectionsScope); !ok {
		return nil, ErrRequiredScopeAbsent(WriteCollectionsScope)
	}
	// make POST request
	endPoint := CollectionsListEndpoint + collectionID + "/add"
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

// RemovePhotoFromCollection takes in a context, collection id, and data with photo details. `data` must have `photo_id`.
// Returns abbreviated versions of the deleted picture and user.
// Removes a photo from one of the logged-in user’s collections. Requires the `write_collections` scope.
// https://unsplash.com/documentation#remove-a-photo-from-a-collection
func (c *Client) RemovePhotoFromCollection(ctx context.Context, collectionID string, data map[string]string) (*CollectionActionResponse, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, ErrClientNotPrivate
	}
	// check if the `write_collections` scope is present in the private Client's scopes
	if ok := c.AuthScopes.Contains(WriteCollectionsScope); !ok {
		return nil, ErrRequiredScopeAbsent(WriteCollectionsScope)
	}
	// make DELETE request
	// responds with a 204 status code and an empty body
	endPoint := CollectionsListEndpoint + collectionID + "/remove"
	resp, err := c.deleteHTTP(ctx, endPoint, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 204 {
		return nil, ErrStatusCode{resp.StatusCode, getErrReasons(resp)}
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
