package api

import (
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
	"github.com/i1kondratiuk/kanban/domain/repository"
)

// TaskManagerApp represents TaskManagerApp application to be called by interface layer
type TaskManagerApp interface {
	GetTaskWithAllCommentsGroupedByCreatedDateTime(taskId common.Id) (*entity.Task, error)
	Create(newTask *entity.Task) (*entity.Task, error)
	Update(columnId common.Id, newName string, newDescription string) (*entity.Task, error)
	DeleteWithAllComments(storedTaskId common.Id) error
	ChangeStatus(taskId common.Id, newParentColumnId common.Id) (*entity.Task, error)
	Prioritize(taskId common.Id, priority int) (*entity.Task, error)
}

// TaskManagerAppImpl is the implementation of UsersCounter
type TaskManagerAppImpl struct{}

var taskManagerApp TaskManagerApp

// InitTaskManagerApp injects implementation for KanbanBoardApp application
func InitTaskManagerApp(a TaskManagerApp) {
	taskManagerApp = a
}

// GetTaskManagerApp returns TaskManagerApp application
func GetTaskManagerApp() TaskManagerApp {
	return taskManagerApp
}

// TaskManagerAppImpl implements the TaskManagerApp interface
var _ TaskManagerApp = &TaskManagerAppImpl{}

func (a *TaskManagerAppImpl) GetTaskWithAllCommentsGroupedByCreatedDateTime(taskId common.Id) (*entity.Task, error) {
	storedTaskWithRelatedComments, err := repository.GetTaskRepository().GetTaskWithAllCommentsGroupedByCreatedDateTime(taskId)

	if err != nil {
		panic(err)
	}

	return storedTaskWithRelatedComments, nil
}

func (a *TaskManagerAppImpl) Create(newTask *entity.Task) (*entity.Task, error) {
	insertedTask, err := repository.GetTaskRepository().Insert(newTask)

	if err != nil {
		return insertedTask, err
	}

	return insertedTask, nil
}

func (a *TaskManagerAppImpl) Update(columnId common.Id, newName string, newDescription string) (*entity.Task, error) {
	updatedColumn, err := repository.GetTaskRepository().Update(columnId, newName, newDescription)

	if err != nil {
		return updatedColumn, err
	}

	return updatedColumn, nil
}

func (a *TaskManagerAppImpl) DeleteWithAllComments(storedTaskId common.Id) error {
	_, err := repository.GetCommentRepository().GetAllBy(storedTaskId)
	// err = repository.GetCommentRepository().DeleteBulk(storedComments)
	err = repository.GetTaskRepository().Delete(storedTaskId)

	if err != nil {
		return err
	}

	return nil
}

func (a *TaskManagerAppImpl) ChangeStatus(taskId common.Id, newParentColumnId common.Id) (*entity.Task, error) {
	updatedTask, err := repository.GetTaskRepository().UpdateParentColumn(taskId, newParentColumnId)

	if err != nil {
		return updatedTask, err
	}

	return updatedTask, nil
}

func (a *TaskManagerAppImpl) Prioritize(taskId common.Id, priority int) (*entity.Task, error) {
	prioritisedTask, err := repository.GetTaskRepository().UpdatePriority(taskId, priority)

	if err != nil {
		return prioritisedTask, err
	}

	return prioritisedTask, nil
}
