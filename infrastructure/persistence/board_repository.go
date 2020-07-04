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
	if b.db == nil {
		return nil, errors.New("database error")
	}

	rows, err := b.db.Query(
		"SELECT b.id, b.name, b.description, c.id, c.name, c.position FROM boards b LEFT JOIN columns c ON c.board_id = b.id WHERE b.id = $1",
		boardId,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	board := aggregate.BoardAggregate{
		BoardAggregateRoot: &entity.Board{},
		ColumnAggregates:   make([]*aggregate.ColumnAggregate, 0),
	}

	for rows.Next() {
		var (
			boardName        sql.NullString
			boardDescription sql.NullString
			columnId         sql.NullInt64
			columnName       sql.NullString
			columnPosition   sql.NullInt32
		)

		err = rows.Scan(
			&board.BoardAggregateRoot.Id,
			&boardName,
			&boardDescription,
			&columnId,
			&columnName,
			&columnPosition,
		)

		if err != nil {
			return nil, err
		}

		if boardName.Valid {
			board.BoardAggregateRoot.Name = boardName.String
		}

		if boardDescription.Valid {
			board.BoardAggregateRoot.Description = boardDescription.String
		}

		if columnId.Valid {
			column := aggregate.ColumnAggregate{
				ColumnAggregateRoot: &entity.Column{
					Board: *board.BoardAggregateRoot,
				},
			}

			column.ColumnAggregateRoot.Id = common.Id(columnId.Int64)

			if columnName.Valid {
				column.ColumnAggregateRoot.Name = columnName.String
			}

			if columnPosition.Valid {
				column.ColumnAggregateRoot.Position = int(columnPosition.Int32)
			}

			board.ColumnAggregates = append(board.ColumnAggregates, &column)
		}
	}

	return &board, nil
}

func (b BoardRepositoryImpl) Insert(newBoard *entity.Board) (*entity.Board, error) {
	var boardId int64
	if err := b.db.QueryRow(
		"INSERT INTO boards (name, description) VALUES ($1, $2) RETURNING id",
		newBoard.Name,
		newBoard.Description,
	).Scan(&boardId); err != nil {
		return nil, err
	}

	newBoard.Id = common.Id(boardId)

	return newBoard, nil
}

func (b BoardRepositoryImpl) Update(modifiedBoard *entity.Board) (*entity.Board, error) {
	_, err := b.db.Exec(
		"UPDATE boards SET name = $1, description = $3 WHERE id = $3",
		modifiedBoard.Name,
		modifiedBoard.Description,
		modifiedBoard.Id,
	)
	if err != nil {
		return nil, err
	}

	return modifiedBoard, nil
}

func (b BoardRepositoryImpl) Delete(storedBoardId common.Id) error {
	res, err := b.db.Exec("DELETE FROM boards WHERE id = $1", storedBoardId)

	if err == nil {
		count, err := res.RowsAffected()
		if err != nil {
			return err
		} else if count != 1 {
			return errors.New("the record cannot be found, thus it is not deleted")
		}
	}

	return nil
}
