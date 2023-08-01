package services

import (
	"go.uber.org/zap"
	"hideki/internal/database/repositories"

	"go.uber.org/fx"
)

var Module = fx.Module("services",
	fx.Provide(func(logger *zap.SugaredLogger, authrepo *repositories.AuthRepository) *AuthService {
		return NewAuthService(logger, authrepo)
	}),
	fx.Provide(func(logger *zap.SugaredLogger, usrrepo *repositories.UserRepository) *UserService {
		return NewUserService(logger, usrrepo)
	}),
)
