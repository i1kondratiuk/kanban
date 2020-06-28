package dto

import (
	"github.com/i1kondratiuk/kanban/application/model"
	"github.com/i1kondratiuk/kanban/domain/aggregate"
)

func NewBoards(bas []*aggregate.BoardAggregate) []*model.Board {
	var bs = make([]*model.Board, len(bas)-1)

	for i, ba := range bas {
		bs[i] = NewBoard(ba)
	}

	return bs
}

func NewBoard(ba *aggregate.BoardAggregate) *model.Board {
	return &model.Board{
		Id:          ba.BoardAggregateRoot.Id,
		Name:        ba.BoardAggregateRoot.Name,
		Description: ba.BoardAggregateRoot.Description,
		Columns:     NewColumns(ba.ColumnAggregates),
	}
}
