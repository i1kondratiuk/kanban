package handler

import (
	"net/http"

	"github.com/i1kondratiuk/kanban/application"
)

// ColumnManagerAppHandler handler
type ColumnManagerAppHandler struct {
	ColumnManagerApp application.ColumnManagerApp
}

// AddRoutes adds ColumnManagerAppHandler routs
func (h ColumnManagerAppHandler) AddRoutes() {
	http.HandleFunc("/columns", h.GetBoardColumns)
}

func (h ColumnManagerAppHandler) GetBoardColumns(w http.ResponseWriter, r *http.Request) {

}
