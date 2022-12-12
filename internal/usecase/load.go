package usecase

import (
	"bez/internal/entity"
	"context"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"net/http"
)

type LoadUseCase struct {
	file File
	srv  Service
	ggl  GoogleAPI
	drv  DriveAPI
}

func NewLoadUseCase(fl File, srv Service, ggl GoogleAPI, drv DriveAPI) *LoadUseCase {
	return &LoadUseCase{fl, srv, ggl, drv}
}

func (l *LoadUseCase) Preload(ctx context.Context) error {
	srv, err := l.srv.GetServices(ctx)
	var fileStorage []entity.FileTorrent
	if err != nil {
		return fmt.Errorf("cannot get service accounts: %v", err)
	}
	var clients []*http.Client
	for _, sr := range srv {
		client, err := l.ggl.CreateClient(ctx, &oauth2.Token{RefreshToken: sr.RefreshToken, TokenType: "Bearer"})
		if err != nil {
			return fmt.Errorf("cannot create http clients")
		}
		clients = append(clients, client)
	}
	for _, cl := range clients {
		driveImpl, err := l.drv.UserDrive(ctx, cl)
		if err != nil {
			return fmt.Errorf("cannot create drive Impl: %v", err)
		}
		personal, err := l.drv.GetPersonalInfo(driveImpl)
		if err != nil {
			return fmt.Errorf("cannot get personal info: %v", err)
		}
		flLs, err := l.drv.GetFileList(driveImpl)
		if err != nil {
			return err
		}
		for _, vl := range flLs {
			fileStorage = append(fileStorage, entity.FileTorrent{
				ID:         uuid.New(),
				FileName:   vl.Name,
				FileType:   vl.MimeType,
				FileID:     vl.Id,
				Count:      0,
				OwnerEmail: personal.Email})
		}
	}
	fmt.Println(len(fileStorage))
	for _, fl := range fileStorage {
		err := l.file.StoreFile(ctx, fl)
		if err != nil {
			return err
		}
	}
	fmt.Println("done")
	return nil
}
