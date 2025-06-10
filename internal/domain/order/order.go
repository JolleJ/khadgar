package order

import "github.com/shopspring/decimal"

type Order struct {
	Id           int
	PortfolioId  int
	InstrumentId int
	Side         string
	Quantity     decimal.Decimal
	Price        decimal.Decimal
	OrderDate    string
	Status       string
	CreatedAt    string
	UpdatedAt    string
	FilledAt     string
}
