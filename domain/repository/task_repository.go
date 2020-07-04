package repository

import (
	"github.com/i1kondratiuk/kanban/domain/aggregate"
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
)

// TaskRepository represents a storage of all existing tasks
type TaskRepository interface {
	GetTaskWithAllCommentsGroupedByCreatedDateTime(taskId common.Id) (*aggregate.TaskAggregate, error)
	GetAllBy(parentColumnId common.Id) ([]*entity.Task, error)
	Insert(newTask *entity.Task) (*entity.Task, error)
	Update(modifiedTask *entity.Task) (*entity.Task, error)
	UpdateName(storedTaskId common.Id, newName string) (*entity.Task, error)
	UpdateDescription(storedTaskId common.Id, newDescription string) (*entity.Task, error)
	UpdateParentColumn(storedTaskId common.Id, newParentColumnId common.Id) (*entity.Task, error)
	UpdatePriority(storedTaskId common.Id, priority int) (*entity.Task, error)
	Delete(storedTaskId common.Id) error
}

var taskRepository TaskRepository

// GetTaskRepository returns the TaskRepository
func GetTaskRepository() TaskRepository {
	return taskRepository
}

// InitTaskRepository injects TaskRepository with its implementation
func InitTaskRepository(r TaskRepository) {
	taskRepository = r
}
