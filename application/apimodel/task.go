package apimodel

import (
	"github.com/i1kondratiuk/kanban/domain/entity/common"
)

// Task represents the task entity stored in repository
type Task struct {
	Id          common.Id  `json:"id"`
	ColumnId    common.Id  `json:"columnId"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Priority    int        `json:"priority"`
	Comments    []*Comment `json:"comments,omitempty"`
}
