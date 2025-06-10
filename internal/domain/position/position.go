package position

import "github.com/shopspring/decimal"

type Position struct {
	Id           int
	PortfolioId  int
	IstrumentId  int
	Quantity     decimal.Decimal
	AveragePrice decimal.Decimal
}
