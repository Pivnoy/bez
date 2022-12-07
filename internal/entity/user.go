package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email   string `json:"email"`
	Refresh string `json:"refresh"`
}
