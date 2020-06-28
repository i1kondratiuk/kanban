package api

import (
	"github.com/i1kondratiuk/kanban/application/apidto"
	"github.com/i1kondratiuk/kanban/application/apimodel"
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
	"github.com/i1kondratiuk/kanban/domain/repository"
	"github.com/i1kondratiuk/kanban/domain/service"
)

// ColumnManagerApp represents ColumnManagerApp application to be called by interface layer
type ColumnManagerApp interface {
	GetAllColumnsWithRelatedTasks(boardId *common.Id) ([]*apimodel.Column, error)
	GetColumn(boardId *common.Id) (*apimodel.Column, error)
	Create(newColumn *entity.Column) (*apimodel.Column, error)
	Update(modifiedColumn *entity.Column) (*apimodel.Column, error)
	Rename(columnId common.Id, newName string) (*apimodel.Column, error)
	ChangePosition(columnId common.Id, newPosition int) (*apimodel.Column, error)
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

func (a *ColumnManagerAppImpl) GetAllColumnsWithRelatedTasks(boardId *common.Id) ([]*apimodel.Column, error) {
	storedColumnsWithRelatedTasks, err := repository.GetColumnRepository().GetAllWithRelatedTasksBy(*boardId)

	if err != nil {
		panic(err)
	}

	return apidto.NewColumnsFromAggregates(storedColumnsWithRelatedTasks), nil
}

func (a *ColumnManagerAppImpl) GetColumn(boardId *common.Id) (*apimodel.Column, error) {
	storedColumn, err := repository.GetColumnRepository().GetBy(*boardId)

	if err != nil {
		panic(err)
	}

	return apidto.NewColumnFromEntity(storedColumn), nil
}

func (a *ColumnManagerAppImpl) Create(newColumn *entity.Column) (*apimodel.Column, error) {
	insertedColumn, err := repository.GetColumnRepository().Insert(newColumn)

	if err != nil {
		return nil, err
	}

	return apidto.NewColumnFromEntity(insertedColumn), nil
}

func (a *ColumnManagerAppImpl) Update(modifiedColumn *entity.Column) (*apimodel.Column, error) {
	updatedColumn, err := repository.GetColumnRepository().Update(modifiedColumn)

	if err != nil {
		return nil, err
	}

	return apidto.NewColumnFromEntity(updatedColumn), nil
}

func (a *ColumnManagerAppImpl) Rename(columnId common.Id, newName string) (*apimodel.Column, error) {
	renamedColumn, err := repository.GetColumnRepository().UpdateName(columnId, newName)

	if err != nil {
		return nil, err
	}

	return apidto.NewColumnFromEntity(renamedColumn), nil
}

func (a *ColumnManagerAppImpl) ChangePosition(columnId common.Id, newPosition int) (*apimodel.Column, error) {
	repositionedColumn, err := repository.GetColumnRepository().UpdatePosition(columnId, newPosition)

	if err != nil {
		return nil, err
	}

	return apidto.NewColumnFromEntity(repositionedColumn), nil
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
