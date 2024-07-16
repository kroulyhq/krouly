package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type PredictionRequest struct {
	Data []float64 `json:"data"`
}

type PredictionResponse struct {
	Predictions []float64 `json:"predictions"`
}

func RunAIInference(data []float64) ([]float64, error) {
	requestBody, err := json.Marshal(PredictionRequest{Data: data})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	resp, err := http.Post("http://localhost:8501/v1/models/your_model:predict", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to send request to AI model: %v", err)
	}
	defer resp.Body.Close()

	var predictionResponse PredictionResponse
	if err := json.NewDecoder(resp.Body).Decode(&predictionResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return predictionResponse.Predictions, nil
}
