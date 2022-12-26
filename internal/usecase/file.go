package usecase

import (
	"bez/internal/entity"
	"context"
)

type FileUseCase struct {
	repo FileRp
}

func NewFileUseCase(repo FileRp) *FileUseCase {
	return &FileUseCase{repo: repo}
}

func (f *FileUseCase) StoreFile(ctx context.Context, fl entity.FileTorrent) error {
	return f.repo.StoreFile(ctx, fl)
}

func (f *FileUseCase) GetFileList(ctx context.Context, owner string) ([]entity.FileTorrent, error) {
	return f.repo.GetFileListByOwner(ctx, owner)
}

func (f *FileUseCase) IncrementByFile(ctx context.Context, fileID string) error {
	return f.repo.IncrementByFileID(ctx, fileID)
}
