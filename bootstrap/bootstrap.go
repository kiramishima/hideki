package bootstrap

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/unrolled/render"
	"hideki/config"
	"hideki/internal/adapters/cache/redis"
	"hideki/internal/adapters/database/postgresql/repository"
	"hideki/internal/core/services"
	"hideki/internal/handlers"
	"hideki/internal/server"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func bootstrap(
	lifecycle fx.Lifecycle,
	logger *zap.SugaredLogger,
	server *server.Server,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				logger.Info("Starting API")

				/*go func() {
					if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
						logger.Fatal("failed to start server")
					}
				}()*/
				_ = server.Run()

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
	repository.DatabaseModule,
	services.Module,
	handlers.Module,
	redis.Module,
	fx.Invoke(bootstrap),
)
