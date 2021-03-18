# Unsplash API wrapper

A simple wrapper around the unsplash API.

## Authentication

Client ID passed in Authorization headers by default, not query parameters.

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
