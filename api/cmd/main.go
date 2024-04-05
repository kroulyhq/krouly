package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
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

	cfg := elasticsearch.Config{
		Addresses: []string{"https://192.168.1.16:5601"},
		Username:  "elastic",
		Password:  "upU+yiqQimy7c-97-3aO",
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	var buf strings.Builder
	for _, item := range jsonData {
		doc := map[string]interface{}{
			"symbol": item.Symbol,
			"price":  item.Price,
		}
		if err := json.NewEncoder(&buf).Encode(doc); err != nil {
			log.Printf("Error encoding document: %s", err)
			continue
		}
		req := esapi.IndexRequest{
			Index:      "collections",
			DocumentID: "", // Auto-generated
			Body:       strings.NewReader(buf.String()),
			Refresh:    "true",
		}
		res, err := req.Do(context.Background(), es)
		if err != nil {
			log.Printf("Error indexing document: %s", err)
			continue
		}
		defer res.Body.Close()
		if res.IsError() {
			log.Printf("Error indexing document: %s", res.String())
			continue
		}
	}

	fmt.Println("Data indexed successfully")

	json.NewEncoder(w).Encode(jsonData)
}
