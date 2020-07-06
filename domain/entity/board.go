package entity

import "github.com/i1kondratiuk/kanban/domain/entity/common"

// Board represents the board entity stored in repository
type Board struct {
	Id          common.Id
	Name        string
	Description string
}

var _ common.Entity = &Board{}

func (b Board) GetId() common.Id {
	return b.Id
}
