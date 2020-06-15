package application

import (
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
)

// TaskManagerApp represents TaskManagerApp application to be called by interface layer
type TaskManagerApp interface {
	Get(taskId common.Id) (*entity.Task, error)
	Create(newTask *entity.Task) (*entity.Task, error)
	Update(columnId common.Id, newName string, newDescription string) (*entity.Task, error)
	DeleteWithAllComments(storedTaskId common.Id) error
	ChangeStatus(taskId common.Id, newParentColumnId common.Id) (*entity.Task, error)
	Prioritize(taskId common.Id, priority int) error
	GetTaskWithAllCommentsGroupedByCreatedDateTime(taskId *common.Id) ([]*entity.Comment, error)
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

func (a *TaskManagerAppImpl) GetAllTaskCommentsGroupedByCreatedDateTime(taskId *common.Id) ([]*entity.Comment, error) {
	panic("implement me")
}

func (a *TaskManagerAppImpl) Get(taskId common.Id) (*entity.Task, error) {
	panic("implement me")
}

func (a *TaskManagerAppImpl) Create(newTask *entity.Task) (*entity.Task, error) {
	panic("implement me")
}

func (a *TaskManagerAppImpl) Update(columnId common.Id, newName string, newDescription string) (*entity.Task, error) {
	panic("implement me")
}

func (a *TaskManagerAppImpl) DeleteWithAllComments(storedTaskId common.Id) error {
	panic("implement me")
}

func (a *TaskManagerAppImpl) ChangeStatus(taskId common.Id, newParentColumnId common.Id) (*entity.Task, error) {
	panic("implement me")
}

func (a *TaskManagerAppImpl) Prioritize(taskId common.Id, priority int) error {
	panic("implement me")
}

func (a *TaskManagerAppImpl) GetTaskWithAllCommentsGroupedByCreatedDateTime(taskId *common.Id) ([]*entity.Comment, error) {
	panic("implement me")
}
