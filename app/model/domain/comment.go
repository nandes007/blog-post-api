package domain

type Comment struct {
	Id       int
	PostId   int
	UserId   int
	ParentId int
	Content  string
}
