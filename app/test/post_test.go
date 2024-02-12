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

func (m *mockPostRepository) Create(ctx context.Context, user *user.UserResponse, req *post.PostRequest) (*post.PostResponse, error) {
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
