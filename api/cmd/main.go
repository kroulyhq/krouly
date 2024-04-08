package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

type Data struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type DataPlaybook struct {
	Name       string `json:"name"`
	Connector  string `json:"connector"`
	Parameters struct {
		URL string `json:"url"`
	} `json:"parameters"`
}

type Task struct {
	Name      string `yaml:"name"`
	Connector string `yaml:"connector"`
	Params    struct {
		URL string `yaml:"url"`
	} `yaml:"parameters"`
}

type Playbook struct {
	Name  string `yaml:"playbook"`
	Tasks []Task `yaml:"tasks"`
}

func main() {
	r := mux.NewRouter()
	port := ":8080"

	r.HandleFunc("/storage", handleData).Methods("GET")

	r.HandleFunc("/sources", handleSources).Methods("GET")

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

/* DRY: We are using this also in connetors */
func loadPlaybook(filename string) (*Playbook, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening playbook file: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	var playbook Playbook
	err = decoder.Decode(&playbook)
	if err != nil {
		return nil, fmt.Errorf("error parsing playbook YAML: %v", err)
	}

	return &playbook, nil
}

func handleSources(w http.ResponseWriter, r *http.Request) {
	var jsonData []DataPlaybook

	playbookFile := "../../playbooks/krouly.sample.yaml"
	playbook, err := loadPlaybook(playbookFile)
	if err != nil {
		fmt.Println("Error loading playbook:", err)
		http.Error(w, "Error loading playbook", http.StatusInternalServerError)
		return
	}

	for _, task := range playbook.Tasks {
		taskData := DataPlaybook{
			Name:      task.Name,
			Connector: task.Connector,
			Parameters: struct {
				URL string `json:"url"`
			}{URL: task.Params.URL},
		}
		jsonData = append(jsonData, taskData)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonData)
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

	// cfg := elasticsearch.Config{
	// 	Addresses: []string{"https://192.168.1.16:5601"},
	// 	Username:  "elastic",
	// 	Password:  "upU+yiqQimy7c-97-3aO",
	// }
	// es, err := elasticsearch.NewClient(cfg)
	// if err != nil {
	// 	log.Fatalf("Error creating the client: %s", err)
	// }

	// var buf strings.Builder
	// for _, item := range jsonData {
	// 	doc := map[string]interface{}{
	// 		"symbol": item.Symbol,
	// 		"price":  item.Price,
	// 	}
	// 	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
	// 		log.Printf("Error encoding document: %s", err)
	// 		continue
	// 	}
	// 	req := esapi.IndexRequest{
	// 		Index:      "collections",
	// 		DocumentID: "", // Auto-generated
	// 		Body:       strings.NewReader(buf.String()),
	// 		Refresh:    "true",
	// 	}
	// 	res, err := req.Do(context.Background(), es)
	// 	if err != nil {
	// 		log.Printf("Error indexing document: %s", err)
	// 		continue
	// 	}
	// 	defer res.Body.Close()
	// 	if res.IsError() {
	// 		log.Printf("Error indexing document: %s", res.String())
	// 		continue
	// 	}
	// }

	// fmt.Println("Data indexed successfully")

	json.NewEncoder(w).Encode(jsonData)
}
