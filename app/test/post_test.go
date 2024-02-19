package test

import (
	"nandes007/blog-post-rest-api/model/web/post"
	"nandes007/blog-post-rest-api/model/web/user"
	"nandes007/blog-post-rest-api/service"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockPostRepository struct {
	mock.Mock
}

func (m *mockPostRepository) Create(req *post.PostRequest, userID int) (*post.PostResponse, error) {
	args := m.Called(req, userID)
	return args.Get(0).(*post.PostResponse), args.Error(1)
}

func (m *mockPostRepository) GetAll() ([]*post.PostResponse, error) {
	args := m.Called()
	return args.Get(0).([]*post.PostResponse), args.Error(1)
}

func (m *mockPostRepository) GetByID(id int) (*post.PostResponse, error) {
	args := m.Called(id)
	return args.Get(0).(*post.PostResponse), args.Error(1)
}

func (m *mockPostRepository) Update(req *post.UpdatePostRequest) (*post.PostResponse, error) {
	args := m.Called(req)
	return args.Get(0).(*post.PostResponse), args.Error(1)
}

func (m *mockPostRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestPostService_CreatePost(t *testing.T) {
	mockRepo := &mockPostRepository{}
	validate := validator.New()
	service := service.NewPostService(mockRepo, validate)

	user := &user.UserResponse{
		ID:    1,
		Name:  "test",
		Email: "test@example.com",
	}

	req := &post.PostRequest{
		Title:   "My First Post",
		Content: "Hello World!",
	}

	expected := &post.PostResponse{
		ID:      1,
		UserID:  user.ID,
		Title:   "My First Post",
		Content: "Hello World!",
		User:    *user,
	}

	mockRepo.On("Create", req, user.ID).Return(expected, nil)
	post, err := service.CreatePost(req, user.ID)
	assert.NoError(t, err)
	assert.Equal(t, expected, post)
	mockRepo.AssertExpectations(t)
}

func TestPostService_GetAllPosts(t *testing.T) {
	mockRepo := &mockPostRepository{}
	validate := validator.New()
	service := service.NewPostService(mockRepo, validate)

	user1 := &user.UserResponse{
		ID:    1,
		Name:  "test1",
		Email: "test1@example.com",
	}

	user2 := &user.UserResponse{
		ID:    2,
		Name:  "test2",
		Email: "test2@example.com",
	}

	expected := []*post.PostResponse{
		{ID: 1, UserID: user1.ID, Title: "User1 Post", Content: "Hello World!", User: *user1},
		{ID: 2, UserID: user2.ID, Title: "user2 Post", Content: "Hello Worlds", User: *user2},
	}

	mockRepo.On("GetAll").Return(expected, nil)
	posts, err := service.GetAllPosts()
	assert.NoError(t, err)
	assert.Equal(t, expected, posts)
	mockRepo.AssertExpectations(t)
}
