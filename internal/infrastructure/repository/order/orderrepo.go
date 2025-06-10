package order

import (
	"context"
	"database/sql"
	"jollej/db-scout/internal/domain/order"
)

type OrderRepo struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) order.OrderRepo {
	return &OrderRepo{db: db}
}

func (o *OrderRepo) Create(ctx context.Context, order order.Order) (int, error) {
	query := `INSERT INTO orders (portfolio_id, instrument_id, side, quantity, price, status) VALUES (?, ?, ?, ?, ?, ?)`
	result, err := o.db.ExecContext(ctx, query, order.PortfolioId, order.InstrumentId, order.Side, order.Quantity, order.Price, order.Status)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
