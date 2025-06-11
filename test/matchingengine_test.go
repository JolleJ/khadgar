package test

import (
	instrumentApplication "jollej/db-scout/internal/application/instrument"
	"jollej/db-scout/internal/application/matchingengine"
	orderApp "jollej/db-scout/internal/application/order"
	"jollej/db-scout/internal/infrastructure"
	"jollej/db-scout/internal/infrastructure/repository/instrument"
	"jollej/db-scout/internal/infrastructure/repository/order"
	"testing"
	"time"
)

func TestMatchingEngineInit(t *testing.T) {
	db, err := infrastructure.InitDb()
	if err != nil {
		t.Fatalf("Failed to initialize the database: %v", err)
	}
	instrumentRepo := instrument.NewInstrumentrepo(db)
	instrumentService := instrumentApplication.NewInstrumentService(instrumentRepo)
	orderRepo := order.NewOrderRepo(db)
	orderService := orderApp.NewOrderService(orderRepo)
	engine := matchingengine.NewMatchingEngine(instrumentService, orderService)

	time.Sleep(4000 * time.Millisecond)
	if engine == nil {
		t.Fatal("Expected MatchingEngine to be initialized, got nil")
	}
}
