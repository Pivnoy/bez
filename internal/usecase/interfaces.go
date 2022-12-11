package usecase

import (
	"bez/internal/entity"
	"context"
	"golang.org/x/oauth2"
	"net/http"
)

type (
	UserRp interface {
		StoreUser(context.Context, entity.User) error
	}

	User interface {
		CreateUser(context.Context, string, string) error
	}

	GoogleAPI interface {
		CreateRegLink() string
		CreateUserToken(ctx context.Context, authCode string) (*oauth2.Token, error)
		CreateClient(ctx context.Context, token *oauth2.Token) (*http.Client, error)
	}

	DriveAPI interface {
		UserDrive(ctx context.Context, client *http.Client) error
		GetPersonalInfo() (*entity.PersonalInfo, error)
		GetFileList()
	}
)
