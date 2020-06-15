package entity

import "github.com/i1kondratiuk/kanban/domain/entity/common"

// Board represents the board entity stored in repository
type Board struct {
	Id          common.Id `json:"id"`
	Name        string    `json:"name"`
	Columns     []Column  `json:"columns"`
}
