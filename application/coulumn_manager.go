package application

import (
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
	"github.com/i1kondratiuk/kanban/domain/repository"
	"github.com/i1kondratiuk/kanban/domain/service"
)

// ColumnManagerApp represents ColumnManagerApp application to be called by interface layer
type ColumnManagerApp interface {
	GetAllColumnsWithRelatedTasks(boardId *common.Id) ([]*entity.Column, error)
	Create(newColumn *entity.Column) (*entity.Column, error)
	Rename(columnId common.Id, newName string) (*entity.Column, error)
	ChangePosition(columnId common.Id, newPosition int) (*entity.Column, error)
	Delete(storedColumnId common.Id) error
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
	insertedColumn, err := repository.GetColumnRepository().Insert(newColumn)

	if err != nil {
		return insertedColumn, err
	}

	return insertedColumn, nil
}

func (a *ColumnManagerAppImpl) Rename(columnId common.Id, newName string) (*entity.Column, error) {
	renamedColumn, err := repository.GetColumnRepository().UpdateName(columnId, newName)

	if err != nil {
		return renamedColumn, err
	}

	return renamedColumn, nil
}

func (a *ColumnManagerAppImpl) ChangePosition(columnId common.Id, newPosition int) (*entity.Column, error) {
	repositionedColumn, err := repository.GetColumnRepository().UpdatePosition(columnId, newPosition)

	if err != nil {
		return repositionedColumn, err
	}

	return repositionedColumn, nil
}

func (a *ColumnManagerAppImpl) Delete(storedColumnId common.Id) error {
	parentBoardId, err := repository.GetColumnRepository().GetBoardId(storedColumnId)

	if err != nil {
		return err
	}

	if err := service.GetBoardService().HasColumns(parentBoardId); err != nil {
		return err
	}

	if err := repository.GetColumnRepository().Delete(storedColumnId); err != nil {
		return err
	}

	return nil
}
