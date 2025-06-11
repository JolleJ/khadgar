package balance

import "jollej/db-scout/internal/domain/transaction"

type BalanceService struct {
	transactionRepo transaction.TransactionRepo
}

func NewBalanceService(transactionRepo transaction.TransactionRepo) *BalanceService {
	return &BalanceService{transactionRepo: transactionRepo}
}

func (b *BalanceService) GetBalanceByPortfolioId(userId int) (float64, error) {
	transactions, err := b.transactionRepo.ListByPortfolioId(userId)
	if err != nil {
		return 0, err
	}

	var balance float64
	for _, t := range transactions {
		if t.Type == "deposit" {
			balance += t.Amount.InexactFloat64()
		} else if t.Type == "withdrawal" {
			balance -= t.Amount.InexactFloat64()
		}
	}

	return balance, nil
}
