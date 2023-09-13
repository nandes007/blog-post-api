package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"log"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/helper/response"
	"nandes007/blog-post-rest-api/model/domain"
	"nandes007/blog-post-rest-api/model/web/user"
	"nandes007/blog-post-rest-api/repository"
	"strings"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []user.Response {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.GetAll(ctx, tx)
	return response.ToUserResponses(users)
}

func (service *UserServiceImpl) Find(ctx context.Context, token string) (user.Response, error) {
	var user domain.User
	tokenFormatted := strings.Replace(token, "Bearer ", "", 1)

	user, err := service.UserRepository.Find(ctx, service.DB, tokenFormatted)

	if err != nil {
		log.Fatal(err)
		return response.ToUserResponse(user), err
	}

	return response.ToUserResponse(user), nil
}
