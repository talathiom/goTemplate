package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/talathiom/goTemplate/api"
)

func main() {
	// Initialize API routes
	router := api.NewRouter()

	// Start the server
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
