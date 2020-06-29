package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

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
func (h CommentManagerAppHandler) AddRoutes(r *mux.Router) { // TODO get rid of the redundant path prefix for the subresource
	r.HandleFunc("/boards/{"+boardIdAnchor+"}/columns/{"+columnIdAnchor+"}/tasks/{"+taskIdAnchor+"}/comments", h.GetAllComments).Methods("GET")
	r.HandleFunc("/boards/{"+boardIdAnchor+"}/columns/{"+columnIdAnchor+"}/tasks/{"+taskIdAnchor+"}/comments", h.CreateComment).Methods("POST")

	r.HandleFunc("/boards/{"+boardIdAnchor+"}/columns/{"+columnIdAnchor+"}/tasks/{"+taskIdAnchor+"}/comments/{"+commentIdAnchor+"}", h.UpdateComment).Methods("PUT")
	r.HandleFunc("/boards/{"+boardIdAnchor+"}/columns/{"+columnIdAnchor+"}/tasks/{"+taskIdAnchor+"}/comments/{"+commentIdAnchor+"}", h.DeleteComment).Methods("DELETE")
}

func (h CommentManagerAppHandler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	h.CommentManagerApp = api.GetCommentManagerApp()

	params := mux.Vars(r)
	taskIdInt64, err := strconv.ParseInt(params[taskIdAnchor], 10, 64)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to get comments; "+err.Error())
		return
	}

	taskId := common.Id(taskIdInt64)

	storedComments, err := h.CommentManagerApp.GetAllParentCommentsGroupedByCreatedDateTime(taskId)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to get comments; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, storedComments)
}

func (h CommentManagerAppHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	h.CommentManagerApp = api.GetCommentManagerApp()

	var newComment entity.Comment

	if err := json.NewDecoder(r.Body).Decode(&newComment); err != nil {
		respondError(w, http.StatusNotFound, "failed to create the column; "+err.Error())
		return
	}

	newCommentStored, err := h.CommentManagerApp.Create(&newComment)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to get comments; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, newCommentStored)
}

func (h CommentManagerAppHandler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	h.CommentManagerApp = api.GetCommentManagerApp()

	params := mux.Vars(r)
	commentIdInt64, err := strconv.ParseInt(params[commentIdAnchor], 10, 64)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to update the column; "+err.Error())
		return
	}

	var bodyText value.BodyText

	if err := json.NewDecoder(r.Body).Decode(&bodyText); err != nil {
		respondError(w, http.StatusNotFound, "failed to update the column; "+err.Error())
		return
	}

	commentId := common.Id(commentIdInt64)

	newCommentStored, err := h.CommentManagerApp.UpdateBodyText(commentId, bodyText)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to update the column; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, newCommentStored)
}

func (h CommentManagerAppHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	h.CommentManagerApp = api.GetCommentManagerApp()

	params := mux.Vars(r)
	commentId, err := strconv.ParseInt(params[commentIdAnchor], 10, 64)

	if err != nil {
		respondError(w, http.StatusNotFound, "failed to delete the task; "+err.Error())
		return
	}

	if err := h.CommentManagerApp.Delete(common.Id(commentId)); err != nil {
		respondError(w, http.StatusNotFound, "failed to delete the task; "+err.Error())
		return
	}

	respondJSON(w, http.StatusOK, "the comment was deleted successfully")
}
