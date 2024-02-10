package service

import (
	"context"
	"fmt"
	"nandes007/blog-post-rest-api/model/web/user"
	"nandes007/blog-post-rest-api/repository"
	"strings"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (s *UserServiceImpl) FindAll(ctx context.Context) ([]*user.UserResponse, error) {
	users, err := s.UserRepository.GetAll(ctx)
	if err != nil {
		fmt.Println("Error when get all users : ", err)
	}
	return users, nil
}

func (s *UserServiceImpl) Find(ctx context.Context, token string) (*user.UserResponse, error) {
	tokenFormatted := strings.Replace(token, "Bearer ", "", 1)
	user, err := s.UserRepository.Find(ctx, tokenFormatted)

	if err != nil {
		fmt.Println("Error when get user : ", err)
		return nil, err
	}

	return user, nil
}
