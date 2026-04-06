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

func (s *BoardService) Remove(ctx context.Context, id string) error {
	err := s.repo.Remove(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *BoardService) GetPins(ctx context.Context, boardID string) ([]models.Pin, error) {
	pins, err := s.repo.GetPins(ctx, boardID)
	if err != nil {
		return nil, err
	}

	return pins, nil
}

func (s *BoardService) AddPin(ctx context.Context, boardID string, pinID string) error {
	err := s.repo.AddPin(ctx, boardID, pinID)
	if err != nil {
		return err
	}

	return nil
}

func (s *BoardService) RemovePin(ctx context.Context, boardID string, pinID string) error {
	err := s.repo.RemovePin(ctx, boardID, pinID)
	if err != nil {
		return err
	}

	return nil
}
