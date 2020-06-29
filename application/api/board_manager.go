package api

import (
	"github.com/i1kondratiuk/kanban/application/apidto"
	"github.com/i1kondratiuk/kanban/application/apimodel"
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
	"github.com/i1kondratiuk/kanban/domain/repository"
	"github.com/i1kondratiuk/kanban/domain/service"
)

// BoardManagerApp represents BoardManagerApp application to be called by interface layer
type BoardManagerApp interface {
	GetAllBoardsSortedByNameAsc() ([]*apimodel.Board, error)
	Get(boardId common.Id) (*apimodel.Board, error)
	Create(newBoard *entity.Board) (*apimodel.Board, error)      // TODO bulk create
	Update(modifiedBoard *entity.Board) (*apimodel.Board, error) // TODO bulk update
	Delete(storedBoardId common.Id) error                        // TODO bulk delete
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

func (a *BoardManagerAppImpl) GetAllBoardsSortedByNameAsc() ([]*apimodel.Board, error) {
	storedBoards, err := repository.GetBoardRepository().GetAllSortedByNameAsc()

	if err != nil {
		return nil, err
	}

	return apidto.NewBoardsFromEntity(storedBoards), nil
}

func (a *BoardManagerAppImpl) Get(boardId common.Id) (*apimodel.Board, error) {
	storedBoard, err := repository.GetBoardRepository().GetBy(boardId)

	if err != nil {
		return nil, err
	}

	return apidto.NewBoardFromAggregate(storedBoard), nil
}

func (a *BoardManagerAppImpl) Create(newBoard *entity.Board) (*apimodel.Board, error) {
	insertedBoard, err := repository.GetBoardRepository().Insert(newBoard)

	if err != nil {
		return nil, err
	}

	_, err = repository.GetColumnRepository().Insert(service.GetColumnService().CreateDefaultColumn(insertedBoard.Id))

	if err != nil {
		return nil, err
	}

	return apidto.NewBoardFromEntity(insertedBoard), nil
}

func (a *BoardManagerAppImpl) Update(modifiedBoard *entity.Board) (*apimodel.Board, error) {
	updatedBoard, err := repository.GetBoardRepository().Update(modifiedBoard)

	if err != nil {
		return nil, err
	}

	return apidto.NewBoardFromEntity(updatedBoard), nil
}

func (a *BoardManagerAppImpl) Delete(storedBoardId common.Id) error {
	if err := repository.GetBoardRepository().Delete(storedBoardId); err != nil {
		return err
	}

	return nil
}
