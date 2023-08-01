package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"hideki/internal/core/domain"
	dbErrors "hideki/pkg/errors"
)

// UserRepository struct
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository Creates a new instance of UserRepository
func NewUserRepository(conn *sql.DB) *UserRepository {
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

	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", dbErrors.ErrPrepareStatement, err)
	}
	defer stmt.Close()

	u := &domain.UserProfile{}

	row := stmt.QueryRowContext(ctx, uid)
	var updatedAt sql.NullTime
	err = row.Scan(&u.ID, &u.UserName, &u.FullName, &u.Bio, &u.Picture, &u.Role, &u.CreatedAt, &updatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, dbErrors.ErrUserNotFound
		} else {
			return nil, fmt.Errorf("%s: %w", dbErrors.ErrScanData, err)
		}
	}

	if updatedAt.Valid {
		u.UpdatedAt = updatedAt.Time
	}

	return u, nil
}
