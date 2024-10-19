// internal/youtube/auth.go

package youtube

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

const tokenFile = "token.json"

// getConfig returns the OAuth2 configuration
func getConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("YOUTUBE_CLIENT_ID"),
		ClientSecret: os.Getenv("YOUTUBE_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/callback",
		Scopes: []string{
			youtube.YoutubeReadonlyScope,
		},
		Endpoint: google.Endpoint,
	}
}

// getTokenFromWeb initiates the OAuth2 flow and returns the token
func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state")
	fmt.Printf("Go to this URL in your browser: \n%v\n", authURL)

	var code string
	codeChan := make(chan string)
	server := &http.Server{Addr: ":8080"}

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code = r.URL.Query().Get("code")
		codeChan <- code
		fmt.Fprintf(w, "Authorization successful! You can close this window.")
		go func() {
			time.Sleep(1 * time.Second)
			server.Shutdown(context.Background())
		}()
	})

	go server.ListenAndServe()

	code = <-codeChan

	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve token from web: %v", err)
	}
	return token, nil
}

// saveToken saves the token to a file
func saveToken(token *oauth2.Token) error {
	f, err := os.OpenFile(tokenFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("unable to save token file: %v", err)
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(token)
}

// loadToken loads the token from a file
func loadToken() (*oauth2.Token, error) {
	f, err := os.Open(tokenFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	token := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(token)
	return token, err
}

// GetAuthenticatedClient returns an authenticated YouTube service
func GetAuthenticatedClient() (*youtube.Service, error) {
	ctx := context.Background()
	config := getConfig()

	token, err := loadToken()
	if err != nil {
		token, err = getTokenFromWeb(config)
		if err != nil {
			return nil, err
		}
		if err := saveToken(token); err != nil {
			return nil, err
		}
	}

	client := config.Client(ctx, token)
	service, err := youtube.New(client)
	if err != nil {
		return nil, fmt.Errorf("error creating YouTube client: %v", err)
	}

	return service, nil
}
