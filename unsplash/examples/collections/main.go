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

	/**Setup private client for actions requiring private authorization**/
	privClient, err := client.NewPrivateAuthClient(
		os.Getenv("CLIENT_ID"),
		os.Getenv("CLIENT_SECRET"),
		os.Getenv("REDIRECT_URI"),
		client.NewAuthScopes("write_likes"),
		client.NewConfig(),
	)
	privateUnsplash := unsplash.New(privClient)
	/*********************************************************************/

	// create collection
	newCollection, err := privateUnsplash.Collections.Create(map[string]string{
		"title": "unsplash-go test",
		"description": "will run travis tests then disappear",
	})
	checkErr(err)
	if newCollection.Title != "unsplash-go test" {
		fmt.Printf("expected %v but got %v", "unsplash-go test", collection.Title)
		os.Exit(1)
	}

	// update collection
	updatedCollection, err := privateUnsplash.Collections.Update(newCollection.ID, map[string]string{
		"title": "unsplash-go TEST",
		"description": "will run travis tests then disappear",
	})
	checkErr(err)
	if updatedCollection.Title != "unsplash-go TEST" {
		fmt.Printf("expected %v but got %v", "unsplash-go TEST", updatedCollection.Title)
		os.Exit(1)
	}

	// delete collection
	err = privateUnsplash.Collections.Delete(newCollection.ID)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("encountered unexpected error: %+v\n", err)
		os.Exit(1)
	}
}
