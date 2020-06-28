package entity

import "github.com/i1kondratiuk/kanban/domain/entity/common"

// Column represents the column entity stored in repository
type Column struct {
	Id       common.Id
	Name     string
	Position int
	Board    Board
}
