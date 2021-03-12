package unsplash

import (
	"context"
	"io/ioutil"
)

// private Client methods
// Note: Without a Bearer token (i.e. using a Client-ID token) these requests will return a 401 Unauthorized response.

// GetUserPrivateProfile gets the user’s private profile
// Note: To access a user’s private data, the user is required to authorize the read_user scope.
func (c *Client) GetUserPrivateProfile(ctx context.Context) (*User, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, errClientNotPrivate
	}
	// check if the `read_user` scope is present in the private Client's scopes
	if ok := c.AuthScopes.Contains(ReadUserScope); !ok {
		return nil, errRequiredScopeAbsent(ReadUserScope)
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

// UpdateUserProfile updates the current user’s profile
// https://unsplash.com/documentation#update-the-current-users-profile
// Note: This action requires the write_user scope. Without it, it will return a 403 Forbidden response.
func (c *Client) UpdateUserProfile(ctx context.Context, updatedData map[string]string) (*User, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, errClientNotPrivate
	}
	// check if the `write_user` scope is present in the private Client's scopes
	if ok := c.AuthScopes.Contains(WriteUserScope); !ok {
		return nil, errRequiredScopeAbsent(WriteUserScope)
	}
	// make PUT request
	// response returns the updated profile
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

// UpdatePhoto updates a photo on behalf of the logged-in user
// This requires the `write_photos` scope
// https://unsplash.com/documentation#update-a-photo
func (c *Client) UpdatePhoto(ctx context.Context, ID string, updatedData map[string]string) (*Photo, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, errClientNotPrivate
	}
	// check if the `write_photo` scope is present in the private Client's scopes
	if ok := c.AuthScopes.Contains(WritePhotosScope); !ok {
		return nil, errRequiredScopeAbsent(WritePhotosScope)
	}
	// make PUT request
	// response returns the updated profile
	endPoint := AllPhotosEndpoint + ID
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

// LikePhoto likes a photo on behalf of the logged-in user
// This requires the `write_likes` scope
// https://unsplash.com/documentation#like-a-photo
func (c *Client) LikePhoto(ctx context.Context, ID string) (*LikeResponse, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, errClientNotPrivate
	}
	// check if the `write_likes` scope is present in the private Client's scopes
	if ok := c.AuthScopes.Contains(WriteLikesScope); !ok {
		return nil, errRequiredScopeAbsent(WriteLikesScope)
	}
	// make POST request
	// response returns abbreviated versions of the picture and user
	endPoint := AllPhotosEndpoint + ID + "/like"
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

// UnlikePhoto removes the logged-in user’s like of a photo.
// https://unsplash.com/documentation#unlike-a-photo
func (c *Client) UnlikePhoto(ctx context.Context, ID string) error {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return errClientNotPrivate
	}
	// make DELETE request
	// responds with a 204 status code and an empty body
	endPoint := AllPhotosEndpoint + ID + "/like"
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

// CreateCollection creates a new collection. This requires the `write_collections` scope.
// https://unsplash.com/documentation#create-a-new-collection
func (c *Client) CreateCollection(ctx context.Context, ID string, data map[string]string) (*Collection, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, errClientNotPrivate
	}
	// check if the `write_collections` scope is present in the private Client's scopes
	if ok := c.AuthScopes.Contains(WriteCollectionsScope); !ok {
		return nil, errRequiredScopeAbsent(WriteCollectionsScope)
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

// UpdateCollection updates an existing collection belonging to the logged-in user. This requires the `write_collections` scope.
// https://unsplash.com/documentation#update-an-existing-collection
// check if client is private to do private requests
func (c *Client) UpdateCollection(ctx context.Context, ID string, data map[string]string) (*Collection, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, errClientNotPrivate
	}
	// check if the `write_collections` scope is present in the private Client's scopes
	if ok := c.AuthScopes.Contains(WriteCollectionsScope); !ok {
		return nil, errRequiredScopeAbsent(WriteCollectionsScope)
	}
	// make PUT request
	// responds with the updated collection
	endPoint := CollectionsListEndpoint + ID
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

// DeleteCollection deletes a collection belonging to the logged-in user. This requires the `write_collections` scope.
// https://unsplash.com/documentation#delete-a-collection
func (c *Client) DeleteCollection(ctx context.Context, ID string) error {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return errClientNotPrivate
	}
	// check if the `write_collections` scope is present in the private Client's scopes
	if ok := c.AuthScopes.Contains(WriteCollectionsScope); !ok {
		return errRequiredScopeAbsent(WriteCollectionsScope)
	}
	// make DELETE request
	// responds with a 204 status code and an empty body
	endPoint := CollectionsListEndpoint + ID
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

// AddPhotoToCollection adds a photo to one of the logged-in user’s collections. Requires the `write_collections` scope.
// https://unsplash.com/documentation#add-a-photo-to-a-collection
func (c *Client) AddPhotoToCollection(ctx context.Context, ID string, data map[string]string) (*CollectionActionResponse, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, errClientNotPrivate
	}
	// check if the `write_collections` scope is present in the private Client's scopes
	if ok := c.AuthScopes.Contains(WriteCollectionsScope); !ok {
		return nil, errRequiredScopeAbsent(WriteCollectionsScope)
	}
	// make POST request
	// response returns abbreviated versions of the picture and user
	endPoint := CollectionsListEndpoint + ID + "/add"
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

// RemovePhotoFromCollection removes a photo from one of the logged-in user’s collections. Requires the `write_collections` scope.
// https://unsplash.com/documentation#remove-a-photo-from-a-collection
func (c *Client) RemovePhotoFromCollection(ctx context.Context, ID string, data map[string]string) (*CollectionActionResponse, error) {
	// check if client is private to do private requests
	if !isClientPrivate(c) {
		return nil, errClientNotPrivate
	}
	// check if the `write_collections` scope is present in the private Client's scopes
	if ok := c.AuthScopes.Contains(WriteCollectionsScope); !ok {
		return nil, errRequiredScopeAbsent(WriteCollectionsScope)
	}
	// make DELETE request
	// responds with a 204 status code and an empty body
	endPoint := CollectionsListEndpoint + ID + "/remove"
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
