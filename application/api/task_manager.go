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
	GetAllColumnTasks(parentColumnId common.Id) ([]*apimodel.Task, error)
	GeTask(taskId common.Id) (*apimodel.Task, error)
	Create(newTask *entity.Task) (*apimodel.Task, error)
	Update(modifiedTask *entity.Task) (*apimodel.Task, error)
	ChangeDescription(taskId common.Id, newDescription string) (*apimodel.Task, error)
	ChangeName(taskId common.Id, newName string) (*apimodel.Task, error)
	ChangeStatus(taskId common.Id, newParentColumnId common.Id) (*apimodel.Task, error)
	Prioritize(taskId common.Id, priority int) (*apimodel.Task, error)
	DeleteWithAllComments(storedTaskId common.Id) error
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

func (a *TaskManagerAppImpl) GetTaskWithAllCommentsGroupedByCreatedDateTime(parentColumnId common.Id) (*apimodel.Task, error) {
	storedTaskWithRelatedComments, err := repository.GetTaskRepository().GetTaskWithAllCommentsGroupedByCreatedDateTime(parentColumnId)

	if err != nil {
		return nil, err
	}

	return apidto.NewTaskFromAggregate(storedTaskWithRelatedComments), nil
}

func (a *TaskManagerAppImpl) GetAllColumnTasks(parentColumnId common.Id) ([]*apimodel.Task, error) {
	storedTasks, err := repository.GetTaskRepository().GetAllBy(parentColumnId)

	if err != nil {
		return nil, err
	}

	return apidto.NewTasksFromEntities(storedTasks), nil
}

func (a *TaskManagerAppImpl) GeTask(taskId common.Id) (*apimodel.Task, error) {
	storedTask, err := repository.GetTaskRepository().GetBy(taskId)

	if err != nil {
		return nil, err
	}

	return apidto.NewTaskFromEntity(storedTask), nil
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

func (a *TaskManagerAppImpl) ChangeDescription(taskId common.Id, newDescription string) (*apimodel.Task, error) {
	updatedTask, err := repository.GetTaskRepository().UpdateDescription(taskId, newDescription)

	if err != nil {
		return nil, err
	}

	return apidto.NewTaskFromEntity(updatedTask), nil
}

func (a *TaskManagerAppImpl) ChangeName(taskId common.Id, newName string) (*apimodel.Task, error) {
	updatedTask, err := repository.GetTaskRepository().UpdateDescription(taskId, newName)

	if err != nil {
		return nil, err
	}

	return apidto.NewTaskFromEntity(updatedTask), nil
}

func (a *TaskManagerAppImpl) ChangeStatus(taskId common.Id, newParentColumnId common.Id) (*apimodel.Task, error) {
	updatedTask, err := repository.GetTaskRepository().UpdateParentColumn(taskId, newParentColumnId)

	if err != nil {
		return nil, err
	}

	return apidto.NewTaskFromEntity(updatedTask), nil
}

func (a *TaskManagerAppImpl) Prioritize(taskId common.Id, priority int) (*apimodel.Task, error) {
	prioritisedTask, err := repository.GetTaskRepository().UpdatePriority(taskId, priority)

	if err != nil {
		return nil, err
	}

	return apidto.NewTaskFromEntity(prioritisedTask), nil
}

func (a *TaskManagerAppImpl) DeleteWithAllComments(storedTaskId common.Id) error {
	storedComments, err := repository.GetCommentRepository().GetAllBy(storedTaskId)

	if err != nil {
		return err
	}

	storedCommentIds := make([]common.Id, len(storedComments))
	for i, elem := range storedComments {
		storedCommentIds[i] = elem.Id
	}

	if err := repository.GetCommentRepository().DeleteBulk(storedCommentIds); err == nil {
		return err
	}

	if err := repository.GetTaskRepository().Delete(storedTaskId); err == nil {
		return err
	}

	return nil
}
