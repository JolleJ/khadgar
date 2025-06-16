package trade

import (
	"context"
	"database/sql"
	"jollej/db-scout/internal/domain/trade"
	"jollej/db-scout/lib/prettylog"
)

type TradeRepo struct {
	db *sql.DB
}

func NewTradeRepo(db *sql.DB) trade.TradeRepo {
	return &TradeRepo{db: db}
}

func (t *TradeRepo) CreateTrade(ctx context.Context, trade trade.Trade) (int, error) {
	log := prettylog.NewPrettyLog()
	log.Infof("Creating trade: %+v", trade)
	query := `INSERT INTO trades (order_id, quantity, price, executed_at, status) VALUES (?, ?, ?, ?, ?)`
	result, err := t.db.Exec(query, trade.OrderId, trade.Quantity, trade.Price, trade.ExecutedAt, trade.Status)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (t *TradeRepo) ListTradesOnPortfolio(ctx context.Context, portfolioId int) ([]trade.Trade, error) {
	query := `SELECT id, order_id, quantity, price, executed_at, status FROM trades WHERE order_id IN (SELECT id FROM orders WHERE portfolio_id = ?)`
	rows, err := t.db.QueryContext(ctx, query, portfolioId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trades []trade.Trade
	for rows.Next() {
		var trade trade.Trade
		if err := rows.Scan(&trade.Id, &trade.OrderId, &trade.Quantity, &trade.Price, &trade.ExecutedAt, &trade.Status); err != nil {
			return nil, err
		}
		trades = append(trades, trade)
	}

	return trades, nil
}
