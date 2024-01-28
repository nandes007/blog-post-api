package comment

type Request struct {
	ParentId int    `json:"parent_id"`
	Content  string `validate:"required" json:"content"`
}
