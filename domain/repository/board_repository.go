package repository

import (
	"github.com/i1kondratiuk/kanban/domain/aggregate"
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
)

// BoardRepository represents a storage of all existing boards
type BoardRepository interface {
	GetAllSortedByNameAsc() ([]*aggregate.BoardAggregate, error)
	Insert(newBoard *entity.Board) (*entity.Board, error)
	Update(modifiedBoard *entity.Board) (*entity.Board, error)
	Delete(storedBoardId common.Id) error
}

var boardRepository BoardRepository

// GetBoardRepository returns the BoardRepository
func GetBoardRepository() BoardRepository {
	return boardRepository
}

// InitBoardRepository injects BoardRepository with its implementation
func InitBoardRepository(r BoardRepository) {
	boardRepository = r
}
