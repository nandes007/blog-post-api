package repository

import (
	"context"
	"nandes007/blog-post-rest-api/model/web/comment"
	"nandes007/blog-post-rest-api/model/web/post"
	"nandes007/blog-post-rest-api/model/web/user"
)

type CommentRepository interface {
	Save(ctx context.Context, user *user.UserResponse, post *post.PostResponse, r *comment.CommentRequest) (*comment.CommentResponse, error)
}
