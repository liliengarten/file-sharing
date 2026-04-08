package service

import (
	"context"
	"errors"
	"liliengarten/filesharing/internal/models"
	"liliengarten/filesharing/internal/repository"
)

type BoardService struct {
	boardRepo *repository.BoardRepository
	userRepo  *repository.UserRepository
}

func NewBoardService(br *repository.BoardRepository, ur *repository.UserRepository) *BoardService {
	return &BoardService{
		boardRepo: br,
		userRepo:  ur,
	}
}

func (s *BoardService) Index(ctx context.Context) ([]models.Board, error) {
	boards, err := s.boardRepo.Index(ctx)
	if err != nil {
		return nil, err
	}

	return boards, nil
}

func (s *BoardService) GetBoard(ctx context.Context, boardID string) ([]models.Board, error) {
	board, err := s.boardRepo.GetById(ctx, boardID)
	if err != nil {
		return nil, err
	}

	return board, nil
}

func (s *BoardService) Create(ctx context.Context, board *models.Board) error {
	err := s.boardRepo.Create(ctx, board)
	if err != nil {
		return err
	}

	return nil
}

func (s *BoardService) Remove(ctx context.Context, id string) error {
	err := s.boardRepo.Remove(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *BoardService) GetPins(ctx context.Context, boardID string) ([]models.Pin, error) {
	pins, err := s.boardRepo.GetPins(ctx, boardID)
	if err != nil {
		return nil, err
	}

	return pins, nil
}

func (s *BoardService) AddPin(ctx context.Context, boardID string, pinID string) error {
	err := s.boardRepo.AddPin(ctx, boardID, pinID)
	if err != nil {
		return err
	}

	return nil
}

func (s *BoardService) RemovePin(ctx context.Context, boardID string, pinID string) error {
	err := s.boardRepo.RemovePin(ctx, boardID, pinID)
	if err != nil {
		return err
	}

	return nil
}

func (s *BoardService) GetAuthors(ctx context.Context, boardID string) ([]models.BoardAuthor, error) {
	authors, err := s.boardRepo.GetAuthors(ctx, boardID)
	if err != nil {
		return nil, err
	}

	return authors, nil
}

func (s *BoardService) AddAuthor(ctx context.Context, boardID string, userID string) error {
	if userID == ctx.Value("user") {
		return errors.New("can't add yourself as author")
	}

	_, err := s.userRepo.GetById(ctx, userID)
	if err != nil {
		return err
	}

	_, err = s.boardRepo.GetById(ctx, boardID)
	if err != nil {
		return err
	}

	err = s.boardRepo.AddAuthor(ctx, boardID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *BoardService) RemoveAuthor(ctx context.Context, boardID string, userID string) error {
	_, err := s.userRepo.GetById(ctx, userID)
	if err != nil {
		return err
	}

	err = s.boardRepo.RemoveAuthor(ctx, boardID, userID)
	if err != nil {
		return err
	}

	return nil
}
