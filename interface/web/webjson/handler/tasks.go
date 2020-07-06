package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

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
func (h TaskManagerAppHandler) AddRoutes(r *mux.Router) { // TODO get rid of the redundant path prefix for the subresource
	r.HandleFunc("/columns/{"+columnIdAnchor+"}/tasks", h.GetAllTasks).Methods("GET")
	r.HandleFunc("/tasks", h.CreateTask).Methods("POST")

	r.HandleFunc("/tasks/{"+taskIdAnchor+"}", h.GetTaskWithComments).Methods("GET")
	r.HandleFunc("/tasks/{"+taskIdAnchor+"}", h.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{"+taskIdAnchor+"}", h.DeleteTask).Methods("DELETE")

	r.HandleFunc("/tasks/{"+taskIdAnchor+"}/priority", h.ChangeTaskPriority).Methods("PUT")
	r.HandleFunc("/tasks/{"+taskIdAnchor+"}/status", h.ChangeTaskStatus).Methods("PUT")
	r.HandleFunc("/tasks/{"+taskIdAnchor+"}/name", h.ChangeTaskName).Methods("PUT")
	r.HandleFunc("/tasks/{"+taskIdAnchor+"}/description", h.ChangeTaskDescription).Methods("PUT")
}

func (h TaskManagerAppHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	h.TaskManagerApp = api.GetTaskManagerApp()

	columnId, err := strconv.ParseInt(mux.Vars(r)[columnIdAnchor], 10, 64)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to get tasks; "+err.Error())
		return
	}

	storedTasks, err := h.TaskManagerApp.GetTasksBy(common.Id(columnId))

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to get tasks; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, storedTasks)
}

func (h TaskManagerAppHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	h.TaskManagerApp = api.GetTaskManagerApp()

	var newTask entity.Task

	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		respondError(w, http.StatusNotFound, "failed to create the task; "+err.Error())
		return
	}

	newTaskStored, err := h.TaskManagerApp.Create(&newTask)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to create the task; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, newTaskStored)
}

func (h TaskManagerAppHandler) GetTaskWithComments(w http.ResponseWriter, r *http.Request) {
	h.TaskManagerApp = api.GetTaskManagerApp()

	params := mux.Vars(r)
	taskIdInt64, err := strconv.ParseInt(params[taskIdAnchor], 10, 64)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to get the task; "+err.Error())
		return
	}

	taskId := common.Id(taskIdInt64)

	storedTasks, err := h.TaskManagerApp.GetTaskWithAllCommentsGroupedByCreatedDateTime(taskId)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to get the task; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, storedTasks)
}

func (h TaskManagerAppHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	h.TaskManagerApp = api.GetTaskManagerApp()

	params := mux.Vars(r)
	modifiedTaskId, err := strconv.ParseInt(params[taskIdAnchor], 10, 64)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to update the task; "+err.Error())
		return
	}

	var modifiedTask entity.Task

	if err := json.NewDecoder(r.Body).Decode(&modifiedTask); err != nil {
		respondError(w, http.StatusNotFound, "failed to update the task; "+err.Error())
		return
	}

	modifiedTask.Id = common.Id(modifiedTaskId)

	updatedTasks, err := h.TaskManagerApp.Update(&modifiedTask)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to update the task; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, updatedTasks)
}

func (h TaskManagerAppHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	h.TaskManagerApp = api.GetTaskManagerApp()

	params := mux.Vars(r)
	taskId, err := strconv.ParseInt(params[taskIdAnchor], 10, 64)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to delete the task; "+err.Error())
		return
	}

	if err := h.TaskManagerApp.DeleteWithAllComments(common.Id(taskId)); err != nil {
		respondError(w, http.StatusNotFound, "failed to delete the task; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, "the task and all related comments were deleted successfully")
}

func (h TaskManagerAppHandler) ChangeTaskPriority(w http.ResponseWriter, r *http.Request) {
	h.TaskManagerApp = api.GetTaskManagerApp()

	params := mux.Vars(r)
	taskIdInt64, err := strconv.ParseInt(params[taskIdAnchor], 10, 64)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to update the task priority; "+err.Error())
		return
	}

	var newPriority int

	if err := json.NewDecoder(r.Body).Decode(&newPriority); err != nil {
		respondError(w, http.StatusNotFound, "failed to update the task priority; "+err.Error())
		return
	}

	taskId := common.Id(taskIdInt64)

	if err = h.TaskManagerApp.Prioritize(taskId, newPriority); err != nil {
		respondError(w, http.StatusNotFound, "failed to update the task priority; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, "the task priority was updates successfully")
}

func (h TaskManagerAppHandler) ChangeTaskStatus(w http.ResponseWriter, r *http.Request) {
	h.TaskManagerApp = api.GetTaskManagerApp()

	params := mux.Vars(r)
	taskIdInt64, err := strconv.ParseInt(params[taskIdAnchor], 10, 64)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to update the task status; "+err.Error())
		return
	}

	var columnIdStr string

	if err := json.NewDecoder(r.Body).Decode(&columnIdStr); err != nil {
		respondError(w, http.StatusNotFound, "failed to update the task priority; "+err.Error())
		return
	}

	columnIdInt64, err := strconv.ParseInt(columnIdStr, 10, 64)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to update the task status; "+err.Error())
		return
	}

	columnId := common.Id(columnIdInt64)
	taskId := common.Id(taskIdInt64)

	if err = h.TaskManagerApp.ChangeStatus(taskId, columnId); err != nil {
		respondError(w, http.StatusNotFound, "failed to update the task status; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, "the task status was updates successfully")
}

func (h TaskManagerAppHandler) ChangeTaskName(w http.ResponseWriter, r *http.Request) {
	h.TaskManagerApp = api.GetTaskManagerApp()

	params := mux.Vars(r)
	taskIdInt64, err := strconv.ParseInt(params[taskIdAnchor], 10, 64)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to update the task; "+err.Error())
		return
	}

	var newName string

	if err := json.NewDecoder(r.Body).Decode(&newName); err != nil {
		respondError(w, http.StatusNotFound, "failed to update the task priority; "+err.Error())
		return
	}

	taskId := common.Id(taskIdInt64)

	if err = h.TaskManagerApp.ChangeName(taskId, newName); err != nil {
		respondError(w, http.StatusNotFound, "failed to update the task name; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, "the task name was updates successfully")
}

func (h TaskManagerAppHandler) ChangeTaskDescription(w http.ResponseWriter, r *http.Request) {
	h.TaskManagerApp = api.GetTaskManagerApp()

	params := mux.Vars(r)
	taskIdInt64, err := strconv.ParseInt(params[taskIdAnchor], 10, 64)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to update the task description; "+err.Error())
		return
	}

	var newDescription string

	if err := json.NewDecoder(r.Body).Decode(&newDescription); err != nil {
		respondError(w, http.StatusNotFound, "failed to update the task description; "+err.Error())
		return
	}

	taskId := common.Id(taskIdInt64)

	if err = h.TaskManagerApp.ChangeDescription(taskId, newDescription); err != nil {
		respondError(w, http.StatusNotFound, "failed to update the task description; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, "the task description was updates successfully")
}
