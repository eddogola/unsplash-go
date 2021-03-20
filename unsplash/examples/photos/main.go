package main

import (
	"fmt"
	"os"

	"github.com/eddogola/unsplash-go/unsplash"
	"github.com/eddogola/unsplash-go/unsplash/client"
)

func main() {
	// Initialize unsplash for public actions
	unsplash := unsplash.New(client.New(
		os.Getenv("CLIENT_ID"),
		nil, // when nil is passed, http.DefaultClient is used
		client.NewConfig(),
	))

	// All photos
	pics, err := unsplash.Photos.All(nil)
	checkErr(err)
	fmt.Println(pics[0])

	// get photo
	pic, err := unsplash.Photos.Get(pics[0].ID)
	checkErr(err)
	fmt.Println(pic.Links.Download)

	// Random photo
	res, err := unsplash.Photos.Random(nil)
	checkErr(err)
	randomPhoto := res.(*client.Photo)
	fmt.Println(randomPhoto.Links.Download)

	// photo stats
	stats, err := unsplash.Photos.Stats(pics[0].ID, nil)
	checkErr(err)
	fmt.Println(stats)

	// search photo
	searchResult, err := unsplash.Photos.Search("food", nil)
	checkErr(err)
	fmt.Println(searchResult.Results[0])
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("encountered unexpected error: %v\n", err)
		os.Exit(1)
	}
}

