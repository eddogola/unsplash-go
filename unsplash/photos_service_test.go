package unsplash

import (
	"context"
	"testing"

	"github.com/eddogola/unsplash-go/unsplash/client"
)

func TestPhotosService(t *testing.T) {

	t.Run("random photo when count not passed", func(t *testing.T) {
		cl := client.NewClient("-jLuawEhNTrJByNkD-scww7cz0u-fC4W8DjMOXyKAEY", nil, client.NewConfig())
		unsplash := New(cl)
		res, err := unsplash.Photos.GetRandom(context.Background(), nil)
		random_photo := res.(*client.Photo)
		checkErrorIsNil(t, err)
		checkRsNotNil(t, random_photo)
	})

	t.Run("random photo when count passed", func(t *testing.T) {
		cl := client.NewClient("-jLuawEhNTrJByNkD-scww7cz0u-fC4W8DjMOXyKAEY", nil, client.NewConfig())
		unsplash := New(cl)
		qParams :=  client.QueryParams(map[string]string{
			"count": "1",
		})
		res, err := unsplash.Photos.GetRandom(context.Background(), qParams)
		random_photos := res.([]client.Photo)
		checkErrorIsNil(t, err)
		checkRsNotNil(t, random_photos)
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
