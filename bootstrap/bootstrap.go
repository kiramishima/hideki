package bootstrap

import (
	"context"
	"database/sql"
	"hideki/config"
	"hideki/internal/core/services"
	"hideki/internal/database/repositories"
	"hideki/internal/handlers"
	"hideki/internal/server"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func bootstrap(
	lifecycle fx.Lifecycle,
	cfg *config.Configuration,
	logger *zap.SugaredLogger,
	db *sql.DB,
	authrepo *repositories.AuthRepository,
	authsvc *services.AuthService,
	router *chi.Mux,
	server *http.Server,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				logger.Info("Starting API")

				// logger.Info(server)

				go func() {
					if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
						logger.Fatal("failed to start server")
					}
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				return logger.Sync()
			},
		},
	)
}

var Module = fx.Options(
	config.ConfigModule,
	config.LoggerModule,
	fx.Provide(chi.NewRouter),
	server.ServerModule,
	repositories.DatabaseModule,
	services.ServicesModule,
	handlers.HandlersModule,
	fx.Invoke(bootstrap),
)
