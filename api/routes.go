package api

import (
	"net/http"

	"github.com/talathiom/goTemplate/internal/handlers"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	// Define API routes
	router.HandleFunc("/api/resource", handlers.HandleResource)

	return router
}
