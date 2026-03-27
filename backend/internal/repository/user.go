package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"liliengarten/filesharing/internal/models"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool}
}



func (r *UserRepository) Create(ctx context.Context, user models.User) error {
	_, err := r.pool.Exec(ctx,
		"INSERT INTO USERS (first_name, last_name, username, email, password) VALUES ($1, $2, $3, $4, $5)",
		user.FirstName, user.LastName, user.Username, user.Email, user.Password,
	)

	if err != nil {
		return err
	}

	return nil
}
