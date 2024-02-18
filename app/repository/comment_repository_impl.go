package repository

import (
	"database/sql"
	"nandes007/blog-post-rest-api/model/web/comment"
	"time"
)

type commentRepositoryImpl struct {
	db             *sql.DB
	postRepository PostRepository
	userRepository UserRepository
}

func NewCommentRepository(db *sql.DB, postRepository PostRepository, userRepository UserRepository) CommentRepository {
	return &commentRepositoryImpl{
		db:             db,
		postRepository: postRepository,
		userRepository: userRepository,
	}
}

// Save implements CommentRepository.
func (r *commentRepositoryImpl) Create(req *comment.CommentRequest, postID int, userID int) (*comment.CommentResponse, error) {
	var id int
	sqlQuery := "INSERT INTO post_comments(post_id, user_id, parent_id, content, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id"
	err := r.db.QueryRow(sqlQuery, postID, userID, req.ParentId, req.Content, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return nil, err
	}

	post, err := r.postRepository.GetByID(postID)
	if err != nil {
		return nil, err
	}

	user, err := r.userRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}

	return &comment.CommentResponse{
		ID:        id,
		PostId:    postID,
		UserId:    userID,
		ParentId:  req.ParentId,
		Content:   req.Content,
		User:      *user,
		Post:      *post,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
