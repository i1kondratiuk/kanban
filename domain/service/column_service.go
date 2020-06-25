package service

import (
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
)

const defaultColumnName = "Default"

// ColumnService represents the service to handle business rules related to Kanban Boards' Columns
type ColumnService interface {
	CreateDefaultColumn(parentBoardId common.Id) *entity.Column
}

// ColumnServiceImpl is the implementation of ColumnService
type ColumnServiceImpl struct{}

var columnService ColumnService

// GetColumnService returns a ColumnService
func GetColumnService() ColumnService {
	return columnService
}

// InitColumnService injects ColumnService with its implementation
func InitColumnService(a ColumnService) {
	columnService = a
}

// Creates default Column with parent Board assigned
func (a *ColumnServiceImpl) CreateDefaultColumn(parentBoardId common.Id) *entity.Column {
	return &entity.Column{
		Board: entity.Board{Id: parentBoardId},
		Name:  defaultColumnName,
	}
}
