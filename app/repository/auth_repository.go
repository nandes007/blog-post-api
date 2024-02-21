package repository

import (
	"nandes007/blog-post-rest-api/model/web/auth"
)

type AuthRepository interface {
	Login(req *auth.LoginRequest) (*auth.LoginResponse, error)
	Register(req *auth.RegisterRequest) (*auth.RegisterResponse, error)
}
