package position

import "context"

type PositionRepo interface {
	Get(ctx *context.Context, id int) Position
}
