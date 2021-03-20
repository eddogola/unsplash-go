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

	// all collections
	collections, err := unsplash.Collections.All(nil)
	checkErr(err)
	fmt.Println(collections[0])

	// get collection
	collection, err := unsplash.Collections.Get(collections[0].ID)
	checkErr(err)
	fmt.Println(collection.Title)

	// collection photos
	photos, err := unsplash.Collections.Photos(collections[0].ID, nil)
	checkErr(err)
	fmt.Println(photos[0])

	// related collections
	related, err := unsplash.Collections.Related(collections[0].ID)
	checkErr(err)
	fmt.Println(related[0])

	// search collections
	searchResults, err := unsplash.Collections.Search("code", nil)
	checkErr(err)
	fmt.Println(searchResults.Results[0])
}

func checkErr(err error) {
	fmt.Printf("encountered unexpected error: %v\n", err)
	os.Exit(1)
}
