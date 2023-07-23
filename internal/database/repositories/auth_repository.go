package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"hideki/internal/core/domain"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(conn *sql.DB) *AuthRepository {
	return &AuthRepository{
		db: conn,
	}
}

// Login Repository method for sign in
func (repo *AuthRepository) Login(ctx context.Context, data *domain.AuthRequest) (*domain.User, error) {
	stmt, err := repo.db.PrepareContext(ctx, "SELECT id, email, password, created_at, updated_at FROM users WHERE email = $1")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrPrepareStatement, err)
	}
	defer stmt.Close()

	u := &domain.User{}

	row := stmt.QueryRowContext(ctx, data.Email)
	var updatedAt sql.NullTime
	err = row.Scan(&u.ID, &u.Email, &u.Password, &u.CreatedAt, &updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		} else {
			return nil, fmt.Errorf("%s: %w", ErrScanData, err)
		}
	}

	if updatedAt.Valid {
		u.UpdatedAt = updatedAt.Time
	}

	return u, nil
}
