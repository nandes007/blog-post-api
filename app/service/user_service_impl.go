package service

import (
	"fmt"
	"nandes007/blog-post-rest-api/model/web/user"
	"nandes007/blog-post-rest-api/repository"

	"github.com/go-playground/validator/v10"
)

type userServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &userServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (s *userServiceImpl) GetAllUsers() ([]*user.UserResponse, error) {
	users, err := s.UserRepository.GetAll()
	if err != nil {
		fmt.Println("Error when get all users : ", err)
	}
	return users, nil
}

func (s *userServiceImpl) GetUserByID(id int) (*user.UserResponse, error) {
	user, err := s.UserRepository.GetByID(id)

	if err != nil {
		fmt.Println("Error when get user : ", err)
		return nil, err
	}

	return user, nil
}
