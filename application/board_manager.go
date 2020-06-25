package application

import (
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
	"github.com/i1kondratiuk/kanban/domain/repository"
	"github.com/i1kondratiuk/kanban/domain/service"
)

// BoardManagerApp represents BoardManagerApp application to be called by interface layer
type BoardManagerApp interface {
	GetAllBoardsSortedByNameAsc() ([]*entity.Board, error)
	Create(newBoard *entity.Board) (*entity.Board, error)
	Update(modifiedBoard *entity.Board) (*entity.Board, error)
	Delete(storedBoardId common.Id) error
}

// BoardManagerAppImpl is the implementation of BoardManagerApp
type BoardManagerAppImpl struct{}

var boardManagerApp BoardManagerApp

// InitBoardManagerApp injects implementation for KanbanBoardApp application
func InitBoardManagerApp(a BoardManagerApp) {
	boardManagerApp = a
}

// GetBoardManagerApp returns KanbanBoardApp application
func GetBoardManagerApp() BoardManagerApp {
	return boardManagerApp
}

// BoardManagerAppImpl implements the KanbanBoardApp interface
var _ BoardManagerApp = &BoardManagerAppImpl{}

func (a *BoardManagerAppImpl) GetAllBoardsSortedByNameAsc() ([]*entity.Board, error) {
	storedBoards, err := repository.GetBoardRepository().GetAllSortedByNameAsc()

	if err != nil {
		return nil, err
	}

	return storedBoards, nil
}

func (a *BoardManagerAppImpl) Create(newBoard *entity.Board) (*entity.Board, error) {
	insertedBoard, err := repository.GetBoardRepository().Insert(newBoard)

	if err != nil {
		return nil, err
	}

	_, err = repository.GetColumnRepository().Insert(service.GetColumnService().CreateDefaultColumn(insertedBoard.Id))

	if err != nil {
		return nil, err
	}

	return insertedBoard, nil
}

func (a *BoardManagerAppImpl) Update(modifiedBoard *entity.Board) (*entity.Board, error) {
	updatedBoard, err := repository.GetBoardRepository().Update(modifiedBoard)

	if err != nil {
		return nil, err
	}

	return updatedBoard, nil
}

func (a *BoardManagerAppImpl) Delete(storedBoardId common.Id) error {
	if err := repository.GetBoardRepository().Delete(storedBoardId); err != nil {
		return err
	}

	return nil
}
