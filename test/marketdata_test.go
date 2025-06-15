package test

import (
	marketDataApp "jollej/db-scout/internal/application/marketdata"
	"jollej/db-scout/internal/infrastructure"
	"jollej/db-scout/internal/infrastructure/repository/instrument"
	"jollej/db-scout/internal/infrastructure/repository/marketdata"
	"testing"
)

func TestMarketData(t *testing.T) {

	db, err := infrastructure.InitDb()
	if err != nil {
		t.Fatalf("Failed to initialize the database: %v", err)
	}

	instrumentRepo := instrument.NewInstrumentrepo(db)
	marketdataRepo := marketdata.NewMarketDataRepo(db)
	marketdataService := marketDataApp.NewMarketDataService(instrumentRepo, marketdataRepo)

	err = marketdataService.Load([]string{"AAPL", "GOOGL", "MSFT"})
	if err != nil {
		t.Fatalf("Failed to load market data: %v", err)
	}
}
