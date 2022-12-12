package entity

import "github.com/google/uuid"

type ServiceAccount struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	RefreshToken string    `json:"refresh_token"`
	StorageLimit int64     `json:"storage_limit"`
	StorageUsage int64     `json:"storage_usage"`
}
