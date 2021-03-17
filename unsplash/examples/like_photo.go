package main

import (
	"context"
	"fmt"
	"os"

	"github.com/eddogola/unsplash-go/unsplash"
	"github.com/eddogola/unsplash-go/unsplash/client"
)

func main() {
	privateClient, err := client.NewPrivateAuthClient(context.Background(),
		os.Getenv("CLIENT_ID"),
		os.Getenv("CLIENT_SECRET"),
		os.Getenv("REDIRECT_URI"),
		nil,
		client.NewConfig(),
		client.NewAuthScopes(client.WriteLikesScope),
	)
	if err != nil {
		fmt.Printf("encountered unexpected error: %v\n", err)
		os.Exit(1)
	}

	unsplash := unsplash.New(privateClient)
	res, err := unsplash.Photos.Like(context.Background(), "lb9hi0NDjT0")
	if err != nil {
		fmt.Printf("encountered unexpected error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(res.Photo)
}
