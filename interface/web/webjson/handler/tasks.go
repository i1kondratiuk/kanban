package handler

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/i1kondratiuk/kanban/application/api"
)

// TaskManagerAppHandler handler
type TaskManagerAppHandler struct {
	TaskManagerApp api.TaskManagerApp
}

// AddRoutes adds TaskManagerAppHandler routs
func (h TaskManagerAppHandler) AddRoutes(r *mux.Router) {
	r.HandleFunc("/boards/{boardId}/columns/{columnId}/tasks", h.GetAllTasks).Methods("GET")
	r.HandleFunc("/boards/{boardId}/columns/{columnId}/tasks", h.CreateTask).Methods("POST")

	r.HandleFunc("/boards/{boardId}/columns/{columnId}/tasks/{taskId}", h.GetTask).Methods("GET")
	r.HandleFunc("/boards/{boardId}/columns/{columnId}/tasks/{taskId}", h.UpdateTask).Methods("PUT")
	r.HandleFunc("/boards/{boardId}/columns/{columnId}/tasks/{taskId}", h.DeleteTask).Methods("DELETE")

	r.HandleFunc("/boards/{boardId}/columns/{columnId}/tasks/{taskId}/priority", h.ChangeTaskPriority).Methods("PUT")
	r.HandleFunc("/boards/{boardId}/columns/{columnId}/tasks/{taskId}/status", h.ChangeTaskStatus).Methods("PUT")
	r.HandleFunc("/boards/{boardId}/columns/{columnId}/tasks/{taskId}/name", h.ChangeTaskName).Methods("PUT")
	r.HandleFunc("/boards/{boardId}/columns/{columnId}/tasks/{taskId}/description", h.ChangeTaskDescription).Methods("PUT")
}

func (h TaskManagerAppHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
}

func (h TaskManagerAppHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
}

func (h TaskManagerAppHandler) GetTask(w http.ResponseWriter, r *http.Request) {
}

func (h TaskManagerAppHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
}

func (h TaskManagerAppHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
}

func (h TaskManagerAppHandler) ChangeTaskPriority(w http.ResponseWriter, r *http.Request) {
}

func (h TaskManagerAppHandler) ChangeTaskStatus(w http.ResponseWriter, r *http.Request) {
}

func (h TaskManagerAppHandler) ChangeTaskName(w http.ResponseWriter, r *http.Request) {
}

func (h TaskManagerAppHandler) ChangeTaskDescription(w http.ResponseWriter, r *http.Request) {
}
