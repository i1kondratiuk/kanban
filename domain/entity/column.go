package entity

// Column represents the column entity stored in repository
type Column struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Board Board  `json:"board"`
	Tasks []Task `json:"tasks"`
}
