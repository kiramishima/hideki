package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"hideki/internal/core/domain"
	port "hideki/internal/core/ports/repository"
	httpErrors "hideki/pkg/errors"
	"hideki/pkg/utils"
)

type AuthService struct {
	logger     *zap.SugaredLogger
	repository port.AuthRepository
}

// NewAuthService creates a new auth service
func NewAuthService(logger *zap.SugaredLogger, repo port.AuthRepository) *AuthService {
	return &AuthService{
		logger:     logger,
		repository: repo,
	}
}

// Login To Login users
func (svc *AuthService) Login(ctx context.Context, data *domain.AuthRequest) (*domain.AuthResponse, error) {
	user, err := svc.repository.Login(ctx, data)

	if err != nil {
		svc.logger.Error(err.Error())

		select {
		case <-ctx.Done():
			return nil, httpErrors.ErrTimeout
		default:
			if errors.Is(err, httpErrors.ErrInvalidRequestBody) {
				return nil, httpErrors.ErrBadEmailOrPassword
			} else {
				return nil, httpErrors.ErrBadEmailOrPassword
			}
		}
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		svc.logger.Error(err.Error(), fmt.Sprintf("%T", err))
		return nil, jwt.ErrSignatureInvalid
	}

	return &domain.AuthResponse{Token: token}, nil
}
