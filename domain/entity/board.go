package entity

// Board represents the board entity stored in repository
type Board struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Columns []Column `json:"columns"`
}
