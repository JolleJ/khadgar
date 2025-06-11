package transaction

import "github.com/shopspring/decimal"

type Transaction struct {
	Id              int
	UserId          int
	PortfolioId     int
	Type            string
	Amount          decimal.Decimal
	Currency        string
	TransactionDate string
	CreatedAt       string
	Status          string
}
