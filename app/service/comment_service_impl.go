package service

import (
	"context"
	"database/sql"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/helper/jwt"
	"nandes007/blog-post-rest-api/model/domain"
	"nandes007/blog-post-rest-api/model/web/comment"
	"nandes007/blog-post-rest-api/repository"

	"github.com/go-playground/validator/v10"
)

type CommentServiceImpl struct {
	CommentRepository repository.CommentRepository
	PostRepository    repository.PostRepository
	UserRepository    repository.UserRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewCommentService(commentRepository repository.CommentRepository, postRepository repository.PostRepository, userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) CommentService {
	return &CommentServiceImpl{
		CommentRepository: commentRepository,
		PostRepository:    postRepository,
		UserRepository:    userRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (cs CommentServiceImpl) Save(ctx context.Context, r comment.Request, postId int, token string) domain.Comment {
	err := cs.Validate.Struct(r)
	helper.PanicIfError(err)

	tx, err := cs.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	tokenFormatted := jwt.FormatToken(token)
	user, err := cs.UserRepository.Find(ctx, cs.DB, tokenFormatted)
	helper.PanicIfError(err)

	post := cs.PostRepository.Find(ctx, cs.DB, postId)

	comment := cs.CommentRepository.Save(ctx, tx, user, post, r)
	return comment
}
