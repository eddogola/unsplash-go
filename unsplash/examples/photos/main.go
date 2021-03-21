package main

import (
	"fmt"
	"os"

	"github.com/eddogola/unsplash-go/unsplash"
	"github.com/eddogola/unsplash-go/unsplash/client"
)

func main() {
	// Initialize unsplash for public actions
	publicUnsplash := unsplash.New(client.New(
		os.Getenv("CLIENT_ID"),
		nil, // when nil is passed, http.DefaultClient is used
		client.NewConfig(),
	))

	// All photos
	pics, err := publicUnsplash.Photos.All(nil)
	checkErr(err)
	fmt.Println(pics[0].Links.Download)

	// get photo
	pic, err := publicUnsplash.Photos.Get(pics[0].ID)
	checkErr(err)
	fmt.Println(pic.Links.Download)

	// Random photo
	res, err := publicUnsplash.Photos.Random(nil)
	checkErr(err)
	randomPhoto := res.(*client.Photo)
	fmt.Println(randomPhoto.Links.Download)

	// photo stats
	stats, err := publicUnsplash.Photos.Stats(pics[0].ID, nil)
	checkErr(err)
	fmt.Println(stats.ID)

	// search photo
	searchResult, err := publicUnsplash.Photos.Search("food", nil)
	checkErr(err)
	fmt.Println(searchResult.Results[0].Links.Download)
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("encountered unexpected error: %+v\n", err)
		os.Exit(1)
	}
}
