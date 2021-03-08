package unsplash

import "fmt"

type errStatusCode int

func (e errStatusCode) Error() string {
	return fmt.Sprintf("unexpected status code: %d", e)
}
