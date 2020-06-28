package aggregate

import (
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
)

// TaskAggregate represents the task entity as the aggregate root with Comments (its related child entities)
type TaskAggregate struct {
	Id          common.Id        `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Priority    int              `json:"priority"`
	Position    int              `json:"position"`
	Column      entity.Column    `json:"column"`
	Comments    []entity.Comment `json:"comments"`
}
