package repository

import (
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
)

// ColumnRepository represents a storage of all existing columns
type ColumnRepository interface {
	GetAllWithRelatedTasksBy(parentBoardId common.Id) ([]*entity.Column, error)
}

var columnRepository ColumnRepository

// GetColumnRepository returns the ColumnRepository
func GetColumnRepository() ColumnRepository {
	return columnRepository
}

// InitColumnRepository injects ColumnRepository with its implementation
func InitColumnRepository(r ColumnRepository) {
	columnRepository = r
}
