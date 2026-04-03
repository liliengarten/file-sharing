package repository

import (
	"context"
	"errors"
	"liliengarten/filesharing/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PinRepository struct {
	pool *pgxpool.Pool
}

func NewPinRepository(pool *pgxpool.Pool) *PinRepository {
	return &PinRepository{pool}
}

func (r *PinRepository) Index(ctx context.Context) ([]models.Pin, error) {
	rows, err := r.pool.Query(ctx, "SELECT * FROM pins")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pins, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Pin])
	if err != nil {
		return nil, err
	}

	return pins, nil
}

func (r *PinRepository) SavePin(ctx context.Context, pin *models.Pin) error {
	_, err := r.pool.Exec(ctx, "INSERT INTO pins (image, description) VALUES ($1, $2)", pin.Image, pin.Description)
	if err != nil {
		return err
	}

	return nil
}

func (r *PinRepository) GetById(ctx context.Context, id int) (*models.Pin, error) {
	pin := &models.Pin{}

	row := r.pool.QueryRow(ctx, "SELECT * FROM pins WHERE image = $1", id)

	err := row.Scan(&pin.Image, &pin.Description)
	if err != nil {
		return nil, err
	}

	return pin, nil
}

func (r *PinRepository) Remove(ctx context.Context, id string) error {
	commandTag, err := r.pool.Exec(ctx, "DELETE FROM pins WHERE id = $1", id)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		err = errors.New("Pin not found")
		return err
	}

	return nil
}
