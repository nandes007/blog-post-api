package repository

import (
	"nandes007/blog-post-rest-api/model/web/user"
)

type UserRepository interface {
	GetAll() ([]*user.UserResponse, error)
	GetByID(id int) (*user.UserResponse, error)
}
