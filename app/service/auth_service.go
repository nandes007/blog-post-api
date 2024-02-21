package service

import (
	"nandes007/blog-post-rest-api/model/web/auth"
)

type AuthService interface {
	Login(request *auth.LoginRequest) (*auth.LoginResponse, error)
	Register(request *auth.RegisterRequest) (*auth.RegisterResponse, error)
}
