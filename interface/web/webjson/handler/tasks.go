package handler

import (
	"net/http"

	"github.com/i1kondratiuk/kanban/application"
)

// TaskManagerAppHandler handler
type TaskManagerAppHandler struct {
	TaskManagerApp application.TaskManagerApp
}

// AddRoutes adds TaskManagerAppHandler routs
func (h TaskManagerAppHandler) AddRoutes() {
	http.HandleFunc("/comments", h.GetAllTasks)
}

func (h TaskManagerAppHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {

}
