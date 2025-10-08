package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"loyalty-points-service/internal/handlers"

	"github.com/gorilla/mux"
)

func TestGetCustomerPoints(t *testing.T) {
	// Crée une requête HTTP
	req, err := http.NewRequest("GET", "/api/v1/customers/cust-001/points", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Crée un ResponseRecorder pour capturer la réponse
	rr := httptest.NewRecorder()

	// Crée un router et appelle le handler
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/customers/{customerID}/points", handlers.GetCustomerPoints)
	router.ServeHTTP(rr, req)

	// Vérifie le status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Vérifie le Content-Type
	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Handler returned wrong content type: got %v want %v", contentType, expectedContentType)
	}

	// Vérifie que la réponse contient des données JSON valides
	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Response contains invalid JSON: %v", err)
	}

	// Vérifie les champs attendus
	if _, exists := response["id"]; !exists {
		t.Error("Response missing 'id' field")
	}
	if _, exists := response["points"]; !exists {
		t.Error("Response missing 'points' field")
	}
}

func TestAddPoints(t *testing.T) {
	// Données de test
	requestBody := map[string]interface{}{
		"points": 50,
		"reason": "Test bonus",
	}
	jsonData, _ := json.Marshal(requestBody)

	// Crée une requête POST
	req, err := http.NewRequest("POST", "/api/v1/customers/cust-001/points/add", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Crée un ResponseRecorder
	rr := httptest.NewRecorder()

	// Appelle le handler
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/customers/{customerID}/points/add", handlers.AddPoints)
	router.ServeHTTP(rr, req)

	// Vérifie le status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Vérifie la réponse
	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Response contains invalid JSON: %v", err)
	}

	if message, exists := response["message"]; !exists || message != "Points added successfully" {
		t.Error("Add points operation failed")
	}
}

func TestHealthCheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Simule le handler health check
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "healthy",
			"service": "loyalty-points",
			"version": "1.0.0",
		})
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Health check returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
