package apidto

import (
	"github.com/i1kondratiuk/kanban/application/apimodel"
	"github.com/i1kondratiuk/kanban/domain/aggregate"
	"github.com/i1kondratiuk/kanban/domain/entity"
)

func NewTasksFromEntities(tes []*entity.Task) []*apimodel.Task {
	var ts = make([]*apimodel.Task, len(tes))

	for i, te := range tes {
		ts[i] = NewTaskFromEntity(te)
	}

	return ts
}

func NewTaskFromEntity(te *entity.Task) *apimodel.Task {
	return &apimodel.Task{
		Id:          te.Id,
		Name:        te.Name,
		Description: te.Description,
		Priority:    te.Priority,
	}
}

func NewTasksFromAggregates(tas []*aggregate.TaskAggregate) []*apimodel.Task {
	var ts = make([]*apimodel.Task, len(tas))

	for i, ta := range tas {
		ts[i] = NewTaskFromAggregate(ta)
	}

	return ts
}

func NewTaskFromAggregate(ta *aggregate.TaskAggregate) *apimodel.Task {

	var comments = make([]*apimodel.Comment, len(ta.Comments))
	for i, comment := range ta.Comments {
		comments[i] = &apimodel.Comment{
			Id:              comment.Id,
			CreatedDateTime: comment.CreatedDateTime,
			Comment:         comment.Comment,
		}
	}

	return &apimodel.Task{
		Id:          ta.Id,
		Name:        ta.Name,
		Description: ta.Description,
		Priority:    ta.Priority,
		Comments:    comments,
	}
}
