package repository

import (
	"context"
	"nandes007/blog-post-rest-api/model/web/user"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]*user.UserResponse, error)
	Find(ctx context.Context, token string) (*user.UserResponse, error)
}
