package entity

import (
	"github.com/i1kondratiuk/kanban/domain/entity/common"
	"github.com/i1kondratiuk/kanban/domain/value"
)

// Comment represents a comment
type Comment struct {
	Id              common.Id
	CreatedDateTime common.CreatedDateTime
	Comment         value.Comment
}
