package apidto

import (
	"github.com/i1kondratiuk/kanban/application/apimodel"
	"github.com/i1kondratiuk/kanban/domain/aggregate"
	"github.com/i1kondratiuk/kanban/domain/entity"
)

func NewColumnsFromEntities(ces []*entity.Column) []*apimodel.Column {
	var cms = make([]*apimodel.Column, 0, len(ces))

	for _, ce := range ces {
		cms = append(cms, NewColumnFromEntity(ce))
	}

	return cms
}

func NewColumnFromEntity(ce *entity.Column) *apimodel.Column {
	return &apimodel.Column{
		Id:       ce.Id,
		BoardId:  ce.BoardId,
		Name:     ce.Name,
		Position: ce.Position,
	}
}

func NewColumnsFromAggregates(cas []*aggregate.ColumnAggregate) []*apimodel.Column {
	var cs = make([]*apimodel.Column, 0, len(cas))

	for _, ca := range cas {
		cs = append(cs, NewColumnFromAggregate(ca))
	}

	return cs
}

func NewColumnFromAggregate(ca *aggregate.ColumnAggregate) *apimodel.Column {
	column := &apimodel.Column{
		Id:       ca.ColumnAggregateRoot.Id,
		BoardId:  ca.ColumnAggregateRoot.BoardId,
		Name:     ca.ColumnAggregateRoot.Name,
		Position: ca.ColumnAggregateRoot.Position,
	}

	if ca.TaskEntities != nil {
		var tasks = make([]*apimodel.Task, 0, len(ca.TaskEntities))
		for _, taskEntity := range ca.TaskEntities {
			tasks = append(
				tasks,
				&apimodel.Task{
					Id:       taskEntity.Id,
					ColumnId: taskEntity.ColumnId,
					Name:     taskEntity.Name,
					Priority: taskEntity.Priority,
				},
			)
		}

		column.Tasks = tasks
	}

	return column
}
