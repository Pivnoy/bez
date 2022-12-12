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

type DriveAPIUseCase struct{}

func NewDriveAPI() *DriveAPIUseCase {
	return &DriveAPIUseCase{}
}

func (d *DriveAPIUseCase) UserDrive(ctx context.Context, client *http.Client) (*drive.Service, error) {
	return drive.NewService(ctx, option.WithHTTPClient(client))
}

func (d *DriveAPIUseCase) GetPersonalInfo(srv *drive.Service) (*entity.PersonalInfo, error) {
	aboutSrv := drive.NewAboutService(srv)
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

func (d *DriveAPIUseCase) GetFileList(srv *drive.Service) ([]*drive.File, error) {
	fl, err := srv.Files.List().Fields("files(id, name, mimeType, webContentLink, webViewLink)").Do()
	if err != nil {
		return nil, fmt.Errorf("cannot get files: %v", err)
	}
	return fl.Files, nil
}
