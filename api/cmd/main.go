package main

import (
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Define API endpoints and handlers
	r.HandleFunc("/api/v1/endpoint", handleRequest).Methods("GET")

	// Start the HTTP server

	http.ListenAndServe(":8080", r)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Write the response message
	fmt.Fprintf(w, "Hello, I'm Krouly!")
}
