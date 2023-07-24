package bootstrap

import (
	"context"
	"hideki/config"
	"hideki/internal/core/services"
	"hideki/internal/database/repositories"
	"hideki/internal/handlers"
	"hideki/internal/server"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func bootstrap(
	lifecycle fx.Lifecycle,
	logger *zap.SugaredLogger,
	server *http.Server,
	router *chi.Mux,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				logger.Info("Starting API")
				// Middlewares
				/*router.Use(middleware.Timeout(60 * time.Second))
				router.Use(middleware.RequestID)
				router.Use(middleware.RealIP)
				router.Use(middleware.Recoverer)*/
				// logger.Info(server)

				go func() {
					if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
						logger.Fatal("failed to start server")
					}
				}()

				// graceful shutdown
				// waitForShutdown(logger, server)

				return nil
			},
			OnStop: func(ctx context.Context) error {
				return logger.Sync()
			},
		},
	)
}

// waitForShutdown graceful shutdown
func waitForShutdown(logger *zap.SugaredLogger, server *http.Server) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Failed gracefully")
		logger.Fatal("failed to gracefully shut down server", err)
	}
}

var Module = fx.Options(
	config.ConfigModule,
	config.LoggerModule,
	fx.Provide(chi.NewRouter),
	server.Module,
	repositories.DatabaseModule,
	services.Module,
	handlers.Module,
	fx.Invoke(bootstrap),
)
