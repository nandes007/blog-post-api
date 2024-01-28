package domain

import "database/sql"

type Comment struct {
	Id       int
	PostId   int
	UserId   int
	ParentId sql.NullInt64
	Content  string
	User     User
	Post     Post
}
