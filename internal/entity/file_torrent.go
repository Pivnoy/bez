package entity

import (
	"github.com/google/uuid"
)

type FileTorrent struct {
	ID          uuid.UUID `json:"id"`
	FileName    string    `json:"file_name"`
	FileType    string    `json:"file_type"`
	FileID      string    `json:"file_id"`
	Count       int64     `json:"count"`
	OwnerEmail  string    `json:"owner_email"`
	DownloadURL string    `json:"download_url"`
	PreviewURL  string    `json:"preview_url"`
}
