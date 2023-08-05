package services

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"hideki/internal/adapters/database/postgresql/repository"
)

var Module = fx.Module("services",
	fx.Provide(func(logger *zap.SugaredLogger, authrepo *repository.AuthRepository) *AuthService {
		return NewAuthService(logger, authrepo)
	}),
	fx.Provide(func(logger *zap.SugaredLogger, usrrepo *repository.UserRepository) *UserService {
		return NewUserService(logger, usrrepo)
	}),
)
