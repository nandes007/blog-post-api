package post

type CreateRequest struct {
	Title   string `validate:"required" json:"title"`
	Content string `validate:"required" json:"content"`
}
