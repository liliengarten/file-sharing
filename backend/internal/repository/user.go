package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
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

func (r *UserRepository) Login(ctx context.Context, email string) (models.User, error) {
	var user models.User

	err := r.pool.QueryRow(ctx, "SELECT id, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.User{}, err
		}
	}

	return user, nil
}
