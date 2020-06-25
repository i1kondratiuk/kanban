package handler

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/i1kondratiuk/kanban/application"
)

// BoardManagerAppHandler ...
type BoardManagerAppHandler struct {
	BoardManagerApp application.BoardManagerApp
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
	// h.BoardManagerApp = application.GetBoardManagerApp()
	//
	// boards, err := h.BoardManagerApp.GetAllBoardsSortedByNameAsc()
	//
	// if err != nil {
	// 	respondError(w, http.StatusNotFound, "failed to get boards; "+err.Error())
	// 	return
	// }
	//
	// respondJSON(w, http.StatusOK, boards)
}

func (h BoardManagerAppHandler) CreateBoard(w http.ResponseWriter, r *http.Request) {
}

func (h BoardManagerAppHandler) getBoard(w http.ResponseWriter, r *http.Request) {
}

func (h BoardManagerAppHandler) updateBoard(w http.ResponseWriter, r *http.Request) {
}

func (h BoardManagerAppHandler) deleteBoard(w http.ResponseWriter, r *http.Request) {
}
