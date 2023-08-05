package services

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	repositories2 "hideki/internal/adapters/database/repositories"
)

var Module = fx.Module("services",
	fx.Provide(func(logger *zap.SugaredLogger, authrepo *repositories2.AuthRepository) *AuthService {
		return NewAuthService(logger, authrepo)
	}),
	fx.Provide(func(logger *zap.SugaredLogger, usrrepo *repositories2.UserRepository) *UserService {
		return NewUserService(logger, usrrepo)
	}),
)
