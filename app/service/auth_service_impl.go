package service

import (
	"fmt"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/model/web/auth"
	"nandes007/blog-post-rest-api/repository"

	"github.com/go-playground/validator/v10"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	Validate       *validator.Validate
}

func NewAuthService(authRepository repository.AuthRepository, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: authRepository,
		Validate:       validate,
	}
}

func (s *AuthServiceImpl) Login(req *auth.LoginRequest) (*auth.LoginResponse, error) {
	err := s.Validate.Struct(req)
	helper.PanicIfError(err)
	token, err := s.AuthRepository.Login(req)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *AuthServiceImpl) Register(req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	err := s.Validate.Struct(req)
	helper.PanicIfError(err)
	createdUser, err := s.AuthRepository.Register(req)
	if err != nil {
		fmt.Println("Failed registration user : ", err)
		return nil, err
	}

	return createdUser, nil
}
