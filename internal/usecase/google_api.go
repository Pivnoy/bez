package usecase

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"log"
	"net/http"
	"net/url"
)

type GoogleAPIUseCase struct {
	config *oauth2.Config
}

func NewGoogleAPIUseCase(cred []byte) *GoogleAPIUseCase {
	return &GoogleAPIUseCase{
		config: createConfig(cred),
	}
}

func createConfig(cred []byte) *oauth2.Config {
	config, err := google.ConfigFromJSON(cred, drive.DriveMetadataReadonlyScope)
	if err != nil {
		log.Println("Cannot create config")
		log.Fatalln(err)
	}
	return config
}

func (c *GoogleAPIUseCase) CreateRegLink() string {
	c.config.Scopes = []string{
		"https://www.googleapis.com/auth/drive",
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email",
	}
	return c.config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
}

func (c *GoogleAPIUseCase) CreateUserToken(ctx context.Context, authCode string) (*oauth2.Token, error) {
	authCodeURL, err := url.QueryUnescape(authCode)
	if err != nil {
		return nil, fmt.Errorf("cannot unescape URL")
	}
	token, err := c.config.Exchange(ctx, authCodeURL)
	if err != nil {
		log.Println("cannot create token")
		return nil, err
	}
	return token, nil
}

// CreateClient add url regex
func (c *GoogleAPIUseCase) CreateClient(ctx context.Context, token *oauth2.Token) (*http.Client, error) {
	client := c.config.Client(ctx, token)
	return client, nil
}
