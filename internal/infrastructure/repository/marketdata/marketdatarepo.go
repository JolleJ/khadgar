package marketdata

import (
	"context"
	"database/sql"
	"jollej/db-scout/internal/domain/marketdata"
)

type MarketDataRepo struct {
	db *sql.DB
}

func NewMarketDataRepo(db *sql.DB) marketdata.MarketDataRepo {
	return &MarketDataRepo{db: db}
}

func (m *MarketDataRepo) ListBySymbol(ctx context.Context, ticker string) ([]marketdata.MarketDataPoint, error) {
	var res []marketdata.MarketDataPoint
	query := `SELECT * FROM market_data WHERE ticker = ?`

	rows, err := m.db.Query(query, ticker)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var data marketdata.MarketDataPoint
		if err := rows.Scan(&data.Ticker, &data.Date, &data.Open, &data.High, &data.Low, &data.Close, &data.Volume, &data.Dividends, &data.StockSplits); err != nil {
			return nil, err
		}
		res = append(res, data)
	}

	return res, nil
}

func (m *MarketDataRepo) AddMarketData(ctx context.Context, ticker string, marketData []marketdata.MarketDataPoint) error {
	if len(marketData) == 0 {
		return nil // No data to INSERT
	}

	for _, data := range marketData {
		query := `INSERT INTO market_data (ticker, date, open, high, low, close, volume, dividends, stock_splits) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
		_, err := m.db.Exec(query, data.Ticker, data.Date, data.Open, data.High, data.Low, data.Close, data.Volume, data.Dividends, data.StockSplits)
		if err != nil {
			return err
		}
	}
	return nil
}
