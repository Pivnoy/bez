package repo

import (
	"bez/internal/entity"
	"bez/pkg/postgres"
	"context"
	"fmt"
)

type FileTorrentRepo struct {
	*postgres.Postgres
}

func NewFileTorrentRepo(pg *postgres.Postgres) *FileTorrentRepo {
	return &FileTorrentRepo{pg}
}

func (f *FileTorrentRepo) StoreFile(ctx context.Context, fl entity.FileTorrent) error {
	query := `INSERT INTO file(id, file_name, file_type, file_id, count, owner_email) VALUES($1, $2, $3, $4, $5, $6)`

	rows, err := f.Pool.Query(ctx, query, fl.ID, fl.FileName, fl.FileType, fl.FileID, fl.Count, fl.OwnerEmail)
	if err != nil {
		return fmt.Errorf("cannot execute query: %v", err)
	}
	defer rows.Close()
	return nil
}

func (f *FileTorrentRepo) GetFileListByOwner(ctx context.Context, owner string) ([]entity.FileTorrent, error) {
	query := `SELECT * FROM file WHERE owner = $1`

	rows, err := f.Pool.Query(ctx, query, owner)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %v", err)
	}
	defer rows.Close()
	var files []entity.FileTorrent

	for rows.Next() {
		var fl entity.FileTorrent
		err = rows.Scan(
			&fl.ID,
			&fl.FileName,
			&fl.FileType,
			&fl.FileID,
			&fl.Count,
			&fl.OwnerEmail)
		if err != nil {
			return nil, fmt.Errorf("cannot parse file torrent: %v", err)
		}
		files = append(files, fl)
	}
	return files, nil
}
