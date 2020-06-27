package aggregate

import (
	"github.com/i1kondratiuk/kanban/domain/entity"
)

// BoardAggregate represents the board entity as the aggregate root with Columns (its related child entities)
type BoardAggregate struct {
	BoardAggregateRoot *entity.Board
	ColumnAggregates   []*ColumnAggregate
}

// NewBoard creates a new Board Aggregate instance
func (ba *BoardAggregate) NewBoard(b *entity.Board, cs []*ColumnAggregate) *BoardAggregate {
	return &BoardAggregate{
		BoardAggregateRoot: b,
		ColumnAggregates:   cs,
	}
}

// Column represents the column entity as the aggregate root with Tasks (its related child entities)
type ColumnAggregate struct {
	ColumnAggregateRoot *entity.Column
	TaskEntities        []*entity.Task
}

// NewBColumn creates a new Column Aggregate instance
func (ca *ColumnAggregate) NewColumn(c *entity.Column, ts []*entity.Task) *ColumnAggregate {
	return &ColumnAggregate{
		ColumnAggregateRoot: c,
		TaskEntities:        ts,
	}
}
