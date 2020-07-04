package entity

import "github.com/i1kondratiuk/kanban/domain/entity/common"

// Task represents the task entity stored in repository
type Task struct {
	Id          common.Id
	ColumnId    common.Id
	Name        string
	Description string
	Priority    int
}

var _ common.Entity = &Task{}

func (t Task) GetId() common.Id {
	return t.Id
}
