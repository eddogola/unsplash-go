# Unsplash API wrapper

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
      - [All](#all)
      - [Get](#get)
      - [Random](#random)
        - [`count` not provided](#count-not-provided)
        - [`count` provided](#count-provided)
      - [Stats](#stats)
      - [Search](#search)
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
- [unsplash.Photos]

  - [All](#all)
  - [Get](#get)
  - [Random](#random)
  - [Stats](#stats)
  - [Search](#search)

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

#### All

Get a paginated list of all unsplash Photos.

```go
pics, err := unsplash.Photos.All(nil)
```

#### Get

Get a specific photo.

```go
pic, err := unsplash.Photos.Get(`photo-id`)
```

#### Random

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

#### Stats

Get a specific photo's stats. Returns a `*client.PhotoStats` object.

```go
stats, err := unsplash.Photos.Stats(pics[0].ID, nil)
```

#### Search

Search photos. Returns a `*client.PhotoSearchResults` object.

```go
searchResult, err := unsplash.Photos.Search("food", nil)
fmt.Println(searchResult.Results)
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
