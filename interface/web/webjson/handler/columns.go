package handler

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/i1kondratiuk/kanban/application/api"
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
)

// ColumnManagerAppHandler handler
type ColumnManagerAppHandler struct {
	ColumnManagerApp api.ColumnManagerApp
}

// AddRoutes adds ColumnManagerAppHandler routs
func (h ColumnManagerAppHandler) AddRoutes(r *mux.Router) {
	r.HandleFunc("/boards/{boardId}/columns", h.GetAllBoardColumns).Methods("GET")
	r.HandleFunc("/boards/{boardId}/columns", h.CreateBoardColumn).Methods("POST")

	r.HandleFunc("/boards/{boardId}/columns/{columnId}", h.GetColumn).Methods("GET")
	r.HandleFunc("/boards/{boardId}/columns/{columnId}", h.UpdateColumn).Methods("PUT")
	r.HandleFunc("/boards/{boardId}/columns/{columnId}", h.DeleteColumn).Methods("DELETE")

	r.HandleFunc("/boards/{boardId}/columns/{columnId}/name", h.UpdateColumnName).Methods("PUT")
	r.HandleFunc("/boards/{boardId}/columns/{columnId}/position", h.UpdateColumnPosition).Methods("PUT")
}

func (h ColumnManagerAppHandler) GetAllBoardColumns(w http.ResponseWriter, r *http.Request) {
	h.ColumnManagerApp = api.GetColumnManagerApp()

	var boardId common.Id // TODO Implement
	storedColumns, err := h.ColumnManagerApp.GetAllColumnsWithRelatedTasks(&boardId)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to get columns; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, storedColumns)
}

func (h ColumnManagerAppHandler) CreateBoardColumn(w http.ResponseWriter, r *http.Request) {
	h.ColumnManagerApp = api.GetColumnManagerApp()

	newColumn := &entity.Column{} // TODO Implement
	newColumnStored, err := h.ColumnManagerApp.Create(newColumn)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to create the column; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, newColumnStored)
}

func (h ColumnManagerAppHandler) GetColumn(w http.ResponseWriter, r *http.Request) {
	h.ColumnManagerApp = api.GetColumnManagerApp()

	var columnId common.Id // TODO Implement
	storedColumn, err := h.ColumnManagerApp.GetColumn(&columnId)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to get the column; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, storedColumn)
}

func (h ColumnManagerAppHandler) UpdateColumn(w http.ResponseWriter, r *http.Request) {
	h.ColumnManagerApp = api.GetColumnManagerApp()

	modifiedColumn := &entity.Column{} // TODO Implement
	updatedColumn, err := h.ColumnManagerApp.Update(modifiedColumn)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to update the column; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, updatedColumn)
}

func (h ColumnManagerAppHandler) UpdateColumnName(w http.ResponseWriter, r *http.Request) {
	h.ColumnManagerApp = api.GetColumnManagerApp()

	var columnId common.Id // TODO Implement
	var newName string     // TODO Implement
	updatedColumn, err := h.ColumnManagerApp.Rename(columnId, newName)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to update the column name; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, updatedColumn)
}

func (h ColumnManagerAppHandler) UpdateColumnPosition(w http.ResponseWriter, r *http.Request) {
	h.ColumnManagerApp = api.GetColumnManagerApp()

	var columnId common.Id // TODO Implement
	var newPosition int    // TODO Implement
	updatedColumn, err := h.ColumnManagerApp.ChangePosition(columnId, newPosition)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to update the column position; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, updatedColumn)
}

func (h ColumnManagerAppHandler) DeleteColumn(w http.ResponseWriter, r *http.Request) {
	h.ColumnManagerApp = api.GetColumnManagerApp()

	var columnId common.Id // TODO Implement

	if err := h.ColumnManagerApp.Delete(columnId); err != nil {
		respondError(w, http.StatusNotFound, "failed to create the column; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, "the column was deleted successfully")
}
