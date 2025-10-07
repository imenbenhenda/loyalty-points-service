package main

import (
	"encoding/json"
	"log"
	"net/http"

	"loyalty-points-service/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Routes API
	api := r.PathPrefix("/api/v1").Subrouter()

	// Points endpoints
	api.HandleFunc("/customers/{customerID}/points", handlers.GetCustomerPoints).Methods("GET")
	api.HandleFunc("/customers/{customerID}/points/add", handlers.AddPoints).Methods("POST")
	api.HandleFunc("/customers/{customerID}/points/redeem", handlers.RedeemPoints).Methods("POST")

	// Health check
	api.HandleFunc("/health", healthCheck).Methods("GET")

	log.Println("🚀 Loyalty Points Service starting on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}

// Health check endpoint
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"status":  "healthy",
		"service": "loyalty-points",
		"version": "1.0.0",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Encoding error", http.StatusInternalServerError)
	}
}
