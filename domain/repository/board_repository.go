package repository

import "github.com/i1kondratiuk/kanban/domain/entity"

// BoardRepository represents a storage of all existing boards
type BoardRepository interface {
	GetAllSortedByNameAsc() ([]*entity.Board, error)
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
