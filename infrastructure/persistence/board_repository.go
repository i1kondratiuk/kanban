package persistence

import (
	"database/sql"
	"errors"

	"github.com/i1kondratiuk/kanban/domain/aggregate"
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
	"github.com/i1kondratiuk/kanban/domain/repository"
)

// BoardRepositoryImpl is the implementation of BoardRepository
type BoardRepositoryImpl struct {
	db *sql.DB
}

// BoardRepositoryImpl implements the domain BoardRepository interface
var _ repository.BoardRepository = &BoardRepositoryImpl{}

// BoardRepository returns initialized BoardRepositoryImpl
func NewBoardRepository(db *sql.DB) repository.BoardRepository {
	return &BoardRepositoryImpl{db: db}
}

func (b BoardRepositoryImpl) GetAllSortedByNameAsc() ([]*entity.Board, error) {
	return nil, errors.New("GetAllSortedByNameAsc: implement me")
}

func (b BoardRepositoryImpl) GetBy(boardId common.Id) (*aggregate.BoardAggregate, error) {
	return nil, errors.New("GetBy: implement me")
}

func (b BoardRepositoryImpl) Insert(newBoard *entity.Board) (*entity.Board, error) {
	return nil, errors.New("Insert: implement me")
}

func (b BoardRepositoryImpl) Update(modifiedBoard *entity.Board) (*entity.Board, error) {
	return nil, errors.New("Update: implement me")
}

func (b BoardRepositoryImpl) Delete(storedBoardId common.Id) error {
	return errors.New("Delete: implement me")
}
