package model

type (
	Todo struct {
		ID          int    `db:"id" json:"id"`
		Title    string `db:"title" json:"title"`
    Description     string `db:"description" json:"description"`
    Status     int `db:"status" json:"status"`
	}
)
