package entity

import (
	"github.com/google/uuid"
)

type User struct {
	ID      uuid.UUID `json:"id"`
	Email   string    `json:"email"`
	Refresh string    `json:"refresh_token"`
}
