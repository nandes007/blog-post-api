package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/helper/response"
	"nandes007/blog-post-rest-api/model/domain"
	"nandes007/blog-post-rest-api/model/web/auth"
	"nandes007/blog-post-rest-api/repository"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewAuthService(authRepository repository.AuthRepository, DB *sql.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: authRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *AuthServiceImpl) Login(ctx context.Context, request auth.LoginRequest) (string, error) {
	generateToken, err := service.AuthRepository.Login(ctx, service.DB, request)

	if err != nil {
		return "", err
	}

	return response.ToUserLoginResponse(generateToken), nil
}

func (service *AuthServiceImpl) Register(ctx context.Context, request auth.RegisterRequest) (auth.RegisterResponse, error) {
	err := service.Validate.Struct(request)
	response := auth.RegisterResponse{}
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	user, err = service.AuthRepository.Register(ctx, tx, user)
	helper.PanicIfError(err)

	response.Name = user.Name
	response.Email = user.Email

	return response, nil
}
