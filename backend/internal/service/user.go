package service

import (
	"context"
	"liliengarten/filesharing/internal/repository"
	"liliengarten/filesharing/internal/models"
	"golang.org/x/crypto/bcrypt"
)



type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}



func (s *UserService) Register(ctx context.Context, user models.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashed)

	err = s.repo.Create(ctx, user)

	if err != nil {
		return err
	}

	return nil
}
