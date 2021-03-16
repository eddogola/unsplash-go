package unsplash

import (
	"context"
	"os"
	"testing"

	"github.com/eddogola/unsplash-go/unsplash/client"
)

func TestPhotosService(t *testing.T) {

	t.Run("random photo when count not passed", func(t *testing.T) {
		clientID := os.Getenv("CLIENT_ID")
		cl := client.NewClient(clientID, nil, client.NewConfig())
		unsplash := New(cl)
		res, err := unsplash.Photos.GetRandom(context.Background(), nil)
		randomPhoto := res.(*client.Photo)
		checkErrorIsNil(t, err)
		checkRsNotNil(t, randomPhoto)
	})

	t.Run("random photo when count passed", func(t *testing.T) {
		clientID := os.Getenv("CLIENT_ID")
		cl := client.NewClient(clientID, nil, client.NewConfig())
		unsplash := New(cl)

		res, err := unsplash.Photos.GetRandom(context.Background(), client.QueryParams{"count": "1"})
		randomPhotos := res.([]client.Photo)
		checkErrorIsNil(t, err)
		checkRsNotNil(t, randomPhotos)
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
