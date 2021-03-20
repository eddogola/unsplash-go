package main

import (
	"context"
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
	
	// set context.Background() for all requests
	ctx := context.Background()
	
	// all topics
	topics, err := unsplash.Topics.All(ctx, nil)
	checkErr(err)
	fmt.Println(topics[0])

	// get topic
	topic, err := unsplash.Topics.Get(ctx, topics[0].ID)
	checkErr(err)
	fmt.Println(topic.Slug)

	// get topic photos
	photos, err := unsplash.Topics.Photos(ctx, topics[0].ID, nil)
	checkErr(err)
	fmt.Println(photos[0].Links.Download)
}

func checkErr(err error) {
	fmt.Printf("encountered unexpected error: %v\n", err)
	os.Exit(1)
}