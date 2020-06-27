package handler

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/i1kondratiuk/kanban/application/api"
)

// ColumnManagerAppHandler handler
type ColumnManagerAppHandler struct {
	ColumnManagerApp api.ColumnManagerApp
}

// AddRoutes adds ColumnManagerAppHandler routs
func (h ColumnManagerAppHandler) AddRoutes(r *mux.Router) {
	r.HandleFunc("/boards/{boardId}/columns", h.GetAllBoardColumns).Methods("GET")
	r.HandleFunc("/boards/{boardId}/columns", h.CreateBoardColumn).Methods("POST")

	r.HandleFunc("/boards/{boardId}/columns/{columnId}", h.GetBoardColumn).Methods("GET")
	r.HandleFunc("/boards/{boardId}/columns/{columnId}", h.UpdateBoardColumn).Methods("PUT")
	r.HandleFunc("/boards/{boardId}/columns/{columnId}", h.DeleteBoardColumn).Methods("DELETE")

	r.HandleFunc("/boards/{boardId}/columns/{columnId}/name", h.UpdateBoardColumn).Methods("PUT")
	r.HandleFunc("/boards/{boardId}/columns/{columnId}/position", h.UpdateBoardColumn).Methods("PUT")
}

func (h ColumnManagerAppHandler) GetAllBoardColumns(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">> GetAllBoardColumns")
	respondJSON(w, http.StatusOK, "GetAllBoardColumns")
}

func (h ColumnManagerAppHandler) CreateBoardColumn(w http.ResponseWriter, r *http.Request) {
}

func (h ColumnManagerAppHandler) GetBoardColumn(w http.ResponseWriter, r *http.Request) {
}

func (h ColumnManagerAppHandler) UpdateBoardColumn(w http.ResponseWriter, r *http.Request) {
}

func (h ColumnManagerAppHandler) DeleteBoardColumn(w http.ResponseWriter, r *http.Request) {
}
