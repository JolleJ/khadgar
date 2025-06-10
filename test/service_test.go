package test

import (
	"github.com/shopspring/decimal"
	tradeApp "jollej/db-scout/internal/application/trade"
	"jollej/db-scout/internal/domain/trade"
	"jollej/db-scout/internal/infrastructure"
	tradeRepo "jollej/db-scout/internal/infrastructure/repository/trade"
	"testing"
)

func TestCreateTradeService(t *testing.T) {
	db, err := infrastructure.InitDb()
	if err != nil {
		t.Fatalf("Failed to initialize the database: %v", err)
	}

	repo := tradeRepo.NewTradeRepo(db)
	service := tradeApp.NewTradeService(repo)

	trade := trade.Trade{
		OrderId:    1,
		Quantity:   decimal.NewFromFloat(1.0),
		Price:      decimal.NewFromFloat(0.01),
		ExecutedAt: "2023-10-01T12:00:00Z",
		Status:     "completed",
	}

	id, err := service.CreateTrade(nil, trade)
	if err != nil {
		t.Errorf("Failed to create trade: %v", err)
	}
	if id <= 0 {
		t.Errorf("Expected a valid trade ID, got %d", id)
	}

	if err != nil {
		t.Errorf("Failed to create trade: %v", err)
	}
}
