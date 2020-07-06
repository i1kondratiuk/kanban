package dto

import (
	"github.com/i1kondratiuk/kanban/application/model"
	"github.com/i1kondratiuk/kanban/domain/aggregate"
)

func NewColumns(cas []*aggregate.ColumnAggregate) []*model.Column {
	var cs = make([]*model.Column, len(cas)-1)

	for i, ca := range cas {
		cs[i] = NewColumn(ca)
	}

	return cs
}

func NewColumn(ca *aggregate.ColumnAggregate) *model.Column {

	var tasks = make([]*model.Task, len(ca.TaskEntities)-1)
	for i, taskEntity := range tasks {
		tasks[i] = &model.Task{
			Id:       taskEntity.Id,
			Name:     taskEntity.Name,
			Position: taskEntity.Position,
		}
	}

	return &model.Column{
		Id:       ca.ColumnAggregateRoot.Id,
		Name:     ca.ColumnAggregateRoot.Name,
		Position: ca.ColumnAggregateRoot.Position,
		Tasks:    tasks,
	}
}
