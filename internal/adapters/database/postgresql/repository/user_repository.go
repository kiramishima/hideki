package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"hideki/internal/core/domain"
	dbErrors "hideki/pkg/errors"
	"log"
)

// UserRepository struct
type UserRepository struct {
	db *sqlx.DB
}

// NewUserRepository Creates a new instance of UserRepository
func NewUserRepository(conn *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: conn,
	}
}

func (repo *UserRepository) GetProfile(ctx context.Context, uid int16) (*domain.UserProfile, error) {
	var query = `SELECT u.id,
       u.username,
       CONCAT(up.name, ' ', up.first_lastname, ' ', up.second_lastname) full_name,
       COALESCE(up.bio, '') bio,
       up.picture,
       (SELECT r.name FROM model_has_roles mhr
       INNER JOIN roles r ON mhr.role_id = r.id
       WHERE mhr.model_id = up.user_id) role,
       up.created_at,
       up.updated_at 
	FROM user_profile up
	INNER JOIN users u ON up.user_id = u.id
  	WHERE up.user_id = $1`

	stmt, err := repo.db.PreparexContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", dbErrors.ErrPrepareStatement, err)
	}
	defer stmt.Close()

	q := domain.UserProfileQueryItem{}

	row := stmt.QueryRowxContext(ctx, uid)
	// var updatedAt sql.NullTime
	err = row.StructScan(&q) //(&u.ID, &u.UserName, &u.FullName, &u.Bio, &u.Picture, &u.Role, &u.CreatedAt, &updatedAt)
	log.Println(err)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, dbErrors.ErrUserNotFound
		} else {
			return nil, fmt.Errorf("%s: %w", dbErrors.ErrScanData, err)
		}
	}

	//
	u := &domain.UserProfile{
		ID:        q.ID.Int16,
		UserName:  q.UserName.String,
		FullName:  q.FullName.String,
		Role:      q.Role.String,
		Bio:       q.Bio.String,
		Picture:   q.Picture.String,
		CreatedAt: q.CreatedAt.Time,
		UpdatedAt: q.UpdatedAt.Time,
		DeletedAt: q.DeletedAt.Time,
	}

	return u, nil
}
