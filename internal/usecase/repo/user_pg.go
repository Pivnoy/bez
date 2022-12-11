package repo

import (
	"bez/internal/entity"
	"context"
	"gorm.io/gorm"
)

type UserRepo struct {
	*gorm.DB
}

func NewUserRepo(pg *gorm.DB) *UserRepo {
	return &UserRepo{pg}
}

func (u *UserRepo) StoreUser(ctx context.Context, us entity.User) error {
	return u.WithContext(ctx).Create(&us).Error
}
