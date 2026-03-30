package service

import (
	"context"
	"time"
	"liliengarten/filesharing/internal/repository"
	"liliengarten/filesharing/internal/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
)



type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}



func generateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("12345")) //ключ в енв
}

func (s *UserService) Register(ctx context.Context, user models.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	user.Password = string(hashed)

	err = s.repo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Login(ctx context.Context, user models.UserLogin) (string, error) {
	repo_user, err := s.repo.Login(ctx, user.Email)
	if err != nil {
		return "", err
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(repo_user.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	
	token, err := generateToken(repo_user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
