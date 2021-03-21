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
