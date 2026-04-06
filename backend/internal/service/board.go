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
