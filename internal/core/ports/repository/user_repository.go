package repository

import (
	"context"
	"hideki/internal/core/domain"
)

type UserRepository interface {
	GetProfile(ctx context.Context, uid int16) (*domain.User, error)
}
