package service

import (
	"errors"

	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
)

const defaultColumnName = "Default"

// ColumnService represents the service to handle business rules related to Kanban Boards' Columns
type ColumnService interface {
	CreateDefaultColumn(parentBoardId common.Id) *entity.Column
	IsLeftMostColumn(position int) bool
	GetRightMostColumnPosition(boardId common.Id) (int, error)
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

// ColumnServiceImpl implements the ColumnService interface
var _ ColumnService = &ColumnServiceImpl{}

// Creates default Column with parent Board assigned
func (s *ColumnServiceImpl) CreateDefaultColumn(parentBoardId common.Id) *entity.Column {
	return &entity.Column{
		BoardId: parentBoardId,
		Name:    defaultColumnName,
	}
}

func (s *ColumnServiceImpl) IsLeftMostColumn(position int) bool {
	return position == 1
}

func (s *ColumnServiceImpl) GetRightMostColumnPosition(boardId common.Id) (int, error) {
	return 0, errors.New("GetAllSortedByNameAsc: implement me")
}
