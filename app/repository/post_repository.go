package repository

import (
	"context"
	"database/sql"
	"nandes007/blog-post-rest-api/model/domain"
	"nandes007/blog-post-rest-api/model/web/post"
	"nandes007/blog-post-rest-api/model/web/user"
)

type PostRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User, post domain.Post) domain.Post
	GetAll(ctx context.Context, db *sql.DB, user user.Response) []domain.Post
	Find(ctx context.Context, db *sql.DB, id int) domain.Post
	Update(ctx context.Context, tx *sql.Tx, post post.CreateRequest, id int) bool
	Delete(ctx context.Context, tx *sql.Tx, id int) bool
}
