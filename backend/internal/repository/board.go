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

	rows, err = r.pool.Query(ctx, "SELECT * FROM boards WHERE id IN $1", boardIds)
	if err != nil {
		return nil, err
	}

	boards, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Board])
	if err != nil {
		return nil, err
	}

	return boards, nil
}
