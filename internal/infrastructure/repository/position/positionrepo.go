package position

import (
	"context"
	"database/sql"
	"jollej/db-scout/internal/domain/position"
	"log"
)

type PositionRepo struct {
	db *sql.DB
}

func NewPositionRepo(db *sql.DB) position.PositionRepo {
	return &PositionRepo{db: db}
}

func (p *PositionRepo) Get(ctx *context.Context, id int) position.Position {
	query := `SELECT * FROM positions WHERE id = ?`

	row := p.db.QueryRowContext(*ctx, query, id)
	var position position.Position
	if err := row.Scan(&position.Id, &position.PortfolioId, &position.IstrumentId, &position.Quantity, &position.AveragePrice); err != nil {
		log.Fatalf("Error scanning position: %v", err)
	}

	return position
}
