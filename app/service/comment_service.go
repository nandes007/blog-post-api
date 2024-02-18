package service

import (
	"nandes007/blog-post-rest-api/model/web/comment"
)

type CommentService interface {
	CreateComment(req *comment.CommentRequest, postID int, userID int) (*comment.CommentResponse, error)
}
