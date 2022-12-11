package usecase

import (
	"bez/internal/entity"
	"context"
	"github.com/google/uuid"
)

type UserUseCase struct {
	repo UserRp
}

func NewUserUseCase(repo UserRp) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (u *UserUseCase) CreateUser(ctx context.Context, email, token string) error {
	us := entity.User{ID: uuid.New(), Email: email, RefreshToken: token}
	return u.repo.StoreUser(ctx, us)
}
