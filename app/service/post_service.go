package service

import (
	"nandes007/blog-post-rest-api/model/web/post"
)

type PostService interface {
	CreatePost(req *post.PostRequest, userID int) (*post.PostResponse, error)
	GetAllPosts() ([]*post.PostResponse, error)
	GetPostByID(id int) (*post.PostResponse, error)
	UpdatePost(req *post.UpdatePostRequest) (*post.PostResponse, error)
	DeletePost(id int) error
}
