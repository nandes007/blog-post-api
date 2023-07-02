package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"log"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/helper/jwt"
	"nandes007/blog-post-rest-api/helper/response"
	"nandes007/blog-post-rest-api/model/domain"
	"nandes007/blog-post-rest-api/model/web/post"
	"nandes007/blog-post-rest-api/repository"
)

type PostServiceImpl struct {
	PostRepository repository.PostRepository
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewPostService(postRepository repository.PostRepository, userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) PostService {
	return &PostServiceImpl{
		PostRepository: postRepository,
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service PostServiceImpl) Create(ctx context.Context, request post.CreateRequest, token string) post.Response {
	//TODO implement me
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	//helper.PanicIfError(err)
	if err != nil {
		log.Fatal(err)
	}
	defer helper.CommitOrRollback(tx)

	tokenFormatted := jwt.FormatToken(token)
	user, err := service.UserRepository.Find(ctx, service.DB, tokenFormatted)

	//helper.PanicIfError(err)
	if err != nil {
		log.Fatal(err)
	}

	post := domain.Post{
		AuthorId:  user.Id,
		Title:     request.Title,
		Content:   request.Content,
		CreatedAt: helper.GetCurrentTime(),
		UpdatedAt: helper.GetCurrentTime(),
	}

	post = service.PostRepository.Save(ctx, tx, user, post)

	return response.ToPostResponse(post)
}
