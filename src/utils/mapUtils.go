package utils

import (
	"encoding/json"
	"net/http"
)

// MapResponse represents a unified response structure for map services
type MapResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ParseMapResponse parses the HTTP response from a map service
func ParseMapResponse(resp *http.Response) (*MapResponse, error) {
	var mapResponse MapResponse
	if err := json.NewDecoder(resp.Body).Decode(&mapResponse); err != nil {
		return nil, err
	}
	return &mapResponse, nil
}

// HandleErrorResponse creates a standardized error response
func HandleErrorResponse(err error) *MapResponse {
	return &MapResponse{
		Status:  "error",
		Message: err.Error(),
		Data:    nil,
	}
}