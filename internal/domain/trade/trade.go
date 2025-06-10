package trade

import "github.com/shopspring/decimal"

type Trade struct {
	Id         int
	OrderId    int
	Quantity   decimal.Decimal
	Price      decimal.Decimal
	ExecutedAt string
	Status     string
}
