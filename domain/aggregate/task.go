package aggregate

import (
	"github.com/i1kondratiuk/kanban/domain/entity"
)

// TaskAggregate represents the task entity as the aggregate root with Comments (its related child entities)
type TaskAggregate struct {
	TaskAggregateRoot *entity.Task
	Comments          []*entity.Comment
}
