package repo

import (
	"bez/internal/entity"
	"bez/pkg/postgres"
	"context"
)

type UserRepo struct {
	*postgres.Postgres
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (u *UserRepo) StoreUser(ctx context.Context, us entity.User) error {

}
