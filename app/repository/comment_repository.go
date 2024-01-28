package repository

import (
	"context"
	"database/sql"
	"nandes007/blog-post-rest-api/model/domain"
	"nandes007/blog-post-rest-api/model/web/comment"
)

type CommentRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User, post domain.Post, r comment.Request) domain.Comment
}
