package services

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
	"hideki/internal/core/domain"
	mock "hideki/internal/mocks"
	"testing"
	"time"
)

func TestLogin(t *testing.T) {
	logger, _ := zap.NewProduction()
	slogger := logger.Sugar()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repo := mock.NewMockAuthRepository(mockCtrl)
	repo.EXPECT().Login(gomock.Any(), gomock.Any()).Return(&domain.User{
		ID:        "1",
		Email:     "gini@mail.com",
		UserName:  "",
		Password:  "12356",
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}, nil)

	uc := NewAuthService(slogger, repo)

	t.Run("OK", func(t *testing.T) {
		ctx := context.Background()
		data := &domain.AuthRequest{Email: "gini@mail.com", Password: "123456"}
		b, err := uc.Login(ctx, data)
		t.Log("B", b)
		assert.NoError(t, err)
		assert.Equal(t, "gini@mail.com", b.Token)
	})
	t.Run("Not Found", func(t *testing.T) {
		ctx := context.Background()
		data := &domain.AuthRequest{Email: "gini@mail.com", Password: ""}
		b, err := uc.Login(ctx, data)
		t.Log(b)
		t.Log(err)
		assert.Error(t, err)
	})
}
