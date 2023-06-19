package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/helper/response"
	"nandes007/blog-post-rest-api/model/domain"
	"nandes007/blog-post-rest-api/model/web/user"
	"nandes007/blog-post-rest-api/repository"
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

func (service *UserServiceImpl) Create(ctx context.Context, request user.CreateRequest) user.Response {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	user = service.UserRepository.Save(ctx, tx, user)

	return response.ToUserResponse(user)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []user.Response {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.GetAll(ctx, tx)
	return response.ToUserResponses(users)
}

func (service *UserServiceImpl) Login(ctx context.Context, request user.LoginRequest) string {

	generateToken := service.UserRepository.Login(ctx, service.DB, request)
	return response.ToUserLoginResponse(generateToken)
}
