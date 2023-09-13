package service

import (
	"context"
	"nandes007/blog-post-rest-api/model/web/auth"
)

type AuthService interface {
	Login(ctx context.Context, request auth.LoginRequest) (string, error)
	Register(ctx context.Context, request auth.RegisterRequest) (auth.RegisterResponse, error)
}
