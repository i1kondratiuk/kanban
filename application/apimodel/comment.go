package apimodel

import (
	"github.com/i1kondratiuk/kanban/domain/entity/common"
	"github.com/i1kondratiuk/kanban/domain/value"
)

// Comment represents a comment
type Comment struct {
	Id              common.Id              `json:"id"`
	CreatedDateTime common.CreatedDateTime `json:"createdDateTime"`
	Comment         value.Comment          `json:"comment"`
}
