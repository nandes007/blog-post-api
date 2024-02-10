package service

import (
	"context"
	"nandes007/blog-post-rest-api/model/web/post"
)

type PostService interface {
	Create(ctx context.Context, req *post.PostRequest, token string) (*post.PostResponse, error)
	FindAll(ctx context.Context, token string) ([]*post.PostResponse, error)
	Find(ctx context.Context, token string, id int) (*post.PostResponse, error)
	Update(ctx context.Context, req *post.UpdatePostRequest, token string) (*post.PostResponse, error)
	Delete(ctx context.Context, id int) error
}
