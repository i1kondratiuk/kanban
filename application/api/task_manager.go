package api

import (
	"github.com/i1kondratiuk/kanban/application/apidto"
	"github.com/i1kondratiuk/kanban/application/apimodel"
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
	"github.com/i1kondratiuk/kanban/domain/repository"
)

// TaskManagerApp represents TaskManagerApp application to be called by interface layer
type TaskManagerApp interface {
	GetTaskWithAllCommentsGroupedByCreatedDateTime(taskId common.Id) (*apimodel.Task, error)
	GetTasksBy(parentColumnId common.Id) ([]*apimodel.Task, error)
	Create(newTask *entity.Task) (*apimodel.Task, error)              // TODO bulk create
	Update(modifiedTask *entity.Task) (*apimodel.Task, error)         // TODO bulk update
	ChangeDescription(taskId common.Id, newDescription string) error  // TODO use the app update logic instead
	ChangeName(taskId common.Id, newName string) error                // TODO use the app update logic instead
	ChangeStatus(taskId common.Id, newParentColumnId common.Id) error // TODO use the app update logic instead
	Prioritize(taskId common.Id, priority int) error                  // TODO use the app update logic instead
	DeleteWithAllComments(storedTaskId common.Id) error               // TODO bulk delete
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

func (a *TaskManagerAppImpl) GetTaskWithAllCommentsGroupedByCreatedDateTime(taskId common.Id) (*apimodel.Task, error) {
	storedTaskWithRelatedComments, err := repository.GetTaskRepository().GetTaskWithAllCommentsGroupedByCreatedDateTime(taskId)

	if err != nil {
		return nil, err
	}

	return apidto.NewTaskFromAggregate(storedTaskWithRelatedComments), nil
}

func (a *TaskManagerAppImpl) GetTasksBy(parentColumnId common.Id) ([]*apimodel.Task, error) {
	storedTasks, err := repository.GetTaskRepository().GetAllBy(parentColumnId)

	if err != nil {
		return nil, err
	}

	return apidto.NewTasksFromEntities(storedTasks), nil
}

func (a *TaskManagerAppImpl) Create(newTask *entity.Task) (*apimodel.Task, error) {
	insertedTask, err := repository.GetTaskRepository().Insert(newTask)

	if err != nil {
		return nil, err
	}

	return apidto.NewTaskFromEntity(insertedTask), nil
}

func (a *TaskManagerAppImpl) Update(modifiedTask *entity.Task) (*apimodel.Task, error) {
	updatedTask, err := repository.GetTaskRepository().Update(modifiedTask)

	if err != nil {
		return nil, err
	}

	return apidto.NewTaskFromEntity(updatedTask), nil
}

func (a *TaskManagerAppImpl) ChangeDescription(taskId common.Id, newDescription string) (err error) {
	err = repository.GetTaskRepository().UpdateDescription(taskId, newDescription)
	return
}

func (a *TaskManagerAppImpl) ChangeName(taskId common.Id, newName string) (err error) {
	err = repository.GetTaskRepository().UpdateDescription(taskId, newName)
	return
}

func (a *TaskManagerAppImpl) ChangeStatus(taskId common.Id, newParentColumnId common.Id) (err error) {
	err = repository.GetTaskRepository().UpdateParentColumn(taskId, newParentColumnId)
	return
}

func (a *TaskManagerAppImpl) Prioritize(taskId common.Id, priority int) (err error) {
	err = repository.GetTaskRepository().UpdatePriority(taskId, priority)
	return
}

func (a *TaskManagerAppImpl) DeleteWithAllComments(storedTaskId common.Id) (err error) {
	err = repository.GetCommentRepository().DeleteAllBy(storedTaskId)
	err = repository.GetTaskRepository().Delete(storedTaskId)
	return
}
