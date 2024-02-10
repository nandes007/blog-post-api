package service

import (
	"context"
	"fmt"
	"nandes007/blog-post-rest-api/helper/jwt"
	"nandes007/blog-post-rest-api/model/web/comment"
	"nandes007/blog-post-rest-api/repository"

	"github.com/go-playground/validator/v10"
)

type CommentServiceImpl struct {
	CommentRepository repository.CommentRepository
	PostRepository    repository.PostRepository
	UserRepository    repository.UserRepository
	Validate          *validator.Validate
}

func NewCommentService(commentRepository repository.CommentRepository, postRepository repository.PostRepository, userRepository repository.UserRepository, validate *validator.Validate) CommentService {
	return &CommentServiceImpl{
		CommentRepository: commentRepository,
		PostRepository:    postRepository,
		UserRepository:    userRepository,
		Validate:          validate,
	}
}

func (r *CommentServiceImpl) Save(ctx context.Context, req *comment.CommentRequest, postId int, token string) (*comment.CommentResponse, error) {
	err := r.Validate.Struct(req)
	if err != nil {
		fmt.Println("Error validate : ", err)
		return nil, err
	}

	tokenFormatted := jwt.FormatToken(token)
	user, err := r.UserRepository.Find(ctx, tokenFormatted)
	if err != nil {
		fmt.Println("Error when get user : ", err)
		return nil, err
	}

	post, err := r.PostRepository.Find(ctx, user, postId)
	if err != nil {
		fmt.Println("Error when get post : ", err)
	}

	postComment, err := r.CommentRepository.Save(ctx, user, post, req)
	if err != nil {
		fmt.Println("Error when create comment : ", err)
	}
	return postComment, nil
}
