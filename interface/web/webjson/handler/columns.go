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

// ColumnManagerAppHandler handler
type ColumnManagerAppHandler struct {
	ColumnManagerApp api.ColumnManagerApp
}

// AddRoutes adds ColumnManagerAppHandler routs
func (h ColumnManagerAppHandler) AddRoutes(r *mux.Router) { // TODO get rid of the redundant path prefix for the subresource
	r.HandleFunc("/boards/{"+boardIdAnchor+"}/columns", h.GetAllBoardColumns).Methods("GET")
	r.HandleFunc("/boards/{"+boardIdAnchor+"}/columns", h.CreateBoardColumn).Methods("POST")

	r.HandleFunc("/boards/{"+boardIdAnchor+"}/columns/{"+columnIdAnchor+"}", h.GetColumn).Methods("GET")
	r.HandleFunc("/boards/{"+boardIdAnchor+"}/columns/{"+columnIdAnchor+"}", h.UpdateColumn).Methods("PUT")
	r.HandleFunc("/boards/{"+boardIdAnchor+"}/columns/{"+columnIdAnchor+"}", h.DeleteColumn).Methods("DELETE")

	r.HandleFunc("/boards/{"+boardIdAnchor+"}/columns/{"+columnIdAnchor+"}/name", h.UpdateColumnName).Methods("PUT")
	r.HandleFunc("/boards/{"+boardIdAnchor+"}/columns/{"+columnIdAnchor+"}/position", h.UpdateColumnPosition).Methods("PUT")
}

func (h ColumnManagerAppHandler) GetAllBoardColumns(w http.ResponseWriter, r *http.Request) {
	h.ColumnManagerApp = api.GetColumnManagerApp()

	params := mux.Vars(r)
	boardIdInt64, err := strconv.ParseInt(params[boardIdAnchor], 10, 64)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to get the board; "+err.Error())
		return
	}

	boardId := common.Id(boardIdInt64)
	storedColumns, err := h.ColumnManagerApp.GetAllColumnsWithRelatedTasks(&boardId)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to get columns; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, storedColumns)
}

func (h ColumnManagerAppHandler) CreateBoardColumn(w http.ResponseWriter, r *http.Request) {
	h.ColumnManagerApp = api.GetColumnManagerApp()

	var newColumn entity.Column

	if err := json.NewDecoder(r.Body).Decode(&newColumn); err != nil {
		respondError(w, http.StatusNotFound, "failed to create the column; "+err.Error())
		return
	}

	newColumnStored, err := h.ColumnManagerApp.Create(&newColumn)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to create the column; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, newColumnStored)
}

func (h ColumnManagerAppHandler) GetColumn(w http.ResponseWriter, r *http.Request) {
	h.ColumnManagerApp = api.GetColumnManagerApp()

	params := mux.Vars(r)
	columnIdInt64, err := strconv.ParseInt(params[columnIdAnchor], 10, 64)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to get the column; "+err.Error())
		return
	}

	columnId := common.Id(columnIdInt64)

	storedColumn, err := h.ColumnManagerApp.GetColumn(&columnId)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to get the column; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, storedColumn)
}

func (h ColumnManagerAppHandler) UpdateColumn(w http.ResponseWriter, r *http.Request) {
	h.ColumnManagerApp = api.GetColumnManagerApp()

	params := mux.Vars(r)
	modifiedColumnId, err := strconv.ParseInt(params[columnIdAnchor], 10, 64)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to update the column; "+err.Error())
		return
	}

	var modifiedColumn entity.Column

	if err := json.NewDecoder(r.Body).Decode(&modifiedColumn); err != nil {
		respondError(w, http.StatusNotFound, "failed to update the column; "+err.Error())
		return
	}

	modifiedColumn.Id = common.Id(modifiedColumnId)

	updatedColumn, err := h.ColumnManagerApp.Update(&modifiedColumn)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to update the column; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, updatedColumn)
}

func (h ColumnManagerAppHandler) UpdateColumnName(w http.ResponseWriter, r *http.Request) {
	h.ColumnManagerApp = api.GetColumnManagerApp()

	params := mux.Vars(r)
	columnIdInt64, err := strconv.ParseInt(params[columnIdAnchor], 10, 64)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to update the column name; "+err.Error())
		return
	}

	var newName string

	if err := json.NewDecoder(r.Body).Decode(&newName); err != nil {
		respondError(w, http.StatusNotFound, "failed to update the column name; "+err.Error())
		return
	}

	columnId := common.Id(columnIdInt64)

	updatedColumn, err := h.ColumnManagerApp.Rename(columnId, newName)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to update the column name; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, updatedColumn)
}

func (h ColumnManagerAppHandler) UpdateColumnPosition(w http.ResponseWriter, r *http.Request) {
	h.ColumnManagerApp = api.GetColumnManagerApp()

	params := mux.Vars(r)
	columnIdUntyped, err := strconv.ParseInt(params[columnIdAnchor], 10, 64)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to update the column position; "+err.Error())
		return
	}

	var newPosition int

	if err := json.NewDecoder(r.Body).Decode(&newPosition); err != nil {
		respondError(w, http.StatusNotFound, "failed to update the column position; "+err.Error())
		return
	}

	columnId := common.Id(columnIdUntyped)

	updatedColumn, err := h.ColumnManagerApp.ChangePosition(columnId, newPosition)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to update the column position; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, updatedColumn)
}

func (h ColumnManagerAppHandler) DeleteColumn(w http.ResponseWriter, r *http.Request) {
	h.ColumnManagerApp = api.GetColumnManagerApp()
	params := mux.Vars(r)
	columnIdInt64, err := strconv.ParseInt(params[columnIdAnchor], 10, 64)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to create the column; "+err.Error())
		return
	}

	columnId := common.Id(columnIdInt64)

	if err := h.ColumnManagerApp.Delete(columnId); err != nil {
		respondError(w, http.StatusNotFound, "failed to create the column; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, "the column was deleted successfully")
}
