# Unsplash API wrapper

[![Build Status](https://travis-ci.com/eddogola/unsplash-go.svg?branch=main)](https://travis-ci.com/eddogola/unsplash-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/eddogola/unsplash-go)](https://goreportcard.com/report/github.com/eddogola/unsplash-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/eddogola/unsplash-go.svg)](https://pkg.go.dev/github.com/eddogola/unsplash-go)
[![codecov](https://codecov.io/gh/eddogola/unsplash-go/branch/main/graph/badge.svg?token=AH24IA8W7G)](https://codecov.io/gh/eddogola/unsplash-go)

A simple wrapper around the unsplash API.

## Documentation

- [Unsplash API wrapper](#unsplash-api-wrapper)
  - [Documentation](#documentation)
  - [Installation](#installation)
  - [API Guidelines](#api-guidelines)
  - [Registration](#registration)
  - [Usage](#usage)
    - [Importing](#importing)
    - [Create Unsplash instance](#create-unsplash-instance)
    - [unsplash.Photos](#unsplashphotos)
      - [Photos.All](#photosall)
      - [Photos.Get](#photosget)
      - [Photos.Random](#photosrandom)
        - [`count` not provided](#count-not-provided)
        - [`count` provided](#count-provided)
      - [Photos.Stats](#photosstats)
      - [Photos.Search](#photossearch)
    - [unsplash.Users](#unsplashusers)
      - [Users.PublicProfile](#userspublicprofile)
      - [Users.PortfolioURL](#usersportfoliourl)
      - [Users.Photos](#usersphotos)
      - [Users.LikedPhotos](#userslikedphotos)
      - [Users.Collections](#userscollections)
      - [Users.Stats](#usersstats)
      - [Users.Search](#userssearch)
    - [unsplash.Collections](#unsplashcollections)
      - [Collections.All](#collectionsall)
      - [Collections.Get](#collectionsget)
      - [Collections.Photos](#collectionsphotos)
      - [Collections.Related](#collectionsrelated)
      - [Collections.Search](#collectionssearch)
    - [unsplash.Topics](#unsplashtopics)
      - [Topics.All](#topicsall)
      - [Topics.Get](#topicsget)
      - [Topics.Photos](#topicsphotos)
  - [Authentication](#authentication)
  - [Buggy areas](#buggy-areas)
  - [Potential areas of improvement](#potential-areas-of-improvement)

## Installation

```bash
go get github.com/eddogola/unsplash-go/unsplash
```

## API Guidelines

When using the Unsplash API, you need to make sure to abide by their [API guidelines](https://medium.com/unsplash/unsplash-api-guidelines-28e0216e6daa) and [API Terms](https://unsplash.com/api-terms).

## Registration

[Sign up](https://unsplash.com/join) on Unsplash.com and register as a [developer](https://unsplash.com/developers).
You can then [create a new application](https://unsplash.com/oauth/applications/new) and use the AppID and Secret for authentication.

## Usage

- [Importing](#importing)
- [Initialize Unsplash](#create-unsplash-instance)
- [unsplash.Photos](#unsplashphotos)

  - [All](#photosall)
  - [Get](#photosget)
  - [Random](#photosrandom)
  - [Stats](#photossearch)
  - [Search](#photossearch)
- [unsplash.Users](#unsplashusers)

  - [PublicProfile](#userspublicprofile)
  - [PortfolioURL](#usersportfoliourl)
  - [Photos](#usersphotos)
  - [Liked Photos](#userslikedphotos)
  - [Collections](#userscollections)
  - [Stats](#usersstats)
  - [Search](#userssearch)
- [unsplash.Collections](#unsplashcollections)
  - [All](#collectionsall)
  - [Get](#collectionsget)
  - [Photos](#collectionsphotos)
  - [Related](#collectionsrelated)
  - [Search](#collectionssearch)
- [unsplash.Topic](#unsplashtopics)
  - [All](#topicsall)
  - [Get](#topicsget)
  - [Photos](#topicsphotos)

### Importing

Once you've installed the library using [go get](#installation), import it as follows

```go
import (
    "github.com/eddogola/unsplash-go/unsplash"
    "github.com/eddogola/unsplash-go/unsplash/client"
)
```

### Create Unsplash instance

```go
import (
    "github.com/eddogola/unsplash-go/unsplash"
    "github.com/eddogola/unsplash-go/unsplash/client"
)

unsplash := unsplash.New(client.New(
  os.Getenv("CLIENT_ID"), // <YOUR-CLIENT-ID>
  nil, // when nil is passed, http.DefaultClient is used
  client.NewConfig(),
  ))
```

### unsplash.Photos

#### Photos.All

Get a paginated list of all unsplash Photos.

```go
pics, err := unsplash.Photos.All(nil)
```

#### Photos.Get

Get a specific photo.

```go
pic, err := unsplash.Photos.Get(`photo-id`)
```

#### Photos.Random

Get a random photo. Returns an interface depending on whether the `count` query parameter is provided.
If present, a list of photos is returned, otherwise, a single photo is returned

##### `count` not provided

```go
res, err := unsplash.Photos.Random(nil)
randomPhoto := res.(*client.Photo)
```

##### `count` provided

```go
res, err := unsplash.Photos.Random(client.QueryParams{"count": "1"})
randomPhotos := res.([]client.Photo)
```

#### Photos.Stats

Get a specific photo's stats. Returns a `*client.PhotoStats` object.

```go
stats, err := unsplash.Photos.Stats(pics[0].ID, nil)
```

#### Photos.Search

Search photos. Returns a `*client.PhotoSearchResults` object.

```go
searchResult, err := unsplash.Photos.Search("food", nil)
fmt.Println(searchResult.Results)
```

### unsplash.Users

#### Users.PublicProfile

Get a user's public profile. Returns a `*client.User` object.

```go
profile, err := unsplash.Users.PublicProfile(`username`)
```

#### Users.PortfolioURL

Parses a user's portoflio URL, returning it in a `*url.URL` object.

```go
url, err := unsplash.Users.PortfolioURL(`username`)
```

#### Users.Photos

Get a user's photos.

```go
photos, err := unsplash.Users.Photos(`username`, nil)
```

#### Users.LikedPhotos

Get a user's liked photos.

```go
likedPhotos, err := unsplash.Users.LikedPhotos(`username`, nil)
```

#### Users.Collections

Get collections created by user.

```go
collections, err := unsplash.Users.Collections(`username`, nil)
```

#### Users.Stats

Get a user's stats. Returns a `*client.UserStats` object.

```go
stats, err := unsplash.Users.Stats(username, nil)
```

#### Users.Search

Search for a user. Returns search results in a `*client.UserSearchResults` object

```go
searchResult, err := unsplash.Users.Search(`username`, nil)
fmt.Println(searchResult.Results)
```

### unsplash.Collections

#### Collections.All

Returns a list of all collections.

```go
collections, err := unsplash.Collections.All(nil)
```

#### Collections.Get

Get a specific collection

```go
collection, err := unsplash.Collections.Get(`collectionID`)
```

#### Collections.Photos

Returns the given collection's photos.

```go
photos, err := unsplash.Collections.Photos(`collectionID`, nil)
```

#### Collections.Related

Returns a list of collections related to the given collection.

```go
related, err := unsplash.Collections.Related(`collectionID`)
```

#### Collections.Search

Returns the results of searching a collection in a `*client.CollectionSearchResult` object.

```go
searchResults, err := unsplash.Collections.Search("code", nil)
```

### unsplash.Topics

#### Topics.All

Returns a list of all topics.

```go
topics, err := unsplash.Topics.All(nil)
```

#### Topics.Get

Returns a specific `*client.Topic` object, given the topic's ID.

```go
topic, err := unsplash.Topics.Get(`topicID`)
```

#### Topics.Photos

Returns a list of `*client.Photo`s in the given the topic's ID.

```go
photos, err := unsplash.Topics.Photos(`topicID`, nil)
```

## Authentication

The user's client ID is passed in Authorization headers by default, not query parameters.

## Buggy areas

Private client authentication not fully functional.

## Potential areas of improvement

pagination not yet implemented.
The below code can exemplifies a way to work around pagination. Suggestions are welcome.\

```go
package unsplash

import (
 "context"
"fmt"
"net/http"

"github.com/eddogola/unsplash-go/unsplash"
"github.com/eddogola/unsplash-go/unsplash/client"
)

// an example of getting photos in the subsequent pages
func getPics() {
// library initialization
clientID := "<YOUR-CLIENT-ID>"
cl := client.NewClient(clientID, http.DefaultClient, client.NewConfig())
unsplash := unsplash.New(cl)

// query parameters to be passed to the request
qParams := client.QueryParams(map[string]string{
    "order_by": "latest",
})

var photos []Photo
// get the first five pages
for i := 0; i <= 5; i++ {
    qParams["page"] = fmt.Sprint(i)
    pics, err := cl.GetPhotoList(context.Background(), qParams)
    photos = append(photos, pics...)
}
}

```
