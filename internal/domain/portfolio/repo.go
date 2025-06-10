package portfolio

import "context"

type PortfolioRepo interface {
	Get(ctx *context.Context, id int) Portfolio
}
