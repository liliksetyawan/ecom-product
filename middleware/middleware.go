package middleware

import (
	"context"
	"ecom-product/config"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

type contextKey string

const ContextKeyUser contextKey = "user"

type ContextData struct {
	UserID string
	Token  string
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		userID := r.Header.Get("X-UserID")

		if authHeader == "" || userID == "" {
			http.Error(w, "missing Authorization or X-UserID header", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token != config.AppConfig.FixToken {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ContextKeyUser, &ContextData{
			UserID: userID,
			Token:  token,
		})

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

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
