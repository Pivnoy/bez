package entity

import (
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"primaryKey" json:"id"`
	Email        string    `gorm:"email" json:"email"`
	RefreshToken string    `gorm:"refresh_token" json:"refresh_token"`
}
