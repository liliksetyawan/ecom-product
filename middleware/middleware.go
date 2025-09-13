package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func HandleRequest(w http.ResponseWriter, r *http.Request, serviceFunc func(*http.Request) (interface{}, error, interface{})) {
	requestID := GenerateRequestID()

	response, err, data := serviceFunc(r)
	if err != nil {
		Error(w, requestID, http.StatusBadRequest, err.Error())
		logRequest(r.URL.Path, data, requestID, "failed", err)
		return
	}

	Success(w, requestID, response)
	logRequest(r.URL.Path, data, requestID, "success", nil)
}

func logRequest(path string, requestBody interface{}, requestID, status string, err error) {
	logData := map[string]interface{}{
		"request_id": requestID,
		"timestamp":  time.Now().Format(time.RFC3339),
		"path":       path,
		"status":     status,
		"request":    requestBody,
		"error":      "",
	}

	if err != nil {
		logData["error"] = err.Error()
	}

	logJSON, _ := json.Marshal(logData)
	log.Println(string(logJSON))
}
