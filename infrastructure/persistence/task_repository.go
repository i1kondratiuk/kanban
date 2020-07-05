package persistence

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/i1kondratiuk/kanban/domain/aggregate"
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
	"github.com/i1kondratiuk/kanban/domain/repository"
	"github.com/i1kondratiuk/kanban/domain/value"
)

// TaskRepositoryImpl is the implementation of TaskRepository
type TaskRepositoryImpl struct {
	db *sql.DB
}

// TaskRepositoryImpl implements the domain TaskRepository interface
var _ repository.TaskRepository = &TaskRepositoryImpl{}

// TaskRepository returns initialized TaskRepositoryImpl
func NewTaskRepository(db *sql.DB) repository.TaskRepository {
	return &TaskRepositoryImpl{db: db}
}

func (t TaskRepositoryImpl) GetTaskWithAllCommentsGroupedByCreatedDateTime(taskId common.Id) (*aggregate.TaskAggregate, error) {
	if t.db == nil {
		return nil, errors.New("database error")
	}

	rows, err := t.db.Query(`
		SELECT
		       t.column_id,
		       t.name,
		       t.description,
		       t.priority,
		       c.id,
		       c.body,
		       c.created_at
		FROM tasks t LEFT JOIN comments c ON t.id = c.parent_id
		WHERE t.id = $1
		ORDER BY t.priority, c.created_at DESC`,
		taskId,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("there is no record with id " + strconv.Itoa(int(taskId)))
		}

		return nil, err
	}

	defer rows.Close()

	task := aggregate.TaskAggregate{
		TaskAggregateRoot: &entity.Task{Id: taskId},
		Comments:          []*entity.Comment{},
	}

	for rows.Next() {
		var (
			columnId               sql.NullInt64
			taskName               sql.NullString
			taskDescription        sql.NullString
			taskPriority           sql.NullInt32
			commentId              sql.NullInt64
			commentBody            sql.NullString
			commentCreatedDateTime sql.NullTime
		)

		err = rows.Scan(
			&columnId,
			&taskName,
			&taskDescription,
			&taskPriority,
			&commentId,
			&commentBody,
			&commentCreatedDateTime,
		)

		if err != nil {
			return nil, err
		}

		if columnId.Valid {
			task.TaskAggregateRoot.ColumnId = common.Id(columnId.Int64)
		}

		if taskName.Valid {
			task.TaskAggregateRoot.Name = taskName.String
		}

		if taskDescription.Valid {
			task.TaskAggregateRoot.Description = taskDescription.String
		}

		if taskPriority.Valid {
			task.TaskAggregateRoot.Priority = int(taskPriority.Int32)
		}

		if commentId.Valid {
			comment := entity.Comment{
				Id:       common.Id(commentId.Int64),
				ParentId: taskId,
			}

			if commentBody.Valid {
				comment.Comment.BodyText = value.BodyText(commentBody.String)
			}

			if commentCreatedDateTime.Valid {
				comment.CreatedDateTime = commentCreatedDateTime.Time
			}

			task.Comments = append(task.Comments, &comment)
		}
	}

	return &task, nil
}

func (t TaskRepositoryImpl) GetAllBy(parentColumnId common.Id) ([]*entity.Task, error) {
	if t.db == nil {
		return nil, errors.New("database error")
	}

	rows, err := t.db.Query(`
		SELECT
		       id,
		       name,
		       priority,
		       description
		FROM tasks
		WHERE column_id = $1`,
		parentColumnId,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("there is no record with id " + strconv.Itoa(int(parentColumnId)))
		}

		return nil, err
	}

	defer rows.Close()

	tasks := make([]*entity.Task, 0)

	for rows.Next() {
		var (
			taskId          sql.NullInt64
			taskName        sql.NullString
			taskPriority    sql.NullInt32
			taskDescription sql.NullString
		)

		err = rows.Scan(
			&taskId,
			&taskName,
			&taskPriority,
			&taskDescription,
		)

		if err != nil {
			return nil, err
		}

		if taskId.Valid {
			task := &entity.Task{
				Id:       common.Id(taskId.Int64),
				ColumnId: parentColumnId,
			}

			if taskName.Valid {
				task.Name = taskName.String
			}

			if taskPriority.Valid {
				task.Priority = int(taskPriority.Int32)
			}

			if taskDescription.Valid {
				task.Description = taskDescription.String
			}

			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}

func (t TaskRepositoryImpl) Insert(newTask *entity.Task) (*entity.Task, error) {
	var insertedTaskId int64

	if err := t.db.QueryRow(
		`INSERT INTO tasks (column_id, name, priority, description) VALUES ($1, $2, $3, $4) RETURNING id`,
		int64(newTask.ColumnId),
		newTask.Name,
		newTask.Priority,
		newTask.Description,
	).Scan(&insertedTaskId); err != nil {
		return nil, err
	}

	newTask.Id = common.Id(insertedTaskId)

	return newTask, nil
}

func (t TaskRepositoryImpl) Update(modifiedTask *entity.Task) (*entity.Task, error) {
	_, err := t.db.Exec(
		`UPDATE tasks SET column_id = $1, name = $2, priority = $3, description = $4 WHERE id = $5`,
		int64(modifiedTask.ColumnId),
		modifiedTask.Name,
		modifiedTask.Priority,
		modifiedTask.Description,
		int(modifiedTask.Id),
	)

	if err != nil {
		return nil, err
	}

	return modifiedTask, nil
}

func (t TaskRepositoryImpl) UpdateName(storedTaskId common.Id, newName string) (err error) {
	_, err = t.db.Exec(
		`UPDATE tasks SET name = $1 WHERE id = $2`,
		newName,
		int64(storedTaskId),
	)

	return
}

func (t TaskRepositoryImpl) UpdateDescription(storedTaskId common.Id, newDescription string) (err error) {
	_, err = t.db.Exec(
		`UPDATE tasks SET description = $1 WHERE id = $2`,
		newDescription,
		int64(storedTaskId),
	)

	return
}

func (t TaskRepositoryImpl) UpdateParentColumn(storedTaskId common.Id, newParentColumnId common.Id) (err error) {
	_, err = t.db.Exec(
		`UPDATE tasks SET column_id = $1 WHERE id = $2`,
		int64(newParentColumnId),
		int64(storedTaskId),
	)

	return
}

func (t TaskRepositoryImpl) UpdatePriority(storedTaskId common.Id, newPriority int) (err error) {
	_, err = t.db.Exec(
		`UPDATE tasks SET priority = $1 WHERE id = $2`,
		newPriority,
		int64(storedTaskId),
	)

	return
}

func (t TaskRepositoryImpl) Delete(storedTaskId common.Id) error {
	res, err := t.db.Exec(`DELETE FROM tasks WHERE id = $1`, storedTaskId)

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
