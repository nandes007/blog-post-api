package service

import (
	"nandes007/blog-post-rest-api/model/web/user"
)

type UserService interface {
	GetAllUsers() ([]*user.UserResponse, error)
	GetUserByID(id int) (*user.UserResponse, error)
}
