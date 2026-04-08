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

func (r *PinRepository) SavePin(ctx context.Context, pin *models.Pin, userID string) error {
	_, err := r.pool.Exec(ctx, "INSERT INTO pins (owner_id, image, description) VALUES ($1, $2, $3)", userID, pin.Image, pin.Description)
	if err != nil {
		return err
	}

	return nil
}

func (r *PinRepository) GetById(ctx context.Context, id string) ([]models.Pin, error) {
	rows, err := r.pool.Query(ctx, "SELECT * FROM pins WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pin, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Pin])
	if err != nil {
		return nil, err
	}

	if len(pin) == 0 {
		return nil, errors.New("pin not found")
	}

	return pin, nil
}

func (r *PinRepository) Update(ctx context.Context, pin *models.Pin) error {
	commandTag, err := r.pool.Exec(ctx, "UPDATE pins SET description = $1 WHERE id = $2 and owner_id = $3", pin.Description, pin.ID, ctx.Value("user"))

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		err = errors.New("pin not found")
		return err
	}

	return nil
}

func (r *PinRepository) Remove(ctx context.Context, pinID string, userID string) error {
	commandTag, err := r.pool.Exec(ctx, "DELETE FROM pins WHERE id = $1 and owner_id = $2", pinID, userID)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		err = errors.New("pin not found")
		return err
	}

	return nil
}

func (r *PinRepository) LikePin(ctx context.Context, pinID string) error {
	_, err := r.pool.Exec(ctx, "INSERT INTO liked_pins (user_id, pin_id) VALUES ($1, $2)", pinID, ctx.Value("user"))
	if err != nil {
		return err
	}

	return nil
}

func (r *PinRepository) UnlikePin(ctx context.Context, pinID string) error {
	commandTag, err := r.pool.Exec(ctx, "DELETE FROM liked_pins WHERE user_id = $1 and pin_id = $", pinID, ctx.Value("user"))
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		err = errors.New("record not found")
	}

	return nil
}
