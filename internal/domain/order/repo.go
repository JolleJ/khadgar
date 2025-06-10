package order

import "context"

type OrderRepo interface {
	Create(ctx context.Context, order Order) (int, error)
}
