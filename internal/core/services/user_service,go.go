package services

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"hideki/internal/core/domain"
	port "hideki/internal/core/ports/repository"
	httpErrors "hideki/pkg/errors"
)

type UserService struct {
	logger     *zap.SugaredLogger
	repository port.UserRepository
}

// NewUserService creates a new user service
func NewUserService(logger *zap.SugaredLogger, repo port.UserRepository) *UserService {
	return &UserService{
		logger:     logger,
		repository: repo,
	}
}

func (svc *UserService) GetProfile(ctx context.Context, uid int16) (*domain.UserProfile, error) {
	profile, err := svc.repository.GetProfile(ctx, uid)

	if err != nil {
		svc.logger.Error(err.Error())

		select {
		case <-ctx.Done():
			return nil, httpErrors.ErrTimeout
		default:
			if errors.Is(err, httpErrors.ErrInvalidRequestBody) {
				return nil, httpErrors.ErrUserNotFound
			} else {
				return nil, err
			}
		}
	}

	return profile, nil
}
