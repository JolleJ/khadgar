package portfolio

import (
	"context"
	"database/sql"
	"jollej/db-scout/internal/domain/portfolio"
	"log"
)

type PortfolioRepo struct {
	db *sql.DB
}

func NewPortfolioRepo(db *sql.DB) portfolio.PortfolioRepo {
	return &PortfolioRepo{db: db}
}

func (p *PortfolioRepo) Get(ctx *context.Context, id int) portfolio.Portfolio {

	getQuery := `SELECT * FROM portfolios where id = ?;`

	row := p.db.QueryRowContext(*ctx, getQuery, id)
	var portfolio portfolio.Portfolio
	if err := row.Scan(&portfolio.Id, &portfolio.User_id, &portfolio.Name, &portfolio.Description, &portfolio.Created_at, &portfolio.Updated_at); err != nil {
		log.Fatalf("Error fetching portfolio: %v", err)
	}

	return portfolio
}
