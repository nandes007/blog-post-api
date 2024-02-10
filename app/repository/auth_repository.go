package repository

import (
	"context"
	"nandes007/blog-post-rest-api/model/web/auth"
)

type AuthRepository interface {
	Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error)
	Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error)
}
