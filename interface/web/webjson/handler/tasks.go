package handler

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/i1kondratiuk/kanban/application/api"
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
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
	h.TaskManagerApp = api.GetTaskManagerApp()

	var columnId common.Id // TODO Implement
	storedTasks, err := h.TaskManagerApp.GetAllColumnTasks(columnId)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to get tasks; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, storedTasks)
}

func (h TaskManagerAppHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	h.TaskManagerApp = api.GetTaskManagerApp()

	newTask := &entity.Task{} // TODO Implement
	newTaskStored, err := h.TaskManagerApp.Create(newTask)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to create the task; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, newTaskStored)
}

func (h TaskManagerAppHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	h.TaskManagerApp = api.GetTaskManagerApp()

	var taskId common.Id // TODO Implement
	storedTasks, err := h.TaskManagerApp.GeTask(taskId)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to get task; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, storedTasks)
}

func (h TaskManagerAppHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	h.TaskManagerApp = api.GetTaskManagerApp()

	modifiedTask := &entity.Task{} // TODO Implement
	updatedTasks, err := h.TaskManagerApp.Update(modifiedTask)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to update task; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, updatedTasks)
}

func (h TaskManagerAppHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	h.TaskManagerApp = api.GetTaskManagerApp()

	var taskId common.Id // TODO Implement

	if err := h.TaskManagerApp.DeleteWithAllComments(taskId); err != nil {
		respondError(w, http.StatusNotFound, "failed to delete task; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, "the task and all related comments were deleted successfully")
}

func (h TaskManagerAppHandler) ChangeTaskPriority(w http.ResponseWriter, r *http.Request) {
	h.TaskManagerApp = api.GetTaskManagerApp()

	var taskId common.Id // TODO Implement
	var newPriority int  // TODO Implement
	updatedTask, err := h.TaskManagerApp.Prioritize(taskId, newPriority)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to update task; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, updatedTask)
}

func (h TaskManagerAppHandler) ChangeTaskStatus(w http.ResponseWriter, r *http.Request) {
	h.TaskManagerApp = api.GetTaskManagerApp()

	var taskId common.Id   // TODO Implement
	var columnId common.Id // TODO Implement
	updatedTask, err := h.TaskManagerApp.ChangeStatus(taskId, columnId)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to update tasks; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, updatedTask)
}

func (h TaskManagerAppHandler) ChangeTaskName(w http.ResponseWriter, r *http.Request) {
	h.TaskManagerApp = api.GetTaskManagerApp()

	var taskId common.Id // TODO Implement
	var newName string   // TODO Implement
	updatedTask, err := h.TaskManagerApp.ChangeName(taskId, newName)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to update task; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, updatedTask)
}

func (h TaskManagerAppHandler) ChangeTaskDescription(w http.ResponseWriter, r *http.Request) {
	h.TaskManagerApp = api.GetTaskManagerApp()

	var taskId common.Id      // TODO Implement
	var newDescription string // TODO Implement
	updatedTask, err := h.TaskManagerApp.ChangeName(taskId, newDescription)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to update task; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, updatedTask)
}
