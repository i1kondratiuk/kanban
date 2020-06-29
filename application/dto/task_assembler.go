package dto

import (
	"github.com/i1kondratiuk/kanban/application/model"
	"github.com/i1kondratiuk/kanban/domain/aggregate"
)

func NewTasks(tas []*aggregate.TaskAggregate) []*model.Task {
	var ts = make([]*model.Task, len(tas)-1)

	for i, ta := range tas {
		ts[i] = NewTask(ta)
	}

	return ts
}

func NewTask(ta *aggregate.TaskAggregate) *model.Task {

	var comments = make([]*model.Comment, len(ta.Comments)-1)
	for i, comment := range ta.Comments {
		comments[i] = &model.Comment{
			Id:              comment.Id,
			CreatedDateTime: comment.CreatedDateTime,
			Comment:         comment.Comment,
		}
	}

	return &model.Task{
		Id:          ta.Id,
		Name:        ta.Name,
		Description: ta.Description,
		Priority:    ta.Priority,
		Position:    ta.Position,
		Comments:    comments,
	}
}
