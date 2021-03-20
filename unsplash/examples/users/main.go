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

	username := "devsnice"

	// public profile
	profile, err := unsplash.Users.PublicProfile(username)
	checkErr(err)
	fmt.Println(profile)

	// portfolio url
	url, err := unsplash.Users.PortfolioURL(username)
	checkErr(err)
	fmt.Println(url)

	// user photos
	photos, err := unsplash.Users.Photos(username, nil)
	checkErr(err)
	fmt.Println(photos[0].Links.Download)

	// user liked photos
	likedPhotos, err := unsplash.Users.LikedPhotos(username, nil)
	checkErr(err)
	fmt.Println(likedPhotos[0])

	// user collections
	collections, err := unsplash.Users.Collections(username, nil)
	checkErr(err)
	fmt.Println(collections[0])

	// user stats
	stats, err := unsplash.Users.Stats(username, nil)
	checkErr(err)
	fmt.Println(stats)

	// search users
	searchResults, err := unsplash.Users.Search(username, nil)
	checkErr(err)
	fmt.Println(searchResults)
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("encountered unexpected error: %v\n", err)
		os.Exit(1)
	}
}
