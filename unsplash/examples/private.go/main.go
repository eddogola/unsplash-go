package main

import (
	"fmt"
	"os"

	"github.com/eddogola/unsplash-go/unsplash"
	"github.com/eddogola/unsplash-go/unsplash/client"
)

func main() {
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
		"title":       "unsplash-go test",
		"description": "will run travis tests then disappear",
	})
	checkErr(err)
	fmt.Println(newCollection.Title)

	// update collection
	updatedCollection, err := privateUnsplash.Collections.Update(newCollection.ID, map[string]string{
		"title":       "unsplash-go TEST",
		"description": "will run travis tests then disappear",
	})
	checkErr(err)
	fmt.Println(updatedCollection.Title)

	// add photo to collection
	car, err := privateUnsplash.Collections.AddPhoto(newCollection.ID, map[string]string{
		"collection_id": newCollection.ID,
		"photo_id":      "bLqKgljgpf4",
	})
	checkErr(err)
	fmt.Println(car.Photo.AltDescription)

	// remove photo from collection
	car, err = privateUnsplash.Collections.RemovePhoto(newCollection.ID, map[string]string{
		"collection_id": newCollection.ID,
		"photo_id":      "bLqKgljgpf4",
	})
	checkErr(err)
	fmt.Println(car.Photo.AltDescription)

	// delete collection
	err = privateUnsplash.Collections.Delete(newCollection.ID)
	checkErr(err)

	// like photo
	lr, err := privateUnsplash.Photos.Like("bLqKgljgpf4")
	checkErr(err)
	fmt.Println(lr.Photo.Links.Download)

	// unlike photo
	err = privateUnsplash.Photos.Unlike("bLqKgljgpf4")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("encountered unexpected error: %+v\n", err)
		os.Exit(1)
	}
}
