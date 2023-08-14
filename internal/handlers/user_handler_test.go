package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/unrolled/render"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
	"hideki/internal/core/domain"
	mock "hideki/internal/mocks"
	"hideki/pkg/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandlers_GetProfileHandler(t *testing.T) {
	testCases := map[string]struct {
		ID            any
		buildStubs    func(uc *mock.MockUserService)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		"OK": {
			ID: 1,
			buildStubs: func(uc *mock.MockUserService) {
				uc.EXPECT().
					GetProfile(gomock.Any(), gomock.Any()).
					Times(1).
					Return(&domain.UserProfile{}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusAccepted, recorder.Code)
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			svc := mock.NewMockUserService(ctrl)
			tc.buildStubs(svc)

			recorder := httptest.NewRecorder()

			url := "/v1/me"
			user := &domain.User{
				Email:    "gini@mail.com",
				Password: "123456",
				RoleID:   "1",
			}
			// marshall data to json (like json_encode)
			token, err := utils.GenerateJWT(user)
			if err != nil {
				t.Error(err.Error())
			}
			t.Log(token)
			request, err := http.NewRequest(http.MethodPost, url, nil)
			request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
			assert.NoError(t, err)

			router := chi.NewRouter()
			logger, _ := zap.NewProduction()
			slogger := logger.Sugar()
			r := render.New()

			NewUserHandlers(router, slogger, svc, r)
			router.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}

}
