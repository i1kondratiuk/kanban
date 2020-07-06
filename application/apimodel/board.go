package apimodel

import (
	"github.com/i1kondratiuk/kanban/domain/entity/common"
)

// Board represents the board entity stored in repository
type Board struct {
	Id          common.Id `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Columns     []*Column `json:"columns,omitempty"`
}
