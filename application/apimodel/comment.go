package apimodel

import (
	"time"

	"github.com/i1kondratiuk/kanban/domain/entity/common"
	"github.com/i1kondratiuk/kanban/domain/value"
)

// Comment represents a comment
type Comment struct {
	Id              common.Id     `json:"id"`
	ParentId        common.Id     `json:"parentId"`
	CreatedDateTime time.Time     `json:"createdDateTime"`
	Comment         value.Comment `json:"comment"`
}
