package repository

import (
	"context"
	"nandes007/blog-post-rest-api/model/web/post"
	"nandes007/blog-post-rest-api/model/web/user"
)

type PostRepository interface {
	Save(ctx context.Context, user *user.UserResponse, post *post.PostRequest) (*post.PostResponse, error)
	GetAll(ctx context.Context, user *user.UserResponse) ([]*post.PostResponse, error)
	Find(ctx context.Context, user *user.UserResponse, id int) (*post.PostResponse, error)
	Update(ctx context.Context, req *post.UpdatePostRequest, user *user.UserResponse) (*post.PostResponse, error)
	Delete(ctx context.Context, id int) error
}
