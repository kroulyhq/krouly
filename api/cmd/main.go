package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Data struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func main() {
	r := mux.NewRouter()
	port := ":8080"

	r.HandleFunc("/storage", handleData).Methods("GET")

	fmt.Println("Ready to rock n roll at: http://localhost", port)

	// Add CORS middleware
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}

	http.ListenAndServe(port, corsMiddleware(r))
}

func handleData(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Hello, I'm Krouly!")

	filePath := "../../storage/cryptodata.json"
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(w, "Error reading JSON file: %v", err)
		return
	}

	var jsonData []Data
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Fprintf(w, "Error converting JSON: %v", err)
		return
	}

	fmt.Println(jsonData)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonData)
}
