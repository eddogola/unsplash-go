package unsplash

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	errCodeQueryParamNotFound = errors.New("`code` query parameter not found in the request URL")
)

type errQueryNotInURL string

func (e errQueryNotInURL) Error() string {
	return "search query parameter absent in url: " + string(e)
}

type errStatusCode struct {
	statusCode int
	reasons    []string
}

func (e errStatusCode) Error() string {
	return fmt.Sprintf("unexpected status code: %d\n encountered errors: %v", e.statusCode, e.reasons)
}

func getErrReasons(resp *http.Response) []string {
	data, _ := ioutil.ReadAll(resp.Body)

	var errs Errors
	parseJSON(data, &errs)

	return errs.ErrorList
}
