package service

import (
	"context"
	"nandes007/blog-post-rest-api/model/web/user"
)

type UserService interface {
	Create(ctx context.Context, request user.CreateRequest) user.Response
	FindAll(ctx context.Context) []user.Response
	Login(ctx context.Context, request user.LoginRequest) (string, error)
	Find(ctx context.Context, token string) (user.Response, error)
}
