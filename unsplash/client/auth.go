package client

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// AuthScopes lists all scopes used in a particular request
type AuthScopes []string

// AuthResponse defines fields gotten when authenticanting Unsplash using OAuth2
type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	CreatedAt   int    `json:"created_at"`
}

// Contains returns true if string is found in the the underlying slice structure
func (a AuthScopes) Contains(elem string) bool {
	for _, val := range a {
		if val == elem {
			return true
		}
	}
	return false
}

func (a *AuthScopes) String() string {
	return strings.Join(*a, "+")
}

// NewAuthScopes constructs a new AuthScopes object, with an initial
// `public` element
func NewAuthScopes(scopes ...string) *AuthScopes {
	as := AuthScopes{"public"}
	as = append(as, scopes...)
	return &as
}

// NewPrivateAuthClient initializes a new client that has been authorised
// for private actions.
func NewPrivateAuthClient(ctx context.Context, clientID, clientSecret, redirectURI string, client *http.Client, config *Config, as *AuthScopes) (*Client, error) {
	if client == nil {
		client = http.DefaultClient
	}
	c := &Client{ClientID: clientID,
		HTTPClient: client,
		Config:     config,
		Private:    true,
		AuthScopes: *as}
	// get authorization code
	authCode, err := c.authGetCode(ctx, clientID, redirectURI)
	if err != nil {
		return nil, err
	}
	// get access token
	authResponse, err := c.authGetAccessToken(ctx, clientID, clientSecret, redirectURI, authCode)
	if err != nil {
		return nil, err
	}
	accessToken := authResponse.AccessToken
	authHeader := fmt.Sprintf("Bearer %s", accessToken)
	c.Config.Headers.Add("Authentication", authHeader)
	c.Config.Headers.Add("Accept-Version", "v1") // Add api version
	// Unsplash strongly encourages a specific request of the api version
	// do sth to get access token
	return c, nil
}

// returns the authorization code used in the subsequent POST request to get access token
// returns 0 if an error is encountered
func (c *Client) authGetCode(ctx context.Context, clientID, redirectURI string) (string, error) {
	queryParams := QueryParams(map[string]string{
		"client_id":     clientID,
		"redirect_uri":  redirectURI,
		"response_type": "code", // The access response type you are requesting. The authorization workflow Unsplash supports requires the value “code” here.
	})
	link, err := buildURL(AuthCodeEndpoint, queryParams)
	if err != nil {
		return "", err
	}
	link += fmt.Sprintf("&scope=%s", c.AuthScopes.String())

	// User instructions to get authorization code
	fmt.Printf("Navigate to:\n%s\n\n", link)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("You will redirected to the redirect uri, whose link will have a `code` query parameter")
	fmt.Println("Paste the authorization code here: ")
	authCode, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	if authCode == "" {
		return "", ErrAuthCodeEmpty
	}
	return authCode, nil
}

func (c *Client) authGetAccessToken(ctx context.Context, clientID, clientSecret, redirectURI string, authCode string) (*AuthResponse, error) {
	qParams := QueryParams{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"redirect_uri":  redirectURI,
		"code":          strings.TrimSpace(authCode),
		"grant_type":    "authorization_code",
	}
	link, err := buildURL(AuthTokenEndpoint, qParams)
	if err != nil {
		return nil, err
	}
	resp, err := c.postHTTP(ctx, link, nil)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, ErrStatusCode{resp.StatusCode, getErrReasons(resp)}
	}
	defer resp.Body.Close()

	// parse json response
	var ar AuthResponse
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = parseJSON(data, &ar)
	if err != nil {
		return nil, err
	}
	return &ar, nil
}
