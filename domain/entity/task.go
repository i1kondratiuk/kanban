package entity

// Task represents the task entity stored in repository
type Task struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Priority int    `json:"priority"`
	Position int    `json:"position"`
	Column   Column `json:"column"`
}
