package position

import (
	positionDomain "jollej/db-scout/internal/domain/position"

	"github.com/shopspring/decimal"
)

type Position struct {
	Id           int             `json:"id"`
	PortfolioId  int             `json:"portfolioId"`
	IstrumentId  int             `json:"istrumentId"`
	Quantity     decimal.Decimal `json:"quantity"`
	AveragePrice decimal.Decimal `json:"averagePrice"`
}

func ToPositionDto(position positionDomain.Position) Position {
	return Position{
		Id:           position.Id,
		PortfolioId:  position.PortfolioId,
		IstrumentId:  position.IstrumentId,
		Quantity:     position.Quantity,
		AveragePrice: position.AveragePrice,
	}
}

func ToPositionDomain(position Position) positionDomain.Position {
	return positionDomain.Position{
		Id:           position.Id,
		PortfolioId:  position.PortfolioId,
		IstrumentId:  position.IstrumentId,
		Quantity:     position.Quantity,
		AveragePrice: position.AveragePrice,
	}
}
