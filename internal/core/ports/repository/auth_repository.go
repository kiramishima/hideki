package repository

import (
	"context"
	"hideki/internal/core/domain"
)

type AuthRepository interface {
	Login(ctx context.Context, data *domain.AuthRequest) (*domain.User, error)
}
