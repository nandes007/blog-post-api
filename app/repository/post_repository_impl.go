package repository

import (
	"database/sql"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/model/web/post"
	"time"
)

type postRepositoryImpl struct {
	db             *sql.DB
	userRepository UserRepository
}

func NewPostRepository(db *sql.DB, userRepository UserRepository) PostRepository {
	return &postRepositoryImpl{
		db:             db,
		userRepository: userRepository,
	}
}

func (r *postRepositoryImpl) Create(req *post.PostRequest, userID int) (*post.PostResponse, error) {
	var id int
	user, err := r.userRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer helper.CommitOrRollback(tx)
	sqlQuery := "INSERT INTO posts(author_id, title, content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err = tx.QueryRow(sqlQuery, userID, req.Title, req.Content, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &post.PostResponse{
		ID:        id,
		UserID:    userID,
		Title:     req.Title,
		Content:   req.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		User:      *user,
	}, nil
}

func (r *postRepositoryImpl) GetAll() ([]*post.PostResponse, error) {
	sqlQuery := "SELECT id, author_id, title, content, created_at, updated_at FROM posts"
	rows, err := r.db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var posts []*post.PostResponse

	for rows.Next() {
		post := &post.PostResponse{}
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}
		user, err := r.userRepository.GetByID(post.ID)
		if err != nil {
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

func (r *postRepositoryImpl) GetByID(id int) (*post.PostResponse, error) {
	sqlQuery := "SELECT id, author_id, title, content, created_at, updated_at FROM posts WHERE id = $1 LIMIT 1"
	row := r.db.QueryRow(sqlQuery, id)
	post := &post.PostResponse{}
	err := row.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return nil, err
	}

	user, err := r.userRepository.GetByID(post.UserID)
	if err != nil {
		return nil, err
	}
	post.User = *user
	return post, nil
}

func (r *postRepositoryImpl) Update(req *post.UpdatePostRequest) (*post.PostResponse, error) {
	sqlQuery := "UPDATE posts SET title = $1, content = $2, updated_at = $3 WHERE id = $4"
	_, err := r.db.Exec(sqlQuery, req.Title, req.Content, time.Now(), req.ID)
	if err != nil {
		return nil, err
	}

	return r.GetByID(req.ID)
}

func (r *postRepositoryImpl) Delete(id int) error {
	sqlQuery := "DELETE FROM posts WHERE id = $1"
	_, err := r.db.Exec(sqlQuery, id)
	if err != nil {
		return err
	}

	return nil
}
