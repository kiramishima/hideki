package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"hideki/config"
	"time"
)

var DatabaseModule = fx.Module("db",
	fx.Provide(NewDatabase),
	fx.Provide(func(conn *mongo.Database) *AuthRepository {
		return NewAuthRepository(conn)
	}),
)

// NewDatabase creates an instance of DB
func NewDatabase(lc fx.Lifecycle, cfg *config.Configuration, logger *zap.SugaredLogger) (*mongo.Database, error) {

	// context cancel
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Server API
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	// Client Connection
	opts := options.Client().ApplyURI(cfg.DatabaseURL).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			logger.Panic(err)
			panic(err)
		}
	}()

	// Ping
	status := "up"
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		status = "down"
		return nil, err
	}
	logger.Debugf("Status MongoDB: %s", status)
	// choosing database
	db := client.Database(cfg.DatabaseURL)
	return db, nil
}
