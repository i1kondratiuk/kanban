package handler

import (
	"net/http"

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

	r.HandleFunc("/boards/{boardId}", h.getBoard).Methods("GET")
	r.HandleFunc("/boards/{boardId}", h.updateBoard).Methods("PUT")
	r.HandleFunc("/boards/{boardId}", h.deleteBoard).Methods("DELETE")
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

	newBoard := &entity.Board{} // TODO Implement
	newBoardStored, err := h.BoardManagerApp.Create(newBoard)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to create the board; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, newBoardStored)
}

func (h BoardManagerAppHandler) getBoard(w http.ResponseWriter, r *http.Request) {
	h.BoardManagerApp = api.GetBoardManagerApp()

	var boardId common.Id // TODO implement
	retrievedBoard, err := h.BoardManagerApp.Get(boardId)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to get the board; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, retrievedBoard)
}

func (h BoardManagerAppHandler) updateBoard(w http.ResponseWriter, r *http.Request) {
	h.BoardManagerApp = api.GetBoardManagerApp()

	modifiedBoard := &entity.Board{} // TODO implement
	updatedBoard, err := h.BoardManagerApp.Create(modifiedBoard)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to update the board; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, updatedBoard)
}

func (h BoardManagerAppHandler) deleteBoard(w http.ResponseWriter, r *http.Request) {
	h.BoardManagerApp = api.GetBoardManagerApp()

	var boardId common.Id // TODO implement

	if err := h.BoardManagerApp.Delete(boardId); err != nil {
		respondError(w, http.StatusNotFound, "failed to get the board; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, "Deleted Successfully")
}
