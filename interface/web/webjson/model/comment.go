package entity

import (
	"github.com/i1kondratiuk/kanban/domain/value"
)

// Comment represents the model exposed to the API client
type Comment struct {
	value.Comment
}
