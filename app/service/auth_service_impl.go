package service

import (
	"context"
	"fmt"
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

func (s *AuthServiceImpl) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	token, err := s.AuthRepository.Login(ctx, req)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *AuthServiceImpl) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	err := s.Validate.Struct(req)
	if err != nil {
		fmt.Println("Error validate : ", err)
		return nil, err
	}

	createdUser, err := s.AuthRepository.Register(ctx, req)
	if err != nil {
		fmt.Println("Failed registration user : ", err)
		return nil, err
	}

	return createdUser, nil
}
