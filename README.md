# Unsplash API wrapper

A simple wrapper around the unsplash API.

## Potential areas of improvement

pagination not yet implemented.
The below code can exemplifies a way to work around pagination. Suggestions are welcome.\

```go
package unsplash

import (
 "context"
"fmt"
"net/http"
)

// an example of getting photos in the subsequent pages
func getPics() {
clientID := "<YOUR-CLIENT-ID>"
qParams := QueryParams(map[string]string{
"order_by": "latest",
})
c := NewClient(clientID, http.DefaultClient, &Config{AuthInHeader: true, LowContentSafety: false})
if c.Config.AuthInHeader {
    c.Config.Headers.Add("Authorization", fmt.Sprintf("Client-ID %s", clientID))
} else {
    qParams["client_id"] = clientID
}
// how I'd want it to be.
var photos []Photo
// get the first five pages
for i := 0; i <= 5; i++ {
    qParams["page"] = fmt.Sprint(i)
    pics, _ := c.GetPhotoList(context.Background(), qParams)
    photos = append(photos, pics...)
}
}

```
