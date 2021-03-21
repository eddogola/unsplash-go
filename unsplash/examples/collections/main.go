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

	// all collections
	collections, err := publicUnsplash.Collections.All(nil)
	checkErr(err)
	fmt.Println(collections[0].Links.HTML)

	// get collection
	collection, err := publicUnsplash.Collections.Get(collections[0].ID)
	checkErr(err)
	fmt.Println(collection.Title)

	// collection photos
	photos, err := publicUnsplash.Collections.Photos(collections[0].ID, nil)
	checkErr(err)
	fmt.Println(photos[0].Links.HTML)

	// related collections
	related, err := publicUnsplash.Collections.Related(collections[0].ID)
	checkErr(err)
	fmt.Println(related[0].Title)

	// search collections
	searchResults, err := publicUnsplash.Collections.Search("code", nil)
	checkErr(err)
	fmt.Println(searchResults.Results[0].Title)
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("encountered unexpected error: %+v\n", err)
		os.Exit(1)
	}
}
