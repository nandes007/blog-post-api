package service

import (
	"context"
	"nandes007/blog-post-rest-api/model/web/post"
)

type PostService interface {
	Create(ctx context.Context, request post.CreateRequest, token string) post.Response
	FindAll(ctx context.Context) []post.Response
	Find(ctx context.Context, id int) post.Response
	Update(ctx context.Context, request post.CreateRequest, id int) post.Response
	Delete(ctx context.Context, id int) bool
}
