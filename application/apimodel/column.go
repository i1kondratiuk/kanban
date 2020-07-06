package apimodel

import "github.com/i1kondratiuk/kanban/domain/entity/common"

// Column represents the column entity stored in repository
type Column struct {
	Id       common.Id `json:"id"`
	BoardId  common.Id `json:"boardId"`
	Name     string    `json:"name"`
	Position int       `json:"position"`
	Tasks    []*Task   `json:"tasks,omitempty"`
}
