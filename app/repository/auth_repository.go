package repository

import (
	"context"
	"database/sql"
	"nandes007/blog-post-rest-api/model/domain"
	"nandes007/blog-post-rest-api/model/web/auth"
)

type AuthRepository interface {
	Login(ctx context.Context, db *sql.DB, request auth.LoginRequest) (string, error)
	Register(ctx context.Context, db *sql.Tx, user domain.User) (domain.User, error)
}
