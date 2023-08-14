package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"hideki/config"
	"time"

	// _ "github.com/lib/pq"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var DatabaseModule = fx.Module("db",
	fx.Provide(NewDatabase),
	fx.Provide(func(conn *sqlx.DB) *AuthRepository {
		return NewAuthRepository(conn)
	}),
	fx.Provide(func(conn *sqlx.DB) *UserRepository {
		return NewUserRepository(conn)
	}),
)

// NewDatabase creates an instance of DB
func NewDatabase(lc fx.Lifecycle, cfg *config.Configuration, logger *zap.SugaredLogger) (*sqlx.DB, error) {

	db, err := sqlx.Connect(cfg.DatabaseDriver, cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}
	// seteamos el numero maximo de conexiones abiertas. 0 indica sin limite
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	// seteamos el numero maximo de conexiones inactivas. 0 indica sin limite
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	// usamos time.ParseDuration para convertir el string de duracion a time.Duration
	duration, err := time.ParseDuration(cfg.MaxIdleTime)

	if err != nil {
		return nil, err
	}
	// Seteamos el timeout para las inactivas
	db.SetConnMaxIdleTime(duration)

	// creamos el contexto con 5 segundos de timeout deadline
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// PingContext
	status := "up"
	err = db.PingContext(ctx)
	if err != nil {
		status = "down"
		return nil, err
	}
	logger.Debugf("Status DB: %s", status)
	return db, nil
}
