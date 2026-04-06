package repository

import (
	"context"
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
