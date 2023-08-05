package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hideki/internal/core/domain"
	dbErrors "hideki/pkg/errors"
	"time"
)

// AuthRepository struct
type AuthRepository struct {
	db *mongo.Database
}

// NewAuthRepository Creates a new instance of AuthRepository
func NewAuthRepository(conn *mongo.Database) *AuthRepository {
	return &AuthRepository{
		db: conn,
	}
}

// Login Repository method for sign in
func (repo *AuthRepository) Login(ctx context.Context, data *domain.AuthRequest) (*domain.User, error) {
	u := &domain.User{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := repo.db.Collection("users").FindOne(ctx, bson.M{"email": data.Email}).Decode(&u)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", dbErrors.ErrPrepareStatement, err)
	}

	var pipeline []bson.M
	pipeline = append(pipeline, bson.M{"$match": bson.M{"email": data.Email}})
	pipeline = append(pipeline, bson.M{"$lookup": bson.M{
		"from": "user_profile",
		"let": bson.M{
			"user": "$_id",
		},
		"pipeline": bson.A{
			bson.M{
				"$match": bson.M{
					"$expr": bson.M{
						"$and": bson.A{
							bson.M{
								"$eq": bson.A{
									"$$user",
									"$user_id",
								},
							},
						},
					},
				},
			},
		},
		"as": "profile",
	}})
	pipeline = append(pipeline, bson.M{"$unwind": bson.M{"path": "$profile", "preserveNullAndEmptyArrays": true}})
	pipeline = append(pipeline, bson.M{"$project": bson.M{
		"_id":        "$_id",
		"email":      "$email",
		"first_name": "$first_name",
		"last_name":  "$last_name",
		"gender":     "$gender",
		"about":      "$about",
		"role_id":    "$role_id",
	},
	})

	opts := options.Aggregate()
	cur, err := repo.db.Collection("users").Aggregate(ctx, pipeline, opts)
	defer cur.Close(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", dbErrors.ErrPrepareStatement, err)
	}

	for cur.Next(ctx) {
		_ = cur.Decode(&u)
	}

	return u, nil
}
