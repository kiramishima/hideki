// Code generated by MockGen. DO NOT EDIT.
// Source: .\internal\core\ports\repository\auth_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	domain "hideki/internal/core/domain"
	httpErrors "hideki/pkg/errors"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockAuthRepository is a mock of AuthRepository interface.
type MockAuthRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAuthRepositoryMockRecorder
}

// MockAuthRepositoryMockRecorder is the mock recorder for MockAuthRepository.
type MockAuthRepositoryMockRecorder struct {
	mock *MockAuthRepository
}

// NewMockAuthRepository creates a new mock instance.
func NewMockAuthRepository(ctrl *gomock.Controller) *MockAuthRepository {
	mock := &MockAuthRepository{ctrl: ctrl}
	mock.recorder = &MockAuthRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthRepository) EXPECT() *MockAuthRepositoryMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockAuthRepository) Login(ctx context.Context, data *domain.AuthRequest) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, data)
	ret0, _ := ret[0].(*domain.User)
	// ret1, _ := ret[1].(error)
	// return ret0, ret1
	if ret0.Email == data.Email {
		return ret0, nil
	} else {
		return nil, httpErrors.ErrUserNotFound
	}
}

// Login indicates an expected call of Login.
func (mr *MockAuthRepositoryMockRecorder) Login(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthRepository)(nil).Login), ctx, data)
}
