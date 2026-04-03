package service

import (
	"context"
	"io"
	"liliengarten/filesharing/internal/models"
	"liliengarten/filesharing/internal/repository"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type PinService struct {
	repo *repository.PinRepository
}

func NewPinService(repo *repository.PinRepository) *PinService {
	return &PinService{repo}
}

func (s *PinService) Index(ctx context.Context) ([]models.Pin, error) {
	pins, err := s.repo.Index(ctx)
	if err != nil {
		return nil, err
	}

	return pins, nil
}

func (s *PinService) SavePin(ctx context.Context, pin *models.Pin, file multipart.File, header *multipart.FileHeader) error {
	filename := uuid.New().String() + filepath.Ext(header.Filename)

	pin.Image = "uploads/" + filename

	dst, err := os.Create("../../uploads/" + filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return err
	}

	err = s.repo.SavePin(ctx, pin)

	return nil
}

func (s *PinService) Remove(ctx context.Context, id string) error {
	err := s.repo.Remove(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
