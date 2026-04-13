package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"liliengarten/filesharing/internal/models"
	"liliengarten/filesharing/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
	pinRepo  *repository.PinRepository
}

func NewUserService(userRepo *repository.UserRepository, pinRepo *repository.PinRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
		pinRepo:  pinRepo,
	}
}

func generateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"sub": fmt.Sprint(userID),
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("TOKEN_KEY")))
}

func (s *UserService) Register(ctx context.Context, user models.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashed)

	err = s.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Login(ctx context.Context, user models.UserLogin) (string, error) {
	repo_user, err := s.userRepo.Login(ctx, user.Email)
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

func (s *UserService) GetProfile(ctx context.Context, userID string) ([]models.User, error) {
	profile, err := s.userRepo.GetById(ctx, userID)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (s *UserService) GetLikes(ctx context.Context, page string) ([]models.Pin, error) {
	if page == "" {
		page = "1"
	}

	pins, err := s.pinRepo.GetLikes(ctx, page)
	if err != nil {
		return nil, err
	}

	return pins, nil
}
