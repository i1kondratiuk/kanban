package service

import (
	"errors"

	"github.com/i1kondratiuk/kanban/domain/entity/common"
	"github.com/i1kondratiuk/kanban/domain/repository"
)

// BoardService represents the service to handle business rules related to Kanban Boards
type BoardService interface {
	IsColumnDeletable(storedBoardId common.Id, numberOfColumnsToDelete int) (bool, error)
}

// BoardServiceImpl is the implementation of BoardService
type BoardServiceImpl struct{}

var boardService BoardService

// GetBoardService returns a BoardService
func GetBoardService() BoardService {
	return boardService
}

// InitBoardService injects BoardService with its implementation
func InitBoardService(a BoardService) {
	boardService = a
}

// BoardServiceImpl implements the ColumnService interface
var _ BoardService = &BoardServiceImpl{}

// Checks whether the board has columns associated with it
func (s *BoardServiceImpl) IsColumnDeletable(storedBoardId common.Id, numberOfColumnsToDelete int) (bool, error) {
	storedColumnsNumber, err := repository.GetColumnRepository().CountAllBy(storedBoardId)

	if err != nil {
		return false, err
	}

	if storedColumnsNumber-numberOfColumnsToDelete < 1 {
		return false, errors.New("every board should have at least one column")
	}

	return true, nil
}
