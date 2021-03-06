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
	GetColumn(columnId *common.Id) (*apimodel.Column, error)
	Create(newColumn *entity.Column) (*apimodel.Column, error)      // TODO bulk create
	Update(modifiedColumn *entity.Column) (*apimodel.Column, error) // TODO bulk update
	Rename(columnId common.Id, newName string) error                // TODO use the app update logic instead
	ChangePosition(columnId common.Id, newPosition int) error       // TODO use the app update logic instead
	Delete(storedColumnId common.Id) error                          // TODO bulk delete
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
		return nil, err
	}

	return apidto.NewColumnsFromAggregates(storedColumnsWithRelatedTasks), nil
}

func (a *ColumnManagerAppImpl) GetColumn(columnId *common.Id) (*apimodel.Column, error) {
	storedColumn, err := repository.GetColumnRepository().GetBy(*columnId)

	if err != nil {
		return nil, err
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

func (a *ColumnManagerAppImpl) Rename(columnId common.Id, newName string) (err error) {
	err = repository.GetColumnRepository().UpdateName(columnId, newName)
	return
}

func (a *ColumnManagerAppImpl) ChangePosition(columnId common.Id, newPosition int) (err error) {
	err = repository.GetColumnRepository().UpdatePosition(columnId, newPosition)
	return
}

func (a *ColumnManagerAppImpl) Delete(storedColumnId common.Id) error {
	parentBoardId, err := repository.GetColumnRepository().GetBoardId(storedColumnId)

	if err != nil {
		return err
	}

	isColumnDeletable, err := service.GetBoardService().IsColumnDeletable(parentBoardId, 1)

	if err != nil {
		return err
	}

	if !isColumnDeletable {
		return nil
	}

	if err := repository.GetColumnRepository().Delete(storedColumnId); err != nil {
		return err
	}

	return nil
}
