package bootstrap

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/unrolled/render"
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

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func bootstrap(
	lifecycle fx.Lifecycle,
	logger *zap.SugaredLogger,
	server *http.Server,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				logger.Info("Starting API")

				go func() {
					if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
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
	fx.Provide(func() *chi.Mux {
		var r = chi.NewRouter()
		r.Use(middleware.Timeout(60 * time.Second))
		r.Use(middleware.RequestID)
		r.Use(middleware.RealIP)
		r.Use(middleware.Recoverer)
		r.Use(middleware.Logger)
		r.Use(middleware.Compress(5))
		return r
	}),
	fx.Provide(func() *render.Render {
		return render.New()
	}),
	server.Module,
	repositories.DatabaseModule,
	services.Module,
	handlers.Module,
	fx.Invoke(bootstrap),
)
