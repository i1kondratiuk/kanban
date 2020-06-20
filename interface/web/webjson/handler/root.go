package handler

import (
	"fmt"
	"log"
	"net/http"
)

// Run starts server
func Run(port int) {
	log.Printf("Server running at http://localhost:%d/", port)

	new(BoardManagerAppHandler).AddRoutes()
	new(ColumnManagerAppHandler).AddRoutes()
	new(TaskManagerAppHandler).AddRoutes()
	new(CommentManagerAppHandler).AddRoutes()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
