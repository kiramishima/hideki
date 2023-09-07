package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"hideki/internal/core/domain"
	port "hideki/internal/core/ports/repository"
	ports "hideki/internal/core/ports/service"
	httpErrors "hideki/pkg/errors"
	"hideki/pkg/utils"
)

var _ ports.AuthService = (*AuthService)(nil)

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

// FindByCredentials To Login users
func (svc *AuthService) FindByCredentials(ctx context.Context, data *domain.AuthRequest) (*domain.AuthResponse, error) {
	data.Password = data.Hash256Password(data.Password)
	user, err := svc.repository.FindByCredentials(ctx, data)

	if err != nil {
		svc.logger.Error(err.Error())

		select {
		case <-ctx.Done():
			return nil, httpErrors.ErrTimeout
		default:
			if errors.Is(err, httpErrors.ErrInvalidRequestBody) {
				return nil, httpErrors.BadQueryParams
			} else if errors.Is(err, httpErrors.ErrUserNotFound) {
				return nil, httpErrors.ErrBadEmailOrPassword
			} else {
				return nil, httpErrors.InternalServerError
			}
		}
	}

	// Check Password
	if !data.ValidateBcryptPassword(user.Password, data.Password) {
		return nil, errors.New("Password no valid")
	}

	// Generate Token
	token, err := utils.GenerateJWT(user)
	if err != nil {
		svc.logger.Error(err.Error(), fmt.Sprintf("%T", err))
		return nil, jwt.ErrSignatureInvalid
	}

	return &domain.AuthResponse{Token: token}, nil
}

// Register repository method for create a new user.
func (svc *AuthService) Register(ctx context.Context, registerReq *domain.RegisterRequest) error {
	// Hash password
	registerReq.Password = registerReq.Hash256Password(registerReq.Password)
	// Hash Bcrypt
	registerReq.Password, _ = registerReq.BcryptPassword(registerReq.Password)
	// Call repository
	err := svc.repository.Register(ctx, registerReq)

	if err != nil {
		svc.logger.Error(err.Error())

		select {
		case <-ctx.Done():
			return httpErrors.ErrTimeout
		default:
			if errors.Is(err, httpErrors.ErrInvalidRequestBody) {
				return httpErrors.ErrBadEmailOrPassword
			} else {
				return httpErrors.ErrBadEmailOrPassword
			}
		}
	}

	return nil
}
