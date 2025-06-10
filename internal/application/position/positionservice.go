package position

import (
	"context"
	"jollej/db-scout/internal/domain/position"
)

type PositionService struct {
	positionRepo position.PositionRepo
}

func NewPositionService(positionRepo position.PositionRepo) *PositionService {
	return &PositionService{positionRepo: positionRepo}
}

func (p *PositionService) Get(ctx context.Context, id int) position.Position {
	return p.positionRepo.Get(&ctx, id)
}
