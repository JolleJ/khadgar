package transaction

import (
	transactionDomain "jollej/db-scout/internal/domain/transaction"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	Id              int             `json:"Id"`
	UserId          int             `json:"UserId"`
	PortfolioId     int             `json:"PortfolioId"`
	Type            string          `json:"Type"`
	Amount          decimal.Decimal `json:"Amount"`
	Currency        string          `json:"Currency"`
	TransactionDate string          `json:"TransactionDate"`
	CreatedAt       string          `json:"CreatedAt"`
	Status          string          `json:"Status"`
}

func ToTransactionDTO(t transactionDomain.Transaction) Transaction {
	return Transaction{
		Id:              t.Id,
		UserId:          t.UserId,
		PortfolioId:     t.PortfolioId,
		Type:            t.Type,
		Amount:          t.Amount,
		Currency:        t.Currency,
		TransactionDate: t.TransactionDate,
		CreatedAt:       t.CreatedAt,
		Status:          t.Status,
	}
}

func ToTransactionDomain(t Transaction) transactionDomain.Transaction {
	return transactionDomain.Transaction{
		Id:              t.Id,
		UserId:          t.UserId,
		PortfolioId:     t.PortfolioId,
		Type:            t.Type,
		Amount:          t.Amount,
		Currency:        t.Currency,
		TransactionDate: t.TransactionDate,
		CreatedAt:       t.CreatedAt,
		Status:          t.Status,
	}
}
