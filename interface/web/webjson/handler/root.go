package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Run starts server
func Run(port int) {
	log.Printf("Server running at http://localhost:%d/", port)

	r := mux.NewRouter()

	new(BoardManagerAppHandler).AddRoutes(r)
	new(ColumnManagerAppHandler).AddRoutes(r)
	new(TaskManagerAppHandler).AddRoutes(r)
	new(CommentManagerAppHandler).AddRoutes(r)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
