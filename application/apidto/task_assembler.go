package apidto

import (
	"github.com/i1kondratiuk/kanban/application/apimodel"
	"github.com/i1kondratiuk/kanban/domain/aggregate"
	"github.com/i1kondratiuk/kanban/domain/entity"
)

func NewTasksFromEntities(tes []*entity.Task) []*apimodel.Task {
	var ts = make([]*apimodel.Task, 0, len(tes))

	for _, te := range tes {
		ts = append(ts, NewTaskFromEntity(te))
	}

	return ts
}

func NewTaskFromEntity(te *entity.Task) *apimodel.Task {
	return &apimodel.Task{
		Id:          te.Id,
		ColumnId:    te.ColumnId,
		Name:        te.Name,
		Description: te.Description,
		Priority:    te.Priority,
	}
}

func NewTasksFromAggregates(tas []*aggregate.TaskAggregate) []*apimodel.Task {
	var ts = make([]*apimodel.Task, 0, len(tas))

	for _, ta := range tas {
		ts = append(ts, NewTaskFromAggregate(ta))
	}

	return ts
}

func NewTaskFromAggregate(ta *aggregate.TaskAggregate) *apimodel.Task {
	var comments = make([]*apimodel.Comment, 0, len(ta.Comments))

	for _, comment := range ta.Comments {
		comments = append(
			comments,
			&apimodel.Comment{
				Id:              comment.Id,
				ParentId:        comment.ParentId,
				CreatedDateTime: comment.CreatedDateTime,
				Comment:         comment.Comment,
			},
		)
	}

	return &apimodel.Task{
		Id:          ta.TaskAggregateRoot.Id,
		ColumnId:    ta.TaskAggregateRoot.ColumnId,
		Name:        ta.TaskAggregateRoot.Name,
		Description: ta.TaskAggregateRoot.Description,
		Priority:    ta.TaskAggregateRoot.Priority,
		Comments:    comments,
	}
}
