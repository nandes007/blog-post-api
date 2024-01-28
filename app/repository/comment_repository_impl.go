package repository

import (
	"context"
	"database/sql"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/model/domain"
	"nandes007/blog-post-rest-api/model/web/comment"
)

type CommentRepositoryImpl struct {
}

func NewCommentRepository() CommentRepository {
	return &CommentRepositoryImpl{}
}

// Save implements CommentRepository.
func (c CommentRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User, post domain.Post, r comment.Request) domain.Comment {
	comment := domain.Comment{}
	var id int
	sqlQuery := "INSERT INTO post_comments(post_id, user_id, parent_id, content) VALUES ($1, $2, $3, $4) RETURNING id"
	err := tx.QueryRowContext(ctx, sqlQuery, post.Id, user.Id, r.ParentId, r.Content).Scan(&id)
	helper.PanicIfError(err)

	comment.Id = id
	comment.User = user
	comment.Post = post
	comment.ParentId = r.ParentId
	comment.Content = r.Content
	return comment
}
