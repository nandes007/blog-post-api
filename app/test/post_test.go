package test

import (
	"context"
	"nandes007/blog-post-rest-api/model/web/post"
	"nandes007/blog-post-rest-api/model/web/user"

	"github.com/stretchr/testify/mock"
)

type mockPostRepository struct {
	mock.Mock
}

func (m *mockPostRepository) Save(ctx context.Context, user *user.UserResponse, req *post.PostRequest) (*post.PostResponse, error) {
	args := m.Called(ctx, user, req)
	return args.Get(0).(*post.PostResponse), args.Error(1)
}

func (m *mockPostRepository) GetAll(ctx context.Context, user *user.UserResponse) ([]*post.PostResponse, error) {
	args := m.Called(ctx, user)
	return args.Get(0).([]*post.PostResponse), args.Error(1)
}

func (m *mockPostRepository) Find(ctx context.Context, user *user.UserResponse, id int) (*post.PostResponse, error) {
	args := m.Called(ctx, user, id)
	return args.Get(0).(*post.PostResponse), args.Error(1)
}

func (m *mockPostRepository) Update(ctx context.Context, req *post.UpdatePostRequest, user *user.UserResponse) (*post.PostResponse, error) {
	args := m.Called(ctx, req, user)
	return args.Get(0).(*post.PostResponse), args.Error(1)
}

func (m *mockPostRepository) Delete(ctx context.Context, id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// func (m *mockUserRepository) GetAll(ctx context.Context) ([]*user.UserResponse, error) {
// 	args := m.Called(ctx)
// 	return args.Get(0).([]*user.UserResponse), args.Error(1)
// }

// func (m *mockUserRepository) Find(ctx context.Context, token string) (*user.UserResponse, error) {
// 	args := m.Called(ctx, token)
// 	return args.Get(0).(*user.UserResponse), args.Error(1)
// }

// func TestPostService_CreatePost(t *testing.T) {
// 	postMockRepo := &mockPostRepository{}
// 	userMockRepo := &mockUserRepository{}
// 	validate := validator.New()

// 	token := "TOKEN"
// 	req := &post.PostRequest{
// 		Title:   "Daily Blog",
// 		Content: "Hello world!",
// 	}
// 	user := &user.UserResponse{
// 		Id:        1,
// 		Name:      "test",
// 		Email:     "test@example.com",
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}

// 	expected := &post.PostResponse{
// 		Id:      1,
// 		Title:   "test",
// 		Content: "Hello world!",
// 		User:    *user,
// 	}

// 	service := service.NewPostService(postMockRepo, userMockRepo, validate)
// 	postMockRepo.On("Save", context.Background(), user, req).Return(expected, nil)
// 	createdPost, err := service.Create(context.Background(), req, token)
// 	assert.NoError(t, err)
// 	assert.Equal(t, expected, createdPost)
// 	postMockRepo.AssertExpectations(t)
// }
