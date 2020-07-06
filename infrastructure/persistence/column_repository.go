package persistence

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/i1kondratiuk/kanban/domain/aggregate"
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
	"github.com/i1kondratiuk/kanban/domain/repository"
)

// ColumnRepositoryImpl is the implementation of ColumnRepository
type ColumnRepositoryImpl struct {
	db *sql.DB
}

// ColumnRepositoryImpl implements the domain ColumnRepository interface
var _ repository.ColumnRepository = &ColumnRepositoryImpl{}

// ColumnRepository returns initialized ColumnRepositoryImpl
func NewColumnRepository(db *sql.DB) repository.ColumnRepository {
	return &ColumnRepositoryImpl{db: db}
}

func (c ColumnRepositoryImpl) GetAllWithRelatedTasksBy(parentBoardId common.Id) ([]*aggregate.ColumnAggregate, error) {
	if c.db == nil {
		return nil, errors.New("database error")
	}

	rows, err := c.db.Query(`
		SELECT
		       c.id,
		       c.name,
		       c.position,
		       t.id,
		       t.name,
		       t.priority,
		       t.description
		FROM columns c LEFT JOIN tasks t ON t.column_id = c.id
		WHERE c.board_id = $1
		ORDER BY c.position, t.priority`,
		parentBoardId,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	columnAggregatesByColumnIds := make(map[common.Id]*aggregate.ColumnAggregate)

	columnAggregates := make([]*aggregate.ColumnAggregate, 0, len(columnAggregatesByColumnIds))

	for rows.Next() {
		var (
			columnId        sql.NullInt64
			columnName      sql.NullString
			columnPosition  sql.NullInt32
			taskId          sql.NullInt64
			taskName        sql.NullString
			taskPriority    sql.NullInt32
			taskDescription sql.NullString
		)

		err = rows.Scan(
			&columnId,
			&columnName,
			&columnPosition,
			&taskId,
			&taskName,
			&taskPriority,
			&taskDescription,
		)

		if err != nil {
			return nil, err
		}

		if columnId.Valid {
			columnId := common.Id(columnId.Int64)
			columnAggregate, found := columnAggregatesByColumnIds[columnId]
			if found {
				columnAggregate.TaskEntities = append(
					columnAggregate.TaskEntities,
					createTask(columnId, taskId, taskName, taskDescription, taskPriority),
				)
			} else {
				columnAggregateToPutInMap := aggregate.ColumnAggregate{
					ColumnAggregateRoot: &entity.Column{
						Id:      columnId,
						BoardId: parentBoardId,
					},
				}

				if columnName.Valid {
					columnAggregateToPutInMap.ColumnAggregateRoot.Name = columnName.String
				}

				if columnPosition.Valid {
					columnAggregateToPutInMap.ColumnAggregateRoot.Position = int(columnPosition.Int32)
				}

				if taskId.Valid {
					columnAggregateToPutInMap.TaskEntities = []*entity.Task{
						createTask(columnId, taskId, taskName, taskDescription, taskPriority),
					}
				}

				columnAggregates = append(columnAggregates, &columnAggregateToPutInMap)

				columnAggregatesByColumnIds[columnId] = &columnAggregateToPutInMap
			}
		}
	}

	return columnAggregates, nil
}

func createTask(columnId common.Id, taskId sql.NullInt64, taskName sql.NullString, taskDescription sql.NullString, taskPriority sql.NullInt32) *entity.Task {
	task := entity.Task{
		Id:       common.Id(int(taskId.Int64)),
		ColumnId: columnId,
	}

	if taskName.Valid {
		task.Name = taskName.String
	}

	if taskDescription.Valid {
		task.Description = taskDescription.String
	}

	if taskPriority.Valid {
		task.Priority = int(taskPriority.Int32)
	}

	return &task
}

func (c ColumnRepositoryImpl) GetBy(columnId common.Id) (*entity.Column, error) {
	if c.db == nil {
		return nil, errors.New("database error")
	}

	var (
		boardId        sql.NullInt64
		columnName     sql.NullString
		columnPosition sql.NullInt32
	)

	err := c.db.QueryRow(`
		SELECT
		       board_id,
		       name,
		       position
		FROM columns
		WHERE id = $1`,
		columnId,
	).Scan(
		&boardId,
		&columnName,
		&columnPosition,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("there is no record with id " + strconv.Itoa(int(columnId)))
		}

		return nil, err
	}

	column := entity.Column{Id: columnId}

	if boardId.Valid {
		column.BoardId = common.Id(boardId.Int64)
	}

	if columnName.Valid {
		column.Name = columnName.String
	}

	if columnPosition.Valid {
		column.Position = int(columnPosition.Int32)
	}

	return &column, nil
}

func (c ColumnRepositoryImpl) GetByChildTaskId(taskId common.Id) (entity.Column, error) {
	panic("implement me")
}

func (c ColumnRepositoryImpl) GetByParentBoardIdAndPosition(parentBoardId common.Id, position int) (*entity.Column, error) {
	panic("implement me")
}

func (c ColumnRepositoryImpl) GetBoardId(columnId common.Id) (parentBoardId common.Id, err error) {
	if c.db == nil {
		return 0, errors.New("database error")
	}

	var parentBoardIdToRetrieve sql.NullInt64
	err = c.db.QueryRow(`SELECT board_id FROM columns WHERE id = $1`, columnId).Scan(&parentBoardIdToRetrieve)

	if parentBoardIdToRetrieve.Valid {
		parentBoardId = common.Id(parentBoardIdToRetrieve.Int64)
	}

	return
}

func (c ColumnRepositoryImpl) CountAllBy(parentBoardId common.Id) (int, error) {
	if c.db == nil {
		return 0, errors.New("database error")
	}

	var count int

	err := c.db.QueryRow(`SELECT COUNT(*) FROM columns WHERE board_id =$1`, parentBoardId).Scan(&count)

	switch {
	case err != nil:
		return 0, err
	default:
		return count, nil
	}
}

func (c ColumnRepositoryImpl) Insert(newColumn *entity.Column) (*entity.Column, error) {
	var insertedColumnId int64

	if err := c.db.QueryRow(
		`INSERT INTO columns (board_id, name, position) VALUES ($1, $2, $3) RETURNING id`,
		int64(newColumn.BoardId),
		newColumn.Name,
		newColumn.Position,
	).Scan(&insertedColumnId); err != nil {
		return nil, err
	}

	newColumn.Id = common.Id(insertedColumnId)

	return newColumn, nil
}

func (c ColumnRepositoryImpl) Update(updatedColumn *entity.Column) (*entity.Column, error) {
	_, err := c.db.Exec(
		`UPDATE columns SET name = $1, position = $3 WHERE id = $3`,
		updatedColumn.Name,
		updatedColumn.Position,
		updatedColumn.Id,
	)

	if err != nil {
		return nil, err
	}

	return updatedColumn, nil
}

func (c ColumnRepositoryImpl) UpdateName(columnId common.Id, newName string) (err error) {
	_, err = c.db.Exec(
		`UPDATE columns SET name = $1 WHERE id = $2`,
		newName,
		columnId,
	)

	return
}

func (c ColumnRepositoryImpl) UpdatePosition(columnId common.Id, newPosition int) (err error) {
	_, err = c.db.Exec(
		`UPDATE columns SET position = $1 WHERE id = $2`,
		newPosition,
		columnId,
	)

	return
}

func (c ColumnRepositoryImpl) Delete(columnId common.Id) error {
	res, err := c.db.Exec(`DELETE FROM columns WHERE id = $1`, columnId)

	if err == nil {
		count, err := res.RowsAffected()
		if err != nil {
			return err
		} else if count != 1 {
			return errors.New("the record cannot be found, thus it is not deleted")
		}
	} else {
		return err
	}

	return nil
}
