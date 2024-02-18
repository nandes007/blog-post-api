package service

import (
	"context"
	"nandes007/blog-post-rest-api/model/web/user"
)

type UserService interface {
	GetAllUsers(ctx context.Context) ([]*user.UserResponse, error)
	GetUserByID(id int) (*user.UserResponse, error)
}
