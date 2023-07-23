package service

import (
	"context"
	"hideki/internal/core/domain"
)

type UserService interface {
	GetProfile(ctx context.Context, uid int16) (*domain.User, error)
}
