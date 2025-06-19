package order

import (
	"context"
	"jollej/db-scout/internal/domain/order"
)

type OrderService struct {
	orderRepo order.OrderRepo
}

func NewOrderService(orderRepo order.OrderRepo) *OrderService {
	return &OrderService{orderRepo: orderRepo}
}

func (o *OrderService) Create(ctx context.Context, order order.Order) (int, error) {
	return o.orderRepo.Create(ctx, order)
}

func (o *OrderService) ListOrdersByInstrument(ctx context.Context, symbol string) ([]order.Order, error) {
	orders, err := o.orderRepo.ListByInstrument(ctx, symbol)
	return orders, err
}

func (o *OrderService) SetOrderStatus(ctx context.Context, orderID int, status string) error {
	return o.orderRepo.SetStatus(ctx, orderID, status)
}
