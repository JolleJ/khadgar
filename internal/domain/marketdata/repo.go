package marketdata

import "context"

type MarketDataRepo interface {
	ListBySymbol(ctx context.Context, ticker string) ([]MarketDataPoint, error)
	AddMarketData(ctx context.Context, ticker string, data []MarketDataPoint) error
}
