package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	boardIdAnchor   = "boardId"
	columnIdAnchor  = "columnId"
	taskIdAnchor    = "taskId"
	commentIdAnchor = "commentId"
)

// Run starts server
func Run(port int) {
	log.Printf("Server running at http://localhost:%d/", port)

	r := mux.NewRouter()

	r.HandleFunc("/", getWelcomePage).Methods("GET")

	new(BoardManagerAppHandler).AddRoutes(r)
	new(ColumnManagerAppHandler).AddRoutes(r)
	new(TaskManagerAppHandler).AddRoutes(r)
	new(CommentManagerAppHandler).AddRoutes(r)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func getWelcomePage(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, "Welcome to Kanban")
}
