package handlers

import (
	"github.com/unrolled/render"
	"hideki/internal/core/services"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("handlers",
	fx.Invoke(func(r *chi.Mux, logger *zap.SugaredLogger, svc *services.AuthService, render *render.Render) {
		NewAuthHandlers(r, logger, svc, render)
	}),
	fx.Invoke(func(r *chi.Mux, logger *zap.SugaredLogger, svc *services.UserService, render *render.Render) {
		NewUserHandlers(r, logger, svc, render)
	}),
	fx.Invoke(func(logger *zap.SugaredLogger, r *chi.Mux, render *render.Render) {
		NewHealthHandlers(logger, r, render)
	}),
)
