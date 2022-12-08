package usecase

import (
	"context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"log"
	"net/http"
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
	return c.config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
}

func (c *GoogleAPIUseCase) CreateUserToken(ctx context.Context, authCode string) (*oauth2.Token, error) {
	token, err := c.config.Exchange(ctx, authCode)
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
