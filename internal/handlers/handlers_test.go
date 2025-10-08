package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetCustomerPoints(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/customers/cust-001/points", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/customers/{customerID}/points", GetCustomerPoints)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Response contains invalid JSON: %v", err)
	}

	if _, exists := response["id"]; !exists {
		t.Error("Response missing 'id' field")
	}
}