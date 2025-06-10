package order

import (
	"github.com/shopspring/decimal"
	orderDomain "jollej/db-scout/internal/domain/order"
)

type Order struct {
	Id          int          `json:"Id"`
	CreatedAt   string       `json:"CreatedAt"`
	UpdatedAt   string       `json:"UpdatedAt"`
	FilledAt    string       `json:"FilledAt"`
	Details     OrderDetails `json:"Details"`
	PortfolioId int          `json:"PortfolioId"`
}

type OrderDetails struct {
	InstrumentId int             `json:"InstrumentId"`
	Side         string          `json:"Side"`
	Quantity     decimal.Decimal `json:"Quantity"`
	Price        decimal.Decimal `json:"Price"`
	OrderDate    string          `json:"OrderDate"`
	Status       string          `json:"Status"`
}

func ToOrderDetailsDto(order orderDomain.Order) Order {
	return Order{
		Id:          order.Id,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
		FilledAt:    order.FilledAt,
		PortfolioId: order.PortfolioId,
		Details: OrderDetails{
			InstrumentId: order.InstrumentId,
			Side:         order.Side,
			Quantity:     order.Quantity,
			Price:        order.Price,
			OrderDate:    order.OrderDate,
			Status:       order.Status,
		},
	}
}

func ToOrderDomain(order Order) orderDomain.Order {
	return orderDomain.Order{
		Id:           order.Id,
		CreatedAt:    order.CreatedAt,
		UpdatedAt:    order.UpdatedAt,
		FilledAt:     order.FilledAt,
		PortfolioId:  order.PortfolioId,
		InstrumentId: order.Details.InstrumentId,
		Side:         order.Details.Side,
		Quantity:     order.Details.Quantity,
		Price:        order.Details.Price,
		OrderDate:    order.Details.OrderDate,
		Status:       order.Details.Status,
	}
}
