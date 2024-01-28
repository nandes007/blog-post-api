package service

import (
	"context"
	"nandes007/blog-post-rest-api/model/domain"
	"nandes007/blog-post-rest-api/model/web/comment"
)

type CommentService interface {
	Save(ctx context.Context, r comment.Request, postId int, token string) domain.Comment
}
