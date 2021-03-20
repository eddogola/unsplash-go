package unsplash

import (
	"context"
	"fmt"
	"os"

	"github.com/eddogola/unsplash-go/unsplash/client"
)

func ExamplePhotosService_GetRandom() {
	cl := client.New(os.Getenv("CLIENT_ID"), nil, client.NewConfig())
	unsplash := New(cl)
	res, err := unsplash.Photos.GetRandom(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
	}

	randomPhoto := res.(*client.Photo)
	randomPhoto.ID = "losemycool"
	fmt.Println(randomPhoto.URLs)
	// losemycool
}
