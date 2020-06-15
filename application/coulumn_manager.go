package application

import (
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
	"github.com/i1kondratiuk/kanban/domain/repository"
)

// ColumnManagerApp represents ColumnManagerApp application to be called by interface layer
type ColumnManagerApp interface {
	GetAllColumnsWithRelatedTasks(boardId *common.Id) ([]*entity.Column, error)
	Create(newColumn *entity.Column) (*entity.Column, error)
	Rename(columnId common.Id, newName string) error
	Delete(storedColumn *entity.Column) error
	ChangePosition(columnId common.Id, newPosition int) error
}

// ColumnManagerAppImpl is the implementation of ColumnManagerApp
type ColumnManagerAppImpl struct{}

var columnManagerApp ColumnManagerApp

// InitBoardManagerApp injects implementation for KanbanBoardApp application
func InitColumnManagerApp(a ColumnManagerApp) {
	columnManagerApp = a
}

// GetColumnManagerApp returns ColumnManagerApp application
func GetColumnManagerApp() ColumnManagerApp {
	return columnManagerApp
}

// ColumnManagerAppImpl implements the ColumnManagerApp interface
var _ ColumnManagerApp = &ColumnManagerAppImpl{}

func (a *ColumnManagerAppImpl) GetAllColumnsWithRelatedTasks(boardId *common.Id) ([]*entity.Column, error) {
	storedColumnsWithRelatedTasks, err := repository.GetColumnRepository().GetAllWithRelatedTasksBy(*boardId)

	if err != nil {
		panic(err)
	}

	return storedColumnsWithRelatedTasks, nil
}

func (a *ColumnManagerAppImpl) Create(newColumn *entity.Column) (*entity.Column, error) {
	panic("implement me")
}

func (a *ColumnManagerAppImpl) Rename(columnId common.Id, newName string) error {
	panic("implement me")
}

func (a *ColumnManagerAppImpl) Delete(storedColumn *entity.Column) error {
	panic("implement me")
}

func (a *ColumnManagerAppImpl) ChangePosition(columnId common.Id, newPosition int) error {
	panic("implement me")
}
