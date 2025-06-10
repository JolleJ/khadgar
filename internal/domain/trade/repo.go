package trade

import "context"

type TradeRepo interface {
	CreateTrade(ctx context.Context, trade Trade) (int, error)
	ListTradesOnPortfolio(ctx context.Context, portfolioId int) ([]Trade, error)
}
