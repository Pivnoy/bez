package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FileTorrent struct {
	gorm.Model
	ID             uuid.UUID `json:"id"`
	Link           string    `json:"link"`
	DownloadNumber uint64    `json:"download_number"`
}
