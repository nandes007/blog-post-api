package repository

import (
	"nandes007/blog-post-rest-api/model/web/post"
)

type PostRepository interface {
	Create(post *post.PostRequest, userID int) (*post.PostResponse, error)
	GetAll() ([]*post.PostResponse, error)
	GetByID(id int) (*post.PostResponse, error)
	Update(req *post.UpdatePostRequest) (*post.PostResponse, error)
	Delete(id int) error
}
