package service

import (
	"context"
	"nandes007/blog-post-rest-api/model/web/comment"
)

type CommentService interface {
	Save(ctx context.Context, req *comment.CommentRequest, postId int, token string) (*comment.CommentResponse, error)
}
