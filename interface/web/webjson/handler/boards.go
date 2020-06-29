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

// BoardManagerAppHandler ...
type BoardManagerAppHandler struct {
	BoardManagerApp api.BoardManagerApp
}

// AddRoutes adds BoardManagerAppHandler routs
func (h BoardManagerAppHandler) AddRoutes(r *mux.Router) {
	r.HandleFunc("/boards", h.getAllBoards).Methods("GET")
	r.HandleFunc("/boards", h.CreateBoard).Methods("POST")

	r.HandleFunc("/boards/{"+boardIdAnchor+"}", h.getBoard).Methods("GET")
	r.HandleFunc("/boards/{"+boardIdAnchor+"}", h.updateBoard).Methods("PUT")
	r.HandleFunc("/boards/{"+boardIdAnchor+"}", h.deleteBoard).Methods("DELETE")
}

func (h BoardManagerAppHandler) getAllBoards(w http.ResponseWriter, r *http.Request) {
	h.BoardManagerApp = api.GetBoardManagerApp()

	boards, err := h.BoardManagerApp.GetAllBoardsSortedByNameAsc()

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to get boards; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, boards)
}

func (h BoardManagerAppHandler) CreateBoard(w http.ResponseWriter, r *http.Request) {
	h.BoardManagerApp = api.GetBoardManagerApp()

	var newBoard entity.Board

	if err := json.NewDecoder(r.Body).Decode(&newBoard); err != nil {
		respondError(w, http.StatusNotFound, "failed to create the board; "+err.Error())
		return
	}

	newBoardStored, err := h.BoardManagerApp.Create(&newBoard)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to create the board; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, newBoardStored)
}

func (h BoardManagerAppHandler) getBoard(w http.ResponseWriter, r *http.Request) {
	h.BoardManagerApp = api.GetBoardManagerApp()

	params := mux.Vars(r)
	boardId, err := strconv.ParseInt(params[boardIdAnchor], 10, 64)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to get the board; "+err.Error())
		return
	}

	retrievedBoard, err := h.BoardManagerApp.Get(common.Id(boardId))

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to get the board; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, retrievedBoard)
}

func (h BoardManagerAppHandler) updateBoard(w http.ResponseWriter, r *http.Request) {
	h.BoardManagerApp = api.GetBoardManagerApp()

	params := mux.Vars(r)
	modifiedBoardId, err := strconv.ParseInt(params[boardIdAnchor], 10, 64)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to update the board; "+err.Error())
		return
	}

	var modifiedBoard entity.Board

	if err := json.NewDecoder(r.Body).Decode(&modifiedBoard); err != nil {
		respondError(w, http.StatusNotFound, "failed to update the board; "+err.Error())
		return
	}

	modifiedBoard.Id = common.Id(modifiedBoardId)
	updatedBoard, err := h.BoardManagerApp.Update(&modifiedBoard)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to update the board; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, updatedBoard)
}

func (h BoardManagerAppHandler) deleteBoard(w http.ResponseWriter, r *http.Request) {
	h.BoardManagerApp = api.GetBoardManagerApp()

	boardId, err := strconv.ParseInt(mux.Vars(r)[boardIdAnchor], 10, 64)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to delete the board; "+err.Error())
		return
	}

	if err := h.BoardManagerApp.Delete(common.Id(boardId)); err != nil {
		respondError(w, http.StatusNotFound, "failed to delete the board; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, "Deleted Successfully")
}
