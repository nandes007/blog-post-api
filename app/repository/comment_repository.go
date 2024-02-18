package repository

import (
	"nandes007/blog-post-rest-api/model/web/comment"
)

type CommentRepository interface {
	Create(r *comment.CommentRequest, postID int, userID int) (*comment.CommentResponse, error)
}
