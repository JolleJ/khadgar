package transaction

import (
	"context"
	"database/sql"
	"jollej/db-scout/internal/domain/transaction"
)

type TransactionRepo struct {
	db *sql.DB
}

func NewTransactionRepo(db *sql.DB) transaction.TransactionRepo {
	return &TransactionRepo{db: db}
}

func (t *TransactionRepo) Create(ctx context.Context, transaction transaction.Transaction) (int, error) {
	query := `INSERT INTO transactions (user_id, portfolio_id, type, amount, currency, transaction_date, created_at, status) 
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := t.db.ExecContext(ctx, query, transaction.UserId, transaction.PortfolioId, transaction.Type,
		transaction.Amount, transaction.Currency, transaction.TransactionDate, transaction.CreatedAt, transaction.Status)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (t *TransactionRepo) ListByPortfolioId(portfolioId int) ([]transaction.Transaction, error) {
	query := `SELECT id, user_id, portfolio_id, type, amount, currency, transaction_date, created_at, status 
	          FROM transactions WHERE portfolio_id = ?`
	rows, err := t.db.Query(query, portfolioId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []transaction.Transaction
	for rows.Next() {
		var t transaction.Transaction
		if err := rows.Scan(&t.Id, &t.UserId, &t.PortfolioId, &t.Type, &t.Amount, &t.Currency,
			&t.TransactionDate, &t.CreatedAt, &t.Status); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}
