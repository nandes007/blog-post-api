package service

import (
	"context"
	"nandes007/blog-post-rest-api/model/web/user"
)

type UserService interface {
	FindAll(ctx context.Context) []user.Response
	Find(ctx context.Context, token string) (user.Response, error)
}
