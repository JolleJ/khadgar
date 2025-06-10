package portfolio

import (
	"context"
	"jollej/db-scout/internal/domain/portfolio"
)

type PortfolioService struct {
	portfolioRepo portfolio.PortfolioRepo
}

func NewPortfolioService(portfolioRepo portfolio.PortfolioRepo) *PortfolioService {
	return &PortfolioService{portfolioRepo: portfolioRepo}
}

func (p *PortfolioService) Get(ctx context.Context, id int) portfolio.Portfolio {
	return p.portfolioRepo.Get(&ctx, id)
}
