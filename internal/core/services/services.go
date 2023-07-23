package services

import (
	"hideki/internal/database/repositories"

	"go.uber.org/fx"
)

var ServicesModule = fx.Module("services",
	fx.Provide(func(authrepo *repositories.AuthRepository) *AuthService {
		return NewAuthService(authrepo)
	}),
)
