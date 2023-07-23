package services

import (
	"context"
	"hideki/internal/core/domain"
	port "hideki/internal/core/ports/repository"
)

type AuthService struct {
	repository port.AuthRepository
}

// NewAuthService creates a new auth service
func NewAuthService(repo port.AuthRepository) *AuthService {
	return &AuthService{
		repository: repo,
	}
}

// Login To Login users
func (svc *AuthService) Login(ctx context.Context, data *domain.AuthRequest) (*domain.User, error) {
	user, err := svc.repository.Login(ctx, data)

	if err != nil {
		return nil, err
	}

	return user, nil
}
