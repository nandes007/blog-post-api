package service

import (
	"fmt"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/model/web/comment"
	"nandes007/blog-post-rest-api/repository"

	"github.com/go-playground/validator/v10"
)

type commentServiceImpl struct {
	CommentRepository repository.CommentRepository
	Validate          *validator.Validate
}

func NewCommentService(commentRepository repository.CommentRepository, validate *validator.Validate) CommentService {
	return &commentServiceImpl{
		CommentRepository: commentRepository,
		Validate:          validate,
	}
}

func (r *commentServiceImpl) CreateComment(req *comment.CommentRequest, postID int, userID int) (*comment.CommentResponse, error) {
	err := r.Validate.Struct(req)
	helper.PanicIfError(err)
	postComment, err := r.CommentRepository.Create(req, postID, userID)
	if err != nil {
		fmt.Println("Error when create comment : ", err)
	}
	return postComment, nil
}
