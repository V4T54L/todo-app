package models

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Done      bool   `json:"done"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
