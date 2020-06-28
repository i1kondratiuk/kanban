package apidto

import (
	"github.com/i1kondratiuk/kanban/application/apimodel"
	"github.com/i1kondratiuk/kanban/domain/aggregate"
	"github.com/i1kondratiuk/kanban/domain/entity"
)


func NewBoardFromEntity(be *entity.Board) *apimodel.Board {
	return &apimodel.Board{
		Id:          be.Id,
		Name:        be.Name,
		Description: be.Description,
	}
}

func NewBoardsFromEntity(bes []*entity.Board) []*apimodel.Board {
	var bms = make([]*apimodel.Board, len(bes)-1)

	for i, be := range bes {
		bms[i] = NewBoardFromEntity(be)
	}

	return bms
}

func NewBoardsFromAggregate(bas []*aggregate.BoardAggregate) []*apimodel.Board {
	var bms = make([]*apimodel.Board, len(bas)-1)

	for i, ba := range bas {
		bms[i] = NewBoardFromAggregate(ba)
	}

	return bms
}

func NewBoardFromAggregate(ba *aggregate.BoardAggregate) *apimodel.Board {
	return &apimodel.Board{
		Id:          ba.BoardAggregateRoot.Id,
		Name:        ba.BoardAggregateRoot.Name,
		Description: ba.BoardAggregateRoot.Description,
		Columns:     NewColumns(ba.ColumnAggregates),
	}
}
