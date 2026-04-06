package service

import (
	"context"
	"liliengarten/filesharing/internal/models"
	"liliengarten/filesharing/internal/repository"
)

type BoardService struct {
	repo *repository.BoardRepository
}

func NewBoardService(r *repository.BoardRepository) *BoardService {
	return &BoardService{repo: r}
}

func (s *BoardService) Index(ctx context.Context) ([]models.Board, error) {
	boards, err := s.repo.Index(ctx)
	if err != nil {
		return nil, err
	}

	return boards, nil
}

func (s *BoardService) Create(ctx context.Context, board *models.Board) error {
	err := s.repo.Create(ctx, board)
	if err != nil {
		return err
	}

	return nil
}
