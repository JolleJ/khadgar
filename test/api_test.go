package test

import (
	"bytes"
	"encoding/json"
	"jollej/db-scout/internal/infrastructure"
	"jollej/db-scout/internal/interfaces/api"
	"jollej/db-scout/internal/interfaces/api/dto/order"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shopspring/decimal"
)

func TestGetUsersApiEndpoint(t *testing.T) {
	// Initialize the database
	db, err := infrastructure.InitDb()
	if err != nil {
		t.Fatalf("Failed to initialize the database: %v", err)
	}

	mux := api.NewMux(db)

	req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
	res := httptest.NewRecorder()

	mux.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
	}

	resBody := res.Body.String()

	if len(resBody) == 0 {
		t.Error("Expected to have some content in the response")
	}

}
func TestGetPortfolioApiEndpoint(t *testing.T) {
	// Initialize the database
	db, err := infrastructure.InitDb()
	if err != nil {
		t.Fatalf("Failed to initialize the database: %v", err)
	}

	mux := api.NewMux(db)

	req := httptest.NewRequest(http.MethodGet, "/api/portfolio/1", nil)
	res := httptest.NewRecorder()

	mux.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
	}

	resBody := res.Body.String()

	if len(resBody) == 0 {
		t.Error("Expected to have some content in the response")
	}

}
func TestGetListInstruments(t *testing.T) {
	// Initialize the database
	db, err := infrastructure.InitDb()
	if err != nil {
		t.Fatalf("Failed to initialize the database: %v", err)
	}

	mux := api.NewMux(db)

	req := httptest.NewRequest(http.MethodGet, "/api/instruments", nil)
	res := httptest.NewRecorder()

	mux.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
	}

	resBody := res.Body.String()

	if len(resBody) == 0 {
		t.Error("Expected to have some content in the response")
	}

}

func TestCreateOrder(t *testing.T) {

	// Initialize the database
	db, err := infrastructure.InitDb()
	if err != nil {
		t.Fatalf("Failed to initialize the database: %v", err)
	}

	orderDetails := order.OrderDetails{
		InstrumentId: 1,
		Side:         "buy",
		Quantity:     decimal.NewFromFloat(10.0),
		Price:        decimal.NewFromFloat(100.0),
		OrderDate:    "2023-10-01",
		Status:       "pending",
	}

	orderDetailsJson, err := json.Marshal(orderDetails)
	if err != nil {
		t.Fatalf("Failed to marshal order details: %v", err)
	}

	mux := api.NewMux(db)
	req := httptest.NewRequest(http.MethodPost, "/api/portfolio/1/orders", bytes.NewReader(orderDetailsJson))
	res := httptest.NewRecorder()

	mux.ServeHTTP(res, req)

	if res.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, res.Code)
	}

	resBody := res.Body.String()

	if len(resBody) == 0 {
		t.Error("Expected to have some content in the response")
	}

}

func TestGetBalance(t *testing.T) {

	// Initialize the database
	db, err := infrastructure.InitDb()
	if err != nil {
		t.Fatalf("Failed to initialize the database: %v", err)
	}

	mux := api.NewMux(db)
	req := httptest.NewRequest(http.MethodGet, "/api/portfolio/1/balance", nil)
	res := httptest.NewRecorder()

	mux.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
	}

	resBody := res.Body.String()

	if len(resBody) == 0 {
		t.Error("Expected to have some content in the response")
	}

}
