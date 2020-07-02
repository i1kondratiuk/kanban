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
	if b.db == nil {
		return nil, errors.New("database error")
	}

	rows, err := b.db.Query("SELECT id, name, description FROM boards")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	boards := make([]*entity.Board, 0)
	for rows.Next() {
		board := entity.Board{}

		var (
			name        sql.NullString
			description sql.NullString
		)

		err = rows.Scan(
			&board.Id,
			&name,
			&description,
		)

		if err != nil {
			return nil, err
		}

		if name.Valid {
			board.Name = name.String
		}

		if description.Valid {
			board.Description = description.String
		}

		boards = append(boards, &board)
	}

	return boards, nil
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
