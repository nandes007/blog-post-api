package test

import (
	"errors"
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

func TestPostService_GetPostByID(t *testing.T) {
	mockRepo := &mockPostRepository{}
	validate := validator.New()
	service := service.NewPostService(mockRepo, validate)

	expected := &post.PostResponse{
		ID:      1,
		UserID:  1,
		Title:   "My First Post",
		Content: "Hello World!",
	}

	mockRepo.On("GetByID", 1).Return(expected, nil)
	post, err := service.GetPostByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expected, post)
	mockRepo.AssertExpectations(t)
}

func TestPostService_UpdatePost(t *testing.T) {
	mockRepo := &mockPostRepository{}
	validate := validator.New()
	service := service.NewPostService(mockRepo, validate)

	req := &post.UpdatePostRequest{
		ID:      1,
		Title:   "Updated Title",
		Content: "Updated Content",
	}

	expected := &post.PostResponse{
		ID:      1,
		UserID:  1,
		Title:   "Updated Title",
		Content: "Updated Content",
	}

	mockRepo.On("Update", req).Return(expected, nil)
	post, err := service.UpdatePost(req)
	assert.NoError(t, err)
	assert.Equal(t, expected, post)
	mockRepo.AssertExpectations(t)
}

func TestPostService_DeletePost(t *testing.T) {
	mockRepo := &mockPostRepository{}
	validate := validator.New()
	service := service.NewPostService(mockRepo, validate)

	mockRepo.On("Delete", 1).Return(nil)

	err := service.DeletePost(1)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPostService_CreatePost_Error(t *testing.T) {
	mockRepo := &mockPostRepository{}
	validate := validator.New()
	service := service.NewPostService(mockRepo, validate)

	req := &post.PostRequest{
		Title:   "My First Post",
		Content: "Hello World",
	}

	mockRepo.On("Create", req, 1).Return(&post.PostResponse{}, errors.New("failed to create post"))
	post, err := service.CreatePost(req, 1)
	assert.Error(t, err)
	assert.Nil(t, post)
	mockRepo.AssertExpectations(t)
}

func TestPostService_GetAllPosts_Error(t *testing.T) {
	mockRepo := &mockPostRepository{}
	validate := validator.New()
	service := service.NewPostService(mockRepo, validate)

	expected := []*post.PostResponse{
		{ID: 1, UserID: 1, Title: "User1 Post", Content: "Hello World!"},
		{ID: 2, UserID: 2, Title: "user2 Post", Content: "Hello Worlds"},
	}

	mockRepo.On("GetAll").Return(expected, errors.New("failed to retrieve posts"))
	posts, err := service.GetAllPosts()
	assert.Error(t, err)
	assert.Nil(t, posts)
	mockRepo.AssertExpectations(t)
}

func TestPostService_GetPostByID_Error(t *testing.T) {
	mockRepo := &mockPostRepository{}
	validate := validator.New()
	service := service.NewPostService(mockRepo, validate)

	mockRepo.On("GetByID", 1).Return(&post.PostResponse{}, errors.New("failed to retrieve post"))
	post, err := service.GetPostByID(1)
	assert.Error(t, err)
	assert.Nil(t, post)
	mockRepo.AssertExpectations(t)
}

func TestPostService_UpdatePost_Error(t *testing.T) {
	mockRepo := &mockPostRepository{}
	validate := validator.New()
	service := service.NewPostService(mockRepo, validate)

	req := &post.UpdatePostRequest{
		ID:      1,
		Title:   "Updated Title",
		Content: "Updated Content",
	}

	mockRepo.On("Update", req).Return(&post.PostResponse{}, errors.New("failed to update post"))
	post, err := service.UpdatePost(req)
	assert.Error(t, err)
	assert.Nil(t, post)
	mockRepo.AssertExpectations(t)
}

func TestPostService_DeletePost_Error(t *testing.T) {
	mockRepo := &mockPostRepository{}
	validate := validator.New()
	service := service.NewPostService(mockRepo, validate)

	mockRepo.On("Delete", 1).Return(errors.New("failed to delete post"))
	err := service.DeletePost(1)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
