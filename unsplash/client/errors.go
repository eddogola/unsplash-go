package client

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Errors defines the structure Unsplash responds with when an error
// is encountered while using their API.
type Errors struct {
	ErrorList []string `json:"errors"`
}

var (
	// ErrCodeQueryParamNotFound is raised when code query parameter is not in the url used for
	// oauth token access.
	ErrCodeQueryParamNotFound = errors.New("`code` query parameter not found in the request URL")
	// ErrClientNotPrivate is raised when the client usedd with the Unsplash object is not authentocated
	// for private functions.
	ErrClientNotPrivate = errors.New("client not private but used for functions that require private authentication")
	// ErrAuthCodeEmpty is raised when the authorization code is empty.
	ErrAuthCodeEmpty = errors.New("auth code provided is empty")
)

// ErrQueryNotInURL is raised when a search query parameter is not part of the url.
type ErrQueryNotInURL string

// ErrRequiredScopeAbsent is raised on trying to access a private action
// when the required scope is not provided or allowed from the authenticated user's endd.
type ErrRequiredScopeAbsent string

func (e ErrQueryNotInURL) Error() string {
	return "search query parameter absent in url: " + string(e)
}

func (e ErrRequiredScopeAbsent) Error() string {
	return "required scope `%v` not in client auth scopes"
}

// ErrStatusCode defines a http status code error
// with the status code and tthe reasons for the error
type ErrStatusCode struct {
	statusCode int
	reasons    []string
}

func (e ErrStatusCode) Error() string {
	return fmt.Sprintf("unexpected status code: %d\n encountered errors: %v", e.statusCode, e.reasons)
}

func getErrReasons(resp *http.Response) []string {
	data, _ := ioutil.ReadAll(resp.Body)

	var errs Errors
	err := parseJSON(data, &errs)

	errs.ErrorList = append(errs.ErrorList, err.Error())
	return errs.ErrorList
}
