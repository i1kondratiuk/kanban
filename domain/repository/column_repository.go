package repository

import (
	"github.com/i1kondratiuk/kanban/domain/aggregate"
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
)

// ColumnRepository represents a storage of all existing columns
type ColumnRepository interface {
	GetAllWithRelatedTasksBy(parentBoardId common.Id) ([]*aggregate.ColumnAggregate, error)
	GetBy(parentBoardId common.Id) (*entity.Column, error)
	GetByChildTaskId(taskId common.Id) (entity.Column, error)
	GetByParentBoardIdAndPosition(parentBoardId common.Id, position int) (*entity.Column, error)
	GetBoardId(columnId common.Id) (parentBoardId common.Id, err error)
	CountAllBy(parentBoardId common.Id) (int, error)
	Insert(newColumn *entity.Column) (*entity.Column, error)
	Update(updatedColumn *entity.Column) (*entity.Column, error)
	UpdatePosition(columnId common.Id, newName int) (*entity.Column, error)
	UpdateName(columnId common.Id, newName string) (*entity.Column, error)
	Delete(columnId common.Id) error
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
