package test

import (
	"context"
	"nandes007/blog-post-rest-api/model/web/user"
	"nandes007/blog-post-rest-api/service"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserRepository struct {
	mock.Mock
}

func (m *mockUserRepository) GetAll(ctx context.Context) ([]*user.UserResponse, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*user.UserResponse), args.Error(1)
}

func (m *mockUserRepository) GetByID(id int) (*user.UserResponse, error) {
	args := m.Called(id)
	return args.Get(0).(*user.UserResponse), args.Error(1)
}

func TestUserService_GetUserByID(t *testing.T) {
	mockRepo := &mockUserRepository{}
	validate := validator.New()
	service := service.NewUserService(mockRepo, validate)

	expectedUser := &user.UserResponse{
		ID:    1,
		Name:  "test",
		Email: "test@example.com",
	}

	mockRepo.On("GetByID", 1).Return(expectedUser, nil)
	user, err := service.GetUserByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	mockRepo.AssertExpectations(t)
}
