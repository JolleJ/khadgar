package test

import (
	instrumentApplication "jollej/db-scout/internal/application/instrument"
	"jollej/db-scout/internal/application/matchingengine"
	orderApp "jollej/db-scout/internal/application/order"
	tradeApp "jollej/db-scout/internal/application/trade"
	orderDomain "jollej/db-scout/internal/domain/order"
	"jollej/db-scout/internal/infrastructure"
	"jollej/db-scout/internal/infrastructure/repository/instrument"
	"jollej/db-scout/internal/infrastructure/repository/order"
	"jollej/db-scout/internal/infrastructure/repository/trade"
	"testing"
	"time"

	"github.com/shopspring/decimal"
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
	tradeRepo := trade.NewTradeRepo(db)
	tradeService := tradeApp.NewTradeService(tradeRepo)
	engine := matchingengine.NewMatchingEngine(instrumentService, orderService, tradeService)

	// time.Sleep(4000 * time.Millisecond)
	if engine == nil {
		t.Fatal("Expected MatchingEngine to be initialized, got nil")
	}
}

func TestMatchingEngineAddBuyOrder(t *testing.T) {

	db, err := infrastructure.InitDb()
	if err != nil {
		t.Fatalf("Failed to initialize the database: %v", err)
	}
	instrumentRepo := instrument.NewInstrumentrepo(db)
	instrumentService := instrumentApplication.NewInstrumentService(instrumentRepo)
	orderRepo := order.NewOrderRepo(db)
	orderService := orderApp.NewOrderService(orderRepo)
	tradeRepo := trade.NewTradeRepo(db)
	tradeService := tradeApp.NewTradeService(tradeRepo)
	engine := matchingengine.NewMatchingEngine(instrumentService, orderService, tradeService)

	order := orderDomain.Order{
		Id:           999,
		PortfolioId:  1,
		InstrumentId: 1,
		Side:         "buy",
		Price:        decimal.NewFromFloat(50000.0),
		Quantity:     decimal.NewFromFloat(1.0),
		Status:       "pending",
	}
	time.Sleep(1000 * time.Millisecond) // Wait for the engine to initialize
	engine.PlaceOrder("AAPL", order)

}
