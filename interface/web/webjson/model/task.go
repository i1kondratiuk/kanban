package entity

import (
	"github.com/i1kondratiuk/kanban/application/apimodel"
)

// Task represents the model exposed to the API client
type Task struct {
	apimodel.Task
}