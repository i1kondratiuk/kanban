package entity

import "github.com/i1kondratiuk/kanban/domain/entity/common"

// Column represents the column entity stored in repository
type Column struct {
	Id       common.Id
	BoardId  common.Id
	Name     string
	Position int
}

var _ common.Entity = &Column{}

func (c Column) GetId() common.Id {
	return c.Id
}
