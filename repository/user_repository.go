package repository

import (
	"context"
	"database/sql"
	"nandes007/blog-post-rest-api/model/domain"
)

type UserRepository interface {
	GetAll(ctx context.Context, tx *sql.Tx) []domain.User
	Find(ctx context.Context, tx *sql.DB, token string) (domain.User, error)
}
