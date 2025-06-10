package trade

import (
	"jollej/db-scout/internal/application/trade"
	"net/http"
)

type TradeHandler struct {
	tradeService *trade.TradeService
}

func NewTradeHandler(tradeService *trade.TradeService) *TradeHandler {
	return &TradeHandler{tradeService: tradeService}
}

func (t *TradeHandler) ListTradesOnPortfolio(w http.ResponseWriter, r *http.Request) {
}
