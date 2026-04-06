package repository

import (
	"context"
	"errors"
	"liliengarten/filesharing/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BoardRepository struct {
	pool *pgxpool.Pool
}

func NewBoardRepository(pool *pgxpool.Pool) *BoardRepository {
	return &BoardRepository{pool: pool}
}

func (r *BoardRepository) HasRights(ctx context.Context, boardID string) error {
	var exists bool
	err := r.pool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM user_boards WHERE user_id = $1 and board_id = $2)", ctx.Value("user"), boardID).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Board not found")
	}

	return nil
}

func (r *BoardRepository) Index(ctx context.Context) ([]models.Board, error) {
	rows, err := r.pool.Query(ctx, "SELECT board_id FROM user_boards WHERE user_id = $1", ctx.Value("user"))
	if err != nil {
		return nil, err
	}

	boardIds, err := pgx.CollectRows(rows, pgx.RowTo[int])
	if err != nil {
		return nil, err
	}

	rows, err = r.pool.Query(ctx, "SELECT * FROM boards WHERE id = ANY($1)", boardIds)
	if err != nil {
		return nil, err
	}

	boards, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Board])
	if err != nil {
		return nil, err
	}

	return boards, nil
}

func (r *BoardRepository) Create(ctx context.Context, board *models.Board) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}

	err = tx.QueryRow(ctx, "INSERT INTO boards (name, description, private) VALUES ($1, $2, $3) RETURNING id", board.Name, board.Description, board.Private).Scan(&board.ID)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	_, err = tx.Exec(ctx, "INSERT INTO user_boards (user_id, board_id, role) VALUES ($1, $2, $3)", ctx.Value("user"), board.ID, 1)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}

func (r *BoardRepository) Remove(ctx context.Context, id string) error {
	commandTag, err := r.pool.Exec(ctx, "DELETE FROM user_boards WHERE board_id = $1 and user_id = $2 and role = $3", id, ctx.Value("user"), 1)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("Board not found")
	}

	return nil
}

func (r *BoardRepository) GetPins(ctx context.Context, boardID string) ([]models.Pin, error) {
	err := r.HasRights(ctx, boardID)
	if err != nil {
		return nil, err
	}

	rows, err := r.pool.Query(ctx, "SELECT pin_id FROM board_pins WHERE board_id = $1", boardID)
	if err != nil {
		return nil, err
	}
	pinIds, err := pgx.CollectRows(rows, pgx.RowTo[int])
	if err != nil {
		return nil, err
	}

	rows, err = r.pool.Query(ctx, "SELECT * FROM pins WHERE id = ANY($1)", pinIds)
	if err != nil {
		return nil, err
	}
	pins, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Pin])
	if err != nil {
		return nil, err
	}

	return pins, nil
}

func (r *BoardRepository) AddPin(ctx context.Context, boardID string, pinID string) error {
	err := r.HasRights(ctx, boardID)
	if err != nil {
		return err
	}

	_, err = r.pool.Exec(ctx, "INSERT INTO board_pins (pin_id, board_id) VALUES ($1, $2)", pinID, boardID)
	if err != nil {
		return err
	}

	return nil
}

func (r *BoardRepository) RemovePin(ctx context.Context, boardID string, pinID string) error {
	err := r.HasRights(ctx, boardID)
	if err != nil {
		return err
	}

	commandTag, err := r.pool.Exec(ctx, "DELETE FROM board_pins WHERE pin_id = $1 and board_id = $2", pinID, boardID)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("Pin not found")
	}

	return nil
}
