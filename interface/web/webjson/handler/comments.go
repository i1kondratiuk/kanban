package handler

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/i1kondratiuk/kanban/application/api"
)

// CommentManagerAppHandler handler
type CommentManagerAppHandler struct {
	CommentManagerApp api.CommentManagerApp
}

// AddRoutes adds CommentManagerAppHandler routs
func (h CommentManagerAppHandler) AddRoutes(r *mux.Router) {
	r.HandleFunc("/boards/{boardId}/columns/{columnId}/tasks/{taskId}/comments", h.GetAllComments).Methods("GET")
	r.HandleFunc("/boards/{boardId}/columns/{columnId}/tasks/{taskId}/comments", h.CreateComment).Methods("POST")

	r.HandleFunc("/boards/{boardId}/columns/{columnId}/tasks/{taskId}/comments/{taskId}", h.UpdateComment).Methods("PUT")
	r.HandleFunc("/boards/{boardId}/columns/{columnId}/tasks/{taskId}/comments/{taskId}", h.DeleteComment).Methods("DELETE")
}

func (h CommentManagerAppHandler) GetAllComments(w http.ResponseWriter, r *http.Request) {
}

func (h CommentManagerAppHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
}

func (h CommentManagerAppHandler) UpdateComment(w http.ResponseWriter, r *http.Request) {
}

func (h CommentManagerAppHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
}
