package entity

import (
	"github.com/google/uuid"
)

type FileTorrent struct {
	ID             uuid.UUID `json:"id"`
	Link           string    `json:"link"`
	DownloadNumber uint64    `json:"download_number"`
}
