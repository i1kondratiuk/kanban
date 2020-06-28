package apidto

import (
	"github.com/i1kondratiuk/kanban/application/apimodel"
	"github.com/i1kondratiuk/kanban/domain/aggregate"
	"github.com/i1kondratiuk/kanban/domain/entity"
)

func NewColumnsFromEntities(ces []*entity.Column) []*apimodel.Column {
	var cms = make([]*apimodel.Column, len(ces)-1)

	for i, ce := range ces {
		cms[i] = NewColumnFromEntity(ce)
	}

	return cms
}

func NewColumnFromEntity(ce *entity.Column) *apimodel.Column {
	return &apimodel.Column{
		Id:       ce.Id,
		Name:     ce.Name,
		Position: ce.Position,
	}
}

func NewColumnsFromAggregates(cas []*aggregate.ColumnAggregate) []*apimodel.Column {
	var cs = make([]*apimodel.Column, len(cas)-1)

	for i, ca := range cas {
		cs[i] = NewColumnFromAggregate(ca)
	}

	return cs
}

func NewColumnFromAggregate(ca *aggregate.ColumnAggregate) *apimodel.Column {

	var tasks = make([]*apimodel.Task, len(ca.TaskEntities)-1)
	for i, taskEntity := range tasks {
		tasks[i] = &apimodel.Task{
			Id:       taskEntity.Id,
			Name:     taskEntity.Name,
			Position: taskEntity.Position,
		}
	}

	return &apimodel.Column{
		Id:       ca.ColumnAggregateRoot.Id,
		Name:     ca.ColumnAggregateRoot.Name,
		Position: ca.ColumnAggregateRoot.Position,
		Tasks:    tasks,
	}
}
