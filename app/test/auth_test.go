package test

import (
	"nandes007/blog-post-rest-api/model/web/auth"

	"github.com/stretchr/testify/mock"
)

type mockAuthRepository struct {
	mock.Mock
}

func (m *mockAuthRepository) Login(req *auth.LoginRequest) (*auth.LoginResponse, error) {
	args := m.Called(req)
	return args.Get(0).(*auth.LoginResponse), args.Error(1)
}

func (m *mockAuthRepository) Register(req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	args := m.Called(req)
	return args.Get(0).(*auth.RegisterResponse), args.Error(1)
}
