package model

import "github.com/i1kondratiuk/kanban/domain/entity/common"

// Column represents the column entity stored in repository
type Column struct {
	Id       common.Id `json:"id"`
	Name     string    `json:"name"`
	Position int       `json:"position"`
	Tasks    []*Task   `json:"tasks"`
}
