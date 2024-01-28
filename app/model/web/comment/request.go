package comment

import "database/sql"

type Request struct {
	ParentId sql.NullInt64 `json:"parent_id"`
	Content  string        `validate:"required" json:"content"`
}
