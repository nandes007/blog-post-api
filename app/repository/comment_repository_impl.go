package repository

import (
	"context"
	"database/sql"
	"nandes007/blog-post-rest-api/model/web/comment"
	"nandes007/blog-post-rest-api/model/web/post"
	"nandes007/blog-post-rest-api/model/web/user"
	"time"
)

type commentRepositoryImpl struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{
		db: db,
	}
}

// Save implements CommentRepository.
func (r commentRepositoryImpl) Save(ctx context.Context, user *user.UserResponse, post *post.PostResponse, req *comment.CommentRequest) (*comment.CommentResponse, error) {
	var id int
	sqlQuery := "INSERT INTO post_comments(post_id, user_id, parent_id, content, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id"
	err := r.db.QueryRowContext(ctx, sqlQuery, post.Id, user.Id, req.ParentId, req.Content, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &comment.CommentResponse{
		Id:        id,
		PostId:    post.Id,
		UserId:    user.Id,
		ParentId:  req.ParentId,
		Content:   req.Content,
		User:      *user,
		Post:      *post,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
