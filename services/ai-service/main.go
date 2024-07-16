package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PredictionRequest struct {
	Data []float64 `json:"data"`
}

type PredictionResponse struct {
	Predictions []float64 `json:"predictions"`
}

type CryptoData struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

// ReadDataFromFile reads the data from the specified file and parses it into a slice of float64
func ReadDataFromFile(filePath string) ([]float64, error) {
	var cryptoData []CryptoData

	// Read the file
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	// Parse the JSON data
	err = json.Unmarshal(fileContent, &cryptoData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	// Extract prices and convert them to float64
	var data []float64
	for _, item := range cryptoData {
		var price float64
		fmt.Sscanf(item.Price, "%f", &price)
		data = append(data, price)
	}

	return data, nil
}

func RunAIInference(filePath string) ([]float64, error) {
	// Read data from file
	data, err := ReadDataFromFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from file: %v", err)
	}

	// Marshal the data into a JSON request
	requestBody, err := json.Marshal(PredictionRequest{Data: data})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	// Send the request to the AI model
	resp, err := http.Post("http://localhost:8501/v1/models/your_model:predict", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to send request to AI model: %v", err)
	}
	defer resp.Body.Close()

	// Decode the response
	var predictionResponse PredictionResponse
	if err := json.NewDecoder(resp.Body).Decode(&predictionResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return predictionResponse.Predictions, nil
}
