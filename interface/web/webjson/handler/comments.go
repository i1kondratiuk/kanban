package handler

import (
	"net/http"

	"github.com/i1kondratiuk/kanban/application"
)

// CommentManagerAppHandler handler
type CommentManagerAppHandler struct {
	CommentManagerApp application.CommentManagerApp
}

// AddRoutes adds CommentManagerAppHandler routs
func (h CommentManagerAppHandler) AddRoutes() {
	http.HandleFunc("/comments", h.GetAllComments)
}

func (h CommentManagerAppHandler) GetAllComments(w http.ResponseWriter, r *http.Request) {

}

