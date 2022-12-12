package repo

import (
	"bez/internal/entity"
	"bez/pkg/postgres"
	"context"
	"fmt"
	"log"
)

type UserRepo struct {
	*postgres.Postgres
}

func NewUserRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

func (u *UserRepo) StoreUser(ctx context.Context, us entity.User) error {
	query := `INSERT INTO users(id, email, refresh_token) VALUES($1, $2, $3)`

	rows, err := u.Pool.Query(ctx, query, us.ID, us.Email, us.RefreshToken)
	if err != nil {
		log.Println("cannot execute query")
		return fmt.Errorf("cannot execute query: %v", err)
	}
	defer rows.Close()
	return nil
}
