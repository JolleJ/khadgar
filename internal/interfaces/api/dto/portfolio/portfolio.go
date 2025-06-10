package portfolio

import portfolioDomain "jollej/db-scout/internal/domain/portfolio"

type Portfolio struct {
	Id          int    `json:"Id"`
	User_id     int    `json:"User_id"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Created_at  string `json:"Created_at"`
	Updated_at  string `json:"Updated_at"`
}

func ToPortfolioDto(portfolio portfolioDomain.Portfolio) Portfolio {
	return Portfolio{
		Id:          portfolio.Id,
		User_id:     portfolio.User_id,
		Name:        portfolio.Name,
		Description: portfolio.Description,
		Created_at:  portfolio.Created_at,
		Updated_at:  portfolio.Updated_at,
	}
}

func ToPortfolioDomain(portfolio Portfolio) portfolioDomain.Portfolio {
	return portfolioDomain.Portfolio{
		Id:          portfolio.Id,
		User_id:     portfolio.User_id,
		Name:        portfolio.Name,
		Description: portfolio.Description,
		Created_at:  portfolio.Created_at,
		Updated_at:  portfolio.Updated_at,
	}
}
