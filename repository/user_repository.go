package repository

import (
	"context"
	"database/sql"
	"nandes007/blog-post-rest-api/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	GetAll(ctx context.Context, tx *sql.Tx) []domain.User
}
