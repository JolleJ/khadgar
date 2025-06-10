package trade

import (
	"context"
	"jollej/db-scout/internal/domain/trade"
)

type TradeService struct {
	tradeRepo trade.TradeRepo
}

func NewTradeService(tradeRepo trade.TradeRepo) *TradeService {
	return &TradeService{
		tradeRepo: tradeRepo,
	}
}

func (t *TradeService) CreateTrade(ctx context.Context, trade trade.Trade) (int, error) {
	return t.tradeRepo.CreateTrade(ctx, trade)
}

func (t *TradeService) ListTradesOnPortfolio(ctx context.Context, id int) ([]trade.Trade, error) {
	return t.tradeRepo.ListTradesOnPortfolio(ctx, id)
}
