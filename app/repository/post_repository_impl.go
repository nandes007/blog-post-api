package repository

import (
	"context"
	"database/sql"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/model/web/post"
	"nandes007/blog-post-rest-api/model/web/user"
	"time"
)

type postRepositoryImpl struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepositoryImpl{
		db: db,
	}
}

func (r postRepositoryImpl) Save(ctx context.Context, user *user.UserResponse, req *post.PostRequest) (*post.PostResponse, error) {
	var id int
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer helper.CommitOrRollback(tx)
	sqlQuery := "INSERT INTO posts(author_id, title, content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err = tx.QueryRowContext(ctx, sqlQuery, user.Id, req.Title, req.Content, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &post.PostResponse{
		Id:        id,
		Title:     req.Title,
		Content:   req.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		User:      *user,
	}, nil
}

func (r postRepositoryImpl) GetAll(ctx context.Context, user *user.UserResponse) ([]*post.PostResponse, error) {
	sqlQuery := "SELECT id, title, content, created_at, updated_at FROM posts"
	rows, err := r.db.QueryContext(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var posts []*post.PostResponse

	for rows.Next() {
		post := &post.PostResponse{}
		if err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}
		post.User = *user
		posts = append(posts, post)
	}

	err = rows.Close()
	if err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r postRepositoryImpl) Find(ctx context.Context, user *user.UserResponse, id int) (*post.PostResponse, error) {
	sqlQuery := "SELECT id, title, content, created_at, updated_at FROM posts WHERE id = $1 LIMIT 1"
	row := r.db.QueryRowContext(ctx, sqlQuery, id)
	post := &post.PostResponse{}
	err := row.Scan(&post.Id, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return nil, err
	}
	post.User = *user
	return post, nil
}

func (r postRepositoryImpl) Update(ctx context.Context, req *post.UpdatePostRequest, user *user.UserResponse) (*post.PostResponse, error) {
	sqlQuery := "UPDATE posts SET title = $1, content = $2, updated_at = $3 WHERE id = $4"
	_, err := r.db.ExecContext(ctx, sqlQuery, req.Title, req.Content, time.Now(), req.ID)
	if err != nil {
		return nil, err
	}

	return r.Find(context.Background(), user, req.ID)
}

func (r postRepositoryImpl) Delete(ctx context.Context, id int) error {
	sqlQuery := "DELETE FROM posts WHERE id = $1"
	_, err := r.db.ExecContext(ctx, sqlQuery, id)
	if err != nil {
		return err
	}

	return nil
}
