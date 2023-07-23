package repositories

import (
	"context"
	"database/sql"
	"hideki/internal/core/domain"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	ctx := context.Background()

	repo := NewAuthRepository(db)

	user := &domain.User{
		ID:        "1",
		Email:     "gini@mail.com",
		UserName:  "",
		Password:  "12356",
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}

	t.Run("OK", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "email", "password", "created_at", "updated_at"}).
			AddRow(user.ID, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)

		mock.ExpectPrepare("SELECT id, email, password, created_at, updated_at FROM users WHERE email = ").
			ExpectQuery().
			WithArgs(user.Email).
			WillReturnRows(rows)

		userDB, err := repo.Login(ctx, &domain.AuthRequest{Email: user.Email, Password: user.Password})
		assert.NoError(t, err)
		assert.Equal(t, user, userDB)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Query Failed", func(t *testing.T) {
		mock.ExpectPrepare("SELECT id, email, password, created_at, updated_at FROM users WHERE email = ").
			ExpectQuery().
			WithArgs(user.Email).
			WillReturnError(sql.ErrConnDone)

		userProfile, err := repo.Login(ctx, &domain.AuthRequest{Email: user.Email, Password: user.Password})
		assert.Error(t, err)
		assert.Empty(t, userProfile)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Prepare Failed", func(t *testing.T) {
		mock.ExpectPrepare("SELECT id, email, password, created_at, updated_at FROM users WHERE email = ").
			WillReturnError(sql.ErrConnDone)

		userMock, err := repo.Login(ctx, &domain.AuthRequest{Email: user.Email, Password: user.Password})
		assert.Error(t, err)
		assert.Empty(t, userMock)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Not Found", func(t *testing.T) {
		mock.ExpectPrepare("SELECT id, email, password, created_at, updated_at FROM users WHERE email = ").
			ExpectQuery().
			WithArgs(user.Email).
			WillReturnError(sql.ErrNoRows)

		userProfile, err := repo.Login(ctx, &domain.AuthRequest{Email: user.Email, Password: user.Password})
		assert.Error(t, err)
		assert.Empty(t, userProfile)

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
