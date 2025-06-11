package transaction

import "context"

type TransactionRepo interface {
	Create(ctx context.Context, transaction Transaction) (int, error)
	ListByPortfolioId(portfolioId int) ([]Transaction, error)
}
