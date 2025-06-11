package transaction

import (
	"context"
	"jollej/db-scout/internal/domain/transaction"
)

type TransactionService struct {
	transactionRepo transaction.TransactionRepo
}

func NewTransactionService(transactionRepo transaction.TransactionRepo) *TransactionService {
	return &TransactionService{transactionRepo: transactionRepo}
}

func (t *TransactionService) Create(ctx context.Context, transaction transaction.Transaction) (int, error) {
	return t.transactionRepo.Create(ctx, transaction)
}
