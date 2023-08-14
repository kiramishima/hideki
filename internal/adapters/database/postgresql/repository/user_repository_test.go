package repository

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"hideki/internal/core/domain"
	"testing"
	"time"
)

func TestUserRepository_GetProfile(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	ctx := context.Background()
	repo := NewUserRepository(sqlxDB)

	user := &domain.UserProfile{
		ID:        1,
		UserName:  "giny",
		FullName:  "Martha",
		Bio:       "",
		Role:      "user",
		Picture:   "https://i.pravatar.cc/300",
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}

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

	t.Run("OK", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "username", "full_name", "bio", "picture", "role", "created_at", "updated_at"}).
			AddRow(user.ID, user.UserName, user.FullName, user.Bio, user.Picture, user.Role, user.CreatedAt, user.UpdatedAt)

		mock.ExpectPrepare(query).
			ExpectQuery().
			WithArgs(user.ID).
			WillReturnRows(rows)

		userProfile, err := repo.GetProfile(ctx, user.ID)
		// t.Log("up", userProfile)
		// t.Log("err", err)
		assert.NoError(t, err)
		assert.Equal(t, user, userProfile)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Query Failed", func(t *testing.T) {
		mock.ExpectPrepare(query).
			ExpectQuery().
			WithArgs(user.ID).
			WillReturnError(sql.ErrConnDone)

		userProfile, err := repo.GetProfile(ctx, user.ID)
		assert.Error(t, err)
		assert.Empty(t, userProfile)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Prepare Failed", func(t *testing.T) {
		mock.ExpectPrepare(query).
			WillReturnError(sql.ErrConnDone)

		userProfile, err := repo.GetProfile(ctx, user.ID)
		assert.Error(t, err)
		assert.Empty(t, userProfile)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Not Found", func(t *testing.T) {
		mock.ExpectPrepare(query).
			ExpectQuery().
			WithArgs(user.ID).
			WillReturnError(sql.ErrNoRows)

		userProfile, err := repo.GetProfile(ctx, user.ID)
		assert.Error(t, err)
		assert.Empty(t, userProfile)

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
