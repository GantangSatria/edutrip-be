package service

import (
	"context"
	"errors"
	"time"

	"github.com/GantangSatria/edutrip-be/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(ctx context.Context, email, password string) (string, error)
}

type authService struct {
	adminRepo repository.AdminRepository
	jwtSecret string
}

func NewAuthService(adminRepo repository.AdminRepository, jwtSecret string) AuthService {
	return &authService{adminRepo, jwtSecret}
}

func (s *authService) Login(ctx context.Context, email, password string) (string, error) {
	admin, err := s.adminRepo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return "", errors.New("invalid email or password")
		}
		return "", errors.New("internal server error")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  admin.ID.String(),
		"name": admin.Name,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	})

	signed, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}
	return signed, nil
}