package post

import (
	"nandes007/blog-post-rest-api/model/web/user"
	"time"
)

type PostRequest struct {
	Title   string `validate:"required" json:"title"`
	Content string `validate:"required" json:"content"`
}

type UpdatePostRequest struct {
	ID      int    `json:"id"`
	Title   string `validate:"required" json:"title"`
	Content string `validate:"required" json:"content"`
}

type PostResponse struct {
	Id        int               `json:"id"`
	Title     string            `json:"title"`
	Content   string            `json:"content"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	User      user.UserResponse `json:"author"`
}
