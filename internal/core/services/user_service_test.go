package services

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
	"hideki/internal/core/domain"
	mock "hideki/internal/mocks"
	httpErrors "hideki/pkg/errors"
	"testing"
	"time"
)

func TestGetProfile(t *testing.T) {
	logger, _ := zap.NewProduction()
	slogger := logger.Sugar()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	repo := mock.NewMockUserRepository(mockCtrl)
	repo.EXPECT().GetProfile(gomock.Any(), gomock.Eq(int16(1))).Times(1).Return(&domain.UserProfile{
		ID:        1,
		UserName:  "giny",
		FullName:  "Martha",
		Bio:       "",
		Role:      "user",
		Picture:   "https://i.pravatar.cc/300",
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}, nil)
	repo.EXPECT().GetProfile(gomock.Any(), gomock.Any()).Times(1).Return(nil, httpErrors.ErrUserNotFound)

	uc := NewUserService(slogger, repo)

	t.Run("OK", func(t *testing.T) {
		ctx := context.Background()
		b, err := uc.GetProfile(ctx, int16(1))
		t.Log("err", err)
		t.Log("B", b)
		assert.NoError(t, err)
		assert.Equal(t, "giny", b.UserName)
	})
	t.Run("Not Found", func(t *testing.T) {
		ctx := context.Background()
		b, err := uc.GetProfile(ctx, int16(2))
		t.Log("B", err)
		t.Log(b)
		t.Log(err)
		assert.Error(t, err)
	})
}
