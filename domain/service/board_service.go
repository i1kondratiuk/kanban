package service

import (
	"errors"

	"github.com/i1kondratiuk/kanban/domain/entity/common"
	"github.com/i1kondratiuk/kanban/domain/repository"
)

// BoardService represents the service to handle business rules related to Kanban Boards
type BoardService interface {
	HasColumns(storedBoardId common.Id) error
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

// Checks whether the board has columns associated with it
func (a *BoardServiceImpl) HasColumns(storedBoardId common.Id) error {
	storedColumnsNumber, err := repository.GetColumnRepository().CountAllBy(storedBoardId)

	if err != nil {
		panic(err)
	}

	if storedColumnsNumber < 1 {
		return errors.New("every board should have at least one column")
	}

	return nil
}
