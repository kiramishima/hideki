package repository

import (
	"context"
	"hideki/internal/core/domain"
)

type AuthRepository interface {
	FindByCredentials(ctx context.Context, data *domain.AuthRequest) (*domain.User, error)
	Register(ctx context.Context, registerReq *domain.RegisterRequest) error
}
