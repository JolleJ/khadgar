package order

import (
	"context"
	"database/sql"
	"jollej/db-scout/internal/domain/order"
	"sync"
)

type OrderRepo struct {
	// dbMutex is used to ensure that only one goroutine can access the database at a time. This is necessary because SQLite does not support concurrent writes well.
	dbMutex sync.Mutex
	db      *sql.DB
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

func (o *OrderRepo) ListByInstrument(ctx context.Context, ticker string) ([]order.Order, error) {
	o.dbMutex.Lock()
	defer o.dbMutex.Unlock()
	query := `SELECT id, portfolio_id, instrument_id, side, quantity, price, status FROM orders WHERE instrument_id = (SELECT id FROM instruments WHERE ticker = ?)`
	rows, err := o.db.QueryContext(ctx, query, ticker)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []order.Order
	for rows.Next() {
		var ord order.Order
		if err := rows.Scan(&ord.Id, &ord.PortfolioId, &ord.InstrumentId, &ord.Side, &ord.Quantity, &ord.Price, &ord.Status); err != nil {
			return nil, err
		}
		orders = append(orders, ord)
	}

	return orders, nil
}
