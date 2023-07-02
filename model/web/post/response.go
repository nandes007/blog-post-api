package post

type Response struct {
	Id        int    `json:"id"`
	AuthorId  int    `json:"author_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
