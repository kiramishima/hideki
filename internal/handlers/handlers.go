package handlers

import (
	"hideki/internal/core/services"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var HandlersModule = fx.Module("handlers",
	fx.Invoke(func(r *chi.Mux, logger *zap.SugaredLogger, svc *services.AuthService) {
		NewAuthHandlers(r, logger, svc)
	}),
	fx.Invoke(func(logger *zap.SugaredLogger, r *chi.Mux) {
		NewHealthHandlers(logger, r)
	}),
)
