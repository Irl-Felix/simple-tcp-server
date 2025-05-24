package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func HealthCheckHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Path hit:", req.URL.Path)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Service is healthy"))
	fmt.Println("Health check endpoint hit")
}

func ClientinfoHandler(w http.ResponseWriter, req *http.Request) {
	//Request Metadata
	clientIP := getClientIP(req)
	method := req.Method
	path := req.URL.Path
	query := req.URL.Query()
	headers := make(map[string]string)
	for name, values := range req.Header {
		headers[name] = strings.Join(values, ", ")
	}
	timestamp := time.Now().Format(time.RFC3339)

	// Response Payload
	response := map[string]interface{}{
		"timestamp":    timestamp,
		"ip":           clientIP,
		"method":       method,
		"path":         path,
		"query_params": query,
		"headers":      headers,
	}

	//Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
