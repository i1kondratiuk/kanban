package handler

import (
	"net/http"

	"github.com/i1kondratiuk/kanban/application"
)

// BoardManagerAppHandler ...
type BoardManagerAppHandler struct {
	BoardManagerApp application.BoardManagerApp
}

// AddRoutes adds BoardManagerAppHandler routs
func (h BoardManagerAppHandler) AddRoutes() {
	http.HandleFunc("/boards", h.getAllBoards)
}

func (h BoardManagerAppHandler) getAllBoards(w http.ResponseWriter, r *http.Request) {
	h.BoardManagerApp = application.GetBoardManagerApp()

	boards, err := h.BoardManagerApp.GetAllBoardsSortedByNameAsc()

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to get boards; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, boards)
}
