package service

import (
	"context"
	"hideki/internal/core/domain"
)

type AuthService interface {
	Login(ctx context.Context, data *domain.AuthRequest) (*domain.AuthResponse, error)
}
