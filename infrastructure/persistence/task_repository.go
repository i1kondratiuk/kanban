package persistence

import (
	"database/sql"
	"errors"

	"github.com/i1kondratiuk/kanban/domain/aggregate"
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
	"github.com/i1kondratiuk/kanban/domain/repository"
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
	return nil, errors.New("GetTaskWithAllCommentsGroupedByCreatedDateTime: implement me")
}

func (t TaskRepositoryImpl) GetBy(taskId common.Id) (*entity.Task, error) {
	return nil, errors.New("GetBy: implement me")
}

func (t TaskRepositoryImpl) GetAllBy(parentColumnId common.Id) ([]*entity.Task, error) {
	return nil, errors.New("GetAllBy: implement me")
}

func (t TaskRepositoryImpl) Insert(newTask *entity.Task) (*entity.Task, error) {
	return nil, errors.New("Insert: implement me")
}

func (t TaskRepositoryImpl) Update(modifiedTask *entity.Task) (*entity.Task, error) {
	return nil, errors.New("Update: implement me")
}

func (t TaskRepositoryImpl) UpdateName(storedTaskId common.Id, newName string) (*entity.Task, error) {
	return nil, errors.New("UpdateName: implement me")
}

func (t TaskRepositoryImpl) UpdateDescription(storedTaskId common.Id, newDescription string) (*entity.Task, error) {
	return nil, errors.New("UpdateDescription: implement me")
}

func (t TaskRepositoryImpl) UpdateParentColumn(storedTaskId common.Id, newParentColumnId common.Id) (*entity.Task, error) {
	return nil, errors.New("UpdateParentColumn: implement me")
}

func (t TaskRepositoryImpl) UpdatePriority(storedTaskId common.Id, priority int) (*entity.Task, error) {
	return nil, errors.New("UpdatePriority: implement me")
}

func (t TaskRepositoryImpl) Delete(storedTaskId common.Id) error {
	return errors.New("Delete: implement me")
}
