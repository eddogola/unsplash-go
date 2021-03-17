package unsplash

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/eddogola/unsplash-go/unsplash/client"
)

var pics = []client.Photo{
	{
		ID:        "whackID",
		Width:     1308,
		Height:    768,
		Downloads: 400,
		Likes:     799,
	},
	{
		ID:     "noID",
		Width:  450,
		Height: 600,
	},
}

var pic = client.Photo{
	ID:        "whackID",
	Width:     1308,
	Height:    768,
	Downloads: 400,
	Likes:     799,
}

type mockPhotoServiceClient struct{}

func (m *mockPhotoServiceClient) GetPhotoList(ctx context.Context, queryParams client.QueryParams) ([]client.Photo, error) {
	return pics, nil
}
func (m *mockPhotoServiceClient) GetPhoto(ctx context.Context, photoID string) (*client.Photo, error) {
	return &pic, nil
}
func (m *mockPhotoServiceClient) GetRandomPhoto(ctx context.Context, queryParams client.QueryParams) (interface{}, error) {
	if _, ok := queryParams["count"]; ok {
		return pics, nil
	}
	return &pic, nil
}
func (m *mockPhotoServiceClient) GetPhotoStats(ctx context.Context, photoID string, queryParams client.QueryParams) (*client.PhotoStats, error) {
	return &client.PhotoStats{ID: "ladida"}, nil
}
func (m *mockPhotoServiceClient) SearchPhotos(ctx context.Context, queryParams client.QueryParams) (*client.PhotoSearchResult, error) {
	return &client.PhotoSearchResult{Results: pics}, nil
}
func (m *mockPhotoServiceClient) UpdatePhoto(ctx context.Context, photoID string, updatedData map[string]string) (*client.Photo, error) {
	return &pic, nil
}
func (m *mockPhotoServiceClient) LikePhoto(ctx context.Context, photoID string) (*client.LikeResponse, error) {
	return &client.LikeResponse{Photo: pic}, nil
}
func (m *mockPhotoServiceClient) UnlikePhoto(ctx context.Context, photoID string) error {
	return nil
}

func TestPhotosService(t *testing.T) {
	mockUnsplash := &Unsplash{
		Photos: &PhotosService{client: &mockPhotoServiceClient{}},
	}
	realUnsplash := New(client.New(os.Getenv("CLIENT_ID"), nil, client.NewConfig()))

	t.Run("all photos", func(t *testing.T) {
		got, err := mockUnsplash.Photos.All(context.Background(), nil)
		checkErrorIsNil(t, err)
		checkRsNotNil(t, got)
		if !reflect.DeepEqual(got, pics) {
			t.Errorf("expected %v but got %v", pics, got)
		}
	})

	t.Run("all photos, when per_page is passed as 2", func(t *testing.T) {
		got, err := realUnsplash.Photos.All(context.Background(), client.QueryParams{"per_page": "2"})
		checkErrorIsNil(t, err)
		checkRsNotNil(t, got)

		lenExpected := 2
		lenGot := len(got)
		if lenExpected != lenGot {
			t.Errorf("expected length %v but got %v", lenExpected, lenGot)
		}
	})

	t.Run("get specific photo", func(t *testing.T) {
		got, err := mockUnsplash.Photos.Get(context.Background(), "someID")
		checkErrorIsNil(t, err)
		checkRsNotNil(t, got)
		if !reflect.DeepEqual(got, &pic) {
			t.Errorf("expected %v but got %v", pics, got)
		}
	})

	t.Run("random photo when count not passed", func(t *testing.T) {
		res, err := mockUnsplash.Photos.GetRandom(context.Background(), nil)
		randomPhoto := res.(*client.Photo)
		checkErrorIsNil(t, err)
		checkRsNotNil(t, randomPhoto)
	})

	t.Run("random photo when count passed", func(t *testing.T) {
		res, err := mockUnsplash.Photos.GetRandom(context.Background(), client.QueryParams{"count": "1"})
		randomPhotos := res.([]client.Photo)
		checkErrorIsNil(t, err)
		checkRsNotNil(t, randomPhotos)
	})

	t.Run("search photos", func(t *testing.T) {
		got, err := mockUnsplash.Photos.Search(context.Background(), "food", nil)
		checkErrorIsNil(t, err)
		checkRsNotNil(t, got)

		if !reflect.DeepEqual(got.Results, pics) {
			t.Errorf("expected %v but got %v", pics, got.Results)
		}
	})

	t.Run("search photos with per_page set", func(t *testing.T) {
		got, err := realUnsplash.Photos.Search(context.Background(), "code", client.QueryParams{"per_page": "5"})
		checkErrorIsNil(t, err)
		checkRsNotNil(t, got)

		lenExpected := 5
		lenGot := len(got.Results)
		if lenGot != lenExpected {
			t.Errorf("expected length %v but got %v", lenExpected, lenGot)
		}
	})

	t.Run("like photo with non-private client", func(t *testing.T) {
		res, err := realUnsplash.Photos.Like(context.Background(), "someID")
		if res != nil {
			t.Errorf("expected nil but got %v", res)
		}

		if err != client.ErrClientNotPrivate {
			t.Errorf("expected error %v but got %v", client.ErrClientNotPrivate, err)
		}
	})

	t.Run("like photo", func(t *testing.T) {
		res, err := mockUnsplash.Photos.Like(context.Background(), "lb9hi0NDjT0")
		checkErrorIsNil(t, err)
		checkRsNotNil(t, res)
	})

	t.Run("unlike photo", func(t *testing.T) {
		err := mockUnsplash.Photos.Unlike(context.Background(), "lb9hi0NDjT0")
		checkErrorIsNil(t, err)
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
	if reflect.ValueOf(rs).IsNil() {
		t.Errorf("resource gotten is nil: %v", rs)
	}
}
