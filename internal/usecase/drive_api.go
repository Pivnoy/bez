package usecase

import (
	"bez/internal/entity"
	"context"
	"fmt"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"net/http"
)

type DriveAPIUseCase struct {
	srv *drive.Service
}

func NewDriveAPI() *DriveAPIUseCase {
	return &DriveAPIUseCase{}
}

func (d *DriveAPIUseCase) UserDrive(ctx context.Context, client *http.Client) error {
	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return fmt.Errorf("cannot create drive service: %v", err)
	}
	d.srv = srv
	return nil
}

func (d *DriveAPIUseCase) GetPersonalInfo() (*entity.PersonalInfo, error) {
	aboutSrv := drive.NewAboutService(d.srv)
	res, err := aboutSrv.Get().Do(googleapi.QueryParameter("fields", "user,storageQuota"))
	if err != nil {
		return nil, fmt.Errorf("cannot get about user info: %v", err)
	}
	return &entity.PersonalInfo{
		DisplayName: res.User.DisplayName,
		Picture:     res.User.PhotoLink,
		Email:       res.User.EmailAddress,
	}, nil
}
