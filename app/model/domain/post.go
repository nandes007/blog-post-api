package domain

import "nandes007/blog-post-rest-api/model/web/user"

type Post struct {
	Id        int
	AuthorId  int
	Title     string
	Content   string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	User      user.Response
}
