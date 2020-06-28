package entity

import "github.com/i1kondratiuk/kanban/domain/entity/common"

// Task represents the task entity stored in repository
type Task struct {
	Id          common.Id
	Name        string
	Description string
	Priority    int
	Position    int
	Column      Column
}
