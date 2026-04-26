package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Response defines the standard JSON response structure
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// HealthCheck defines the health check response
type HealthCheck struct {
	Status string `json:"status"`
	Uptime string `json:"uptime"`
}

// EchoRequest defines the echo endpoint request
type EchoRequest struct {
	Message string `json:"message"`
}

func main() {
	// Register HTTP handlers
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/echo", echoHandler)

	// Start the web server
	port := ":8080"
	log.Printf("Starting web service on http://localhost%s\n", port)
	log.Printf("Available endpoints:")
	log.Printf("  GET  /hello")
	log.Printf("  GET  /health")
	log.Printf("  POST /echo")

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server error: %v\n", err)
	}
}

// helloHandler responds with a greeting
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := Response{
		Status:  "success",
		Message: "Hello from Go web service!",
		Data: map[string]string{
			"version": "1.0",
			"service": "Basic Go Web Service",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// healthHandler checks the service health
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	health := HealthCheck{
		Status: "healthy",
		Uptime: "running",
	}

	response := Response{
		Status:  "success",
		Message: "Service is healthy",
		Data:    health,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// echoHandler echoes back a message sent by the client
func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req EchoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response := Response{
			Status:  "error",
			Message: "Invalid request body",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := Response{
		Status:  "success",
		Message: "Echo response",
		Data: map[string]string{
			"echo": req.Message,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
