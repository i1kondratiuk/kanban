package repository

import (
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/value"
)

// TaskRepository represents a storage of all existing tasks
type TaskRepository interface {
	GetAllBy(parentColumnId value.Id) ([]*entity.Task, error)
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
