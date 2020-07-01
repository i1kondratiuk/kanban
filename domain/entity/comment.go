package entity

import (
	"github.com/i1kondratiuk/kanban/domain/entity/common"
	"github.com/i1kondratiuk/kanban/domain/value"
)

// Comment represents a comment
type Comment struct {
	Id              common.Id
	CreatedDateTime common.CreatedDateTime
	Parent          common.Entity
	Comment         value.Comment
}

var _ common.Entity = &Comment{}

func (c Comment) GetId() common.Id {
	return c.Id
}
