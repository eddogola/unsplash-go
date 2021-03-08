package unsplash

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

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
