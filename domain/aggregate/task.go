package aggregate

import (
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
)

// TaskAggregate represents the task entity as the aggregate root with Comments (its related child entities)
type TaskAggregate struct {
	Id          common.Id
	Name        string
	Description string
	Priority    int
	Column      entity.Column
	Comments    []entity.Comment
}
