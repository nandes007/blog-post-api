package comment

import (
	"database/sql"
	"nandes007/blog-post-rest-api/model/web/post"
	"nandes007/blog-post-rest-api/model/web/user"
	"time"
)

type CommentRequest struct {
	ParentId sql.NullInt64 `json:"parent_id"`
	Content  string        `validate:"required" json:"content"`
}

type CommentResponse struct {
	Id        int               `json:"id"`
	PostId    int               `json:"post_id"`
	UserId    int               `json:"user_id"`
	ParentId  sql.NullInt64     `json:"parent_id"`
	Content   string            `json:"content"`
	User      user.UserResponse `json:"user"`
	Post      post.PostResponse `json:"post"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}
