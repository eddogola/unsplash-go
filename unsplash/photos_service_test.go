package unsplash

import (
	"context"
	"testing"

	"github.com/eddogola/unsplash-go/unsplash/client"
)

func TestPhotosService(t *testing.T) {

	t.Run("random photo", func(t *testing.T) {
		cl := client.NewClient("client_id", nil, client.NewConfig())
		unsplash := New(cl)
		random_photo, err := unsplash.Photos.GetRandom(context.Background(), nil)
		checkErrorIsNil(t, err)
		checkRsNotNil(t, random_photo)
	})
}

func checkErrorIsNil(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

// raise error if resource is nil
func checkRsNotNil(t *testing.T, rs interface{}) {
	t.Helper()
	if rs == nil {
		t.Errorf("resource gotten is nil: %v", rs)
	}
}
