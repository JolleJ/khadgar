package marketdata

import "github.com/shopspring/decimal"

type MarketDataPoint struct {
	Ticker      string
	Date        string          `json:"Date"`
	Open        decimal.Decimal `json:"Open"`
	High        decimal.Decimal `json:"High"`
	Low         decimal.Decimal `json:"Low"`
	Close       decimal.Decimal `json:"Close"`
	Volume      decimal.Decimal `json:"Volume"`
	Dividends   decimal.Decimal `json:"Dividends"`
	StockSplits decimal.Decimal `json:"StockSplits"`
}
