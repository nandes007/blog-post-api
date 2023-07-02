package repository

import (
	"context"
	"database/sql"
	"nandes007/blog-post-rest-api/model/domain"
)

type PostRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User, post domain.Post) domain.Post
	GetAll(ctx context.Context, tx *sql.Tx) []domain.Post
	Find(ctx context.Context, tx *sql.DB, id int) domain.Post
	Update(ctx context.Context, tx *sql.DB, post domain.Post, id int) bool
	Delete(ctx context.Context, tx *sql.DB, id int) bool
}
