package handler

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/i1kondratiuk/kanban/application/api"
	"github.com/i1kondratiuk/kanban/domain/entity"
	"github.com/i1kondratiuk/kanban/domain/entity/common"
	"github.com/i1kondratiuk/kanban/domain/value"
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
	h.CommentManagerApp = api.GetCommentManagerApp()

	var taskId common.Id // TODO Implement
	storedComments, err := h.CommentManagerApp.GetAllParentCommentsGroupedByCreatedDateTime(taskId)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to get comments; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, storedComments)
}

func (h CommentManagerAppHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	h.CommentManagerApp = api.GetCommentManagerApp()

	newComment := &entity.Comment{} // TODO Implement
	newCommentStored, err := h.CommentManagerApp.Create(newComment)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to get comments; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, newCommentStored)
}

func (h CommentManagerAppHandler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	h.CommentManagerApp = api.GetCommentManagerApp()

	var commentId common.Id     // TODO Implement
	var bodyText value.BodyText // TODO Implement
	newCommentStored, err := h.CommentManagerApp.UpdateBodyText(commentId, bodyText)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to get comments; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, newCommentStored)
}

func (h CommentManagerAppHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	h.CommentManagerApp = api.GetCommentManagerApp()

	var storedCommentId common.Id // TODO Implement

	if err := h.CommentManagerApp.Delete(storedCommentId); err != nil {
		respondError(w, http.StatusNotFound, "failed to get comments; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, "the comment was deleted successfully")
}
