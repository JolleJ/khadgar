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
