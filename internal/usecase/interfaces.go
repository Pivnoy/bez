package usecase

import (
	"bez/internal/entity"
	"context"
	"golang.org/x/oauth2"
	"google.golang.org/api/drive/v3"
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
		UserDrive(ctx context.Context, client *http.Client) (*drive.Service, error)
		GetPersonalInfo(*drive.Service) (*entity.PersonalInfo, error)
		GetFileList(srv *drive.Service) ([]*drive.File, error)
		CopyFile(srv *drive.Service, fileID string) (*drive.File, error)
	}

	ServiceRp interface {
		GetAllServices(context.Context) ([]entity.ServiceAccount, error)
	}

	Service interface {
		GetServices(context.Context) ([]entity.ServiceAccount, error)
	}

	FileRp interface {
		StoreFile(ctx context.Context, fl entity.FileTorrent) error
		GetFileListByOwner(ctx context.Context, owner string) ([]entity.FileTorrent, error)
		IncrementByFileID(ctx context.Context, fileID string) error
	}

	File interface {
		StoreFile(ctx context.Context, fl entity.FileTorrent) error
		GetFileList(ctx context.Context, owner string) ([]entity.FileTorrent, error)
		IncrementByFile(ctx context.Context, fileID string) error
	}
)
