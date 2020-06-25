package persistence

import (
	"database/sql"

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

// VisitLogRepository returns initialized VisitLogRepositoryImpl
func NewBoardRepository(db *sql.DB) repository.BoardRepository {
	return &BoardRepositoryImpl{db: db}
}

func (b BoardRepositoryImpl) GetAllSortedByNameAsc() ([]*entity.Board, error) {
	panic("implement me")
}

func (b BoardRepositoryImpl) Insert(newBoard *entity.Board) (*entity.Board, error) {
	panic("implement me")
}

func (b BoardRepositoryImpl) Update(modifiedBoard *entity.Board) (*entity.Board, error) {
	panic("implement me")
}

func (b BoardRepositoryImpl) Delete(storedBoardId common.Id) error {
	panic("implement me")
}
