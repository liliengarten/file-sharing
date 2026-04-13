package repository

import (
	"context"
	"errors"
	"liliengarten/filesharing/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool}
}

func (r *UserRepository) GetById(ctx context.Context, id string) ([]models.User, error) {
	rows, err := r.pool.Query(ctx, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		return nil, err
	}
	if len(user) == 0 {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (r *UserRepository) Create(ctx context.Context, user models.User) error {
	var exists bool
	err := r.pool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", user.Email).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("user with this email already exists")
	}

	err = r.pool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", user.Username).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("user with this username already exists")
	}

	_, err = r.pool.Exec(ctx,
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
