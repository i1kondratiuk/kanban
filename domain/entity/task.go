package entity

import "github.com/i1kondratiuk/kanban/domain/entity/common"

// Task represents the task entity stored in repository
type Task struct {
	Id          common.Id `json:"id"`
	Name        string    `json:"name"`
	Priority    int       `json:"priority"`
	Position    int       `json:"position"`
	Column      Column    `json:"column"`
}
