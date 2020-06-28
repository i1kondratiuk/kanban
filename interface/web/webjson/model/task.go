package entity

import (
	"github.com/i1kondratiuk/kanban/application/model"
)

// Task represents the model exposed to the API client
type Task struct {
	model.Task
}
