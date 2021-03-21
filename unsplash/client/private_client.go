package client

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
)

// AuthScopes lists all scopes used in a particular request
type AuthScopes []string

// Contains returns true if string is found in the the underlying slice structure
func (a AuthScopes) Contains(elem string) bool {
	for _, val := range a {
		if val == elem {
			return true
		}
	}
	return false
}

// NewAuthScopes constructs a new AuthScopes object, with an initial
// `public` element
func NewAuthScopes(scopes ...string) *AuthScopes {
	as := AuthScopes{"public"}
	as = append(as, scopes...)
	return &as
}

// NewUnsplashOauthConfig constructs an *oauth.Config object
// with all necessary credentials
func NewUnsplashOauthConfig(clientID, clientSecret, redirectURI string, scopes *AuthScopes) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURI,
		Scopes:       *scopes,
		Endpoint:     unsplashEndpoint,
	}
}

var unsplashEndpoint = oauth2.Endpoint{
	AuthURL:  AuthCodeEndpoint,
	TokenURL: AuthTokenEndpoint,
}

// NewPrivateAuthClient initializes a new client that has been authorised
// for private actions.
func NewPrivateAuthClient(clientID, clientSecret, redirectURI string, as *AuthScopes, config *Config) (*Client, error) {
	conf := NewUnsplashOauthConfig(clientID, clientSecret, redirectURI, as)
	link := conf.AuthCodeURL("state", oauth2.AccessTypeOnline)

	// User instructions to get authorization code
	fmt.Printf("Navigate to:\n%s\n\n", link)

	fmt.Println("You will redirected to the redirect uri, whose link will have a `code` query parameter")
	fmt.Println("Paste the authorization code here: ")
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return nil, err
	}
	// get token
	tok, err := conf.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}
	// create client
	client := conf.Client(context.Background(), tok)

	c := &Client{ClientID: clientID,
		HTTPClient: client,
		Config:     config,
		Private:    true,
		AuthScopes: as}
	return c, nil
}
