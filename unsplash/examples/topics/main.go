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

	// all topics
	topics, err := unsplash.Topics.All(nil)
	checkErr(err)
	fmt.Println(topics[0])

	// get topic
	topic, err := unsplash.Topics.Get(topics[0].ID)
	checkErr(err)
	fmt.Println(topic.Slug)

	// get topic photos
	photos, err := unsplash.Topics.Photos(topics[0].ID, nil)
	checkErr(err)
	fmt.Println(photos[0].Links.Download)
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("encountered unexpected error: %+v\n", err)
		os.Exit(1)
	}
}
