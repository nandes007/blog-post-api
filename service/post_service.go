package service

import (
	"context"
	"nandes007/blog-post-rest-api/model/web/post"
)

type PostService interface {
	Create(ctx context.Context, request post.CreateRequest, token string) post.Response
}
