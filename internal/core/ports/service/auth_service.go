package service

import (
	"context"
	"hideki/internal/core/domain"
)

type AuthService interface {
	FindByCredentials(ctx context.Context, data *domain.AuthRequest) (*domain.AuthResponse, error)
	Register(ctx context.Context, registerReq *domain.RegisterRequest) error
}
