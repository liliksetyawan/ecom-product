package middleware

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type BaseResponse struct {
	Timestamp string      `json:"timestamp"`
	RequestID string      `json:"request_id"`
	Data      interface{} `json:"data"`
}

func Success(w http.ResponseWriter, requestID string, data interface{}) {
	resp := BaseResponse{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		RequestID: requestID,
		Data:      data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func Error(w http.ResponseWriter, requestID string, statusCode int, error string) {
	resp := BaseResponse{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		RequestID: requestID,
		Data: map[string]string{
			"error": error,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}

func GenerateRequestID() string {
	return uuid.New().String()
}
