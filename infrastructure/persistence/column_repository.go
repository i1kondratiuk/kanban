package persistence

import (
	"database/sql"

	"github.com/i1kondratiuk/kanban/domain/aggregate"
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
	"github.com/i1kondratiuk/kanban/domain/repository"
)

// ColumnRepositoryImpl is the implementation of ColumnRepository
type ColumnRepositoryImpl struct {
	db *sql.DB
}

// ColumnRepositoryImpl implements the domain ColumnRepository interface
var _ repository.ColumnRepository = &ColumnRepositoryImpl{}

// ColumnRepository returns initialized ColumnRepositoryImpl
func NewColumnRepository(db *sql.DB) repository.ColumnRepository {
	return &ColumnRepositoryImpl{db: db}
}

func (c ColumnRepositoryImpl) GetAllWithRelatedTasksBy(parentBoardId common.Id) ([]*aggregate.ColumnAggregate, error) {
	panic("implement me")
}

func (c ColumnRepositoryImpl) GetBy(parentBoardId common.Id) (*entity.Column, error) {
	panic("implement me")
}

func (c ColumnRepositoryImpl) GetByChildTaskId(taskId common.Id) (entity.Column, error) {
	panic("implement me")
}

func (c ColumnRepositoryImpl) GetByParentBoardIdAndPosition(parentBoardId common.Id, position int) (*entity.Column, error) {
	panic("implement me")
}

func (c ColumnRepositoryImpl) GetBoardId(columnId common.Id) (parentBoardId common.Id, err error) {
	panic("implement me")
}

func (c ColumnRepositoryImpl) CountAllBy(parentBoardId common.Id) (int, error) {
	panic("implement me")
}

func (c ColumnRepositoryImpl) Insert(newColumn *entity.Column) (*entity.Column, error) {
	panic("implement me")
}

func (c ColumnRepositoryImpl) Update(updatedColumn *entity.Column) (*entity.Column, error) {
	panic("implement me")
}

func (c ColumnRepositoryImpl) UpdatePosition(columnId common.Id, newName int) (*entity.Column, error) {
	panic("implement me")
}

func (c ColumnRepositoryImpl) UpdateName(columnId common.Id, newName string) (*entity.Column, error) {
	panic("implement me")
}

func (c ColumnRepositoryImpl) Delete(columnId common.Id) error {
	panic("implement me")
}
