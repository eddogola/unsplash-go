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
	ErrCodeQueryParamNotFound = errors.New("`code` query parameter not found in the request URL")
	ErrClientNotPrivate       = errors.New("client not private but used for functions that require private authentication")
	ErrAuthCodeEmpty          = errors.New("Auth code provided is empty")
)

type ErrQueryNotInURL string
type ErrRequiredScopeAbsent string

func (e ErrQueryNotInURL) Error() string {
	return "search query parameter absent in url: " + string(e)
}

func (e ErrRequiredScopeAbsent) Error() string {
	return "required scope `%v` not in client auth scopes"
}

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
	parseJSON(data, &errs)

	return errs.ErrorList
}
