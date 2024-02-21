package test

import (
	"errors"
	"nandes007/blog-post-rest-api/model/web/auth"
	"nandes007/blog-post-rest-api/service"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
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

func TestAuthService_Login(t *testing.T) {
	mockRepo := &mockAuthRepository{}
	validate := validator.New()
	service := service.NewAuthService(mockRepo, validate)

	expected := &auth.LoginResponse{
		Token: "asdadasdasdasdadsasdasd",
	}

	req := &auth.LoginRequest{
		Email:    "test@example.com",
		Password: "password",
	}

	mockRepo.On("Login", req).Return(expected, nil)
	token, err := service.Login(req)
	assert.NoError(t, err)
	assert.Equal(t, expected, token)
	mockRepo.AssertExpectations(t)
}

func TestAuthService_Register(t *testing.T) {
	mockRepo := &mockAuthRepository{}
	validate := validator.New()
	service := service.NewAuthService(mockRepo, validate)

	expected := &auth.RegisterResponse{
		ID:    1,
		Name:  "test",
		Email: "test@example.com",
	}

	req := &auth.RegisterRequest{
		Name:     "test",
		Email:    "test@example.com",
		Password: "password",
	}

	mockRepo.On("Register", req).Return(expected, nil)

	res, err := service.Register(req)
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
	mockRepo.AssertExpectations(t)
}

func TestAuthService_Login_Error(t *testing.T) {
	mockRepo := &mockAuthRepository{}
	validate := validator.New()
	service := service.NewAuthService(mockRepo, validate)

	req := &auth.LoginRequest{
		Email:    "test@example.com",
		Password: "password",
	}

	mockRepo.On("Login", req).Return(&auth.LoginResponse{}, errors.New("failed to login"))
	res, err := service.Login(req)
	assert.Error(t, err)
	assert.Nil(t, res)
	mockRepo.AssertExpectations(t)
}

func TestAuthService_Register_Error(t *testing.T) {
	mockRepo := &mockAuthRepository{}
	validate := validator.New()
	service := service.NewAuthService(mockRepo, validate)

	req := &auth.RegisterRequest{
		Name:     "test",
		Email:    "test@example.com",
		Password: "password",
	}

	mockRepo.On("Register", req).Return(&auth.RegisterResponse{}, errors.New("failed to register"))
	res, err := service.Register(req)
	assert.Error(t, err)
	assert.Nil(t, res)
	mockRepo.AssertExpectations(t)
}
