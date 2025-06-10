package instrument

import (
	"context"
	"jollej/db-scout/internal/domain/instrument"
)

type InstrumentService struct {
	instrumentRepo instrument.InstrumentRepo
}

func NewInstrumentService(instrumentRepo instrument.InstrumentRepo) *InstrumentService {
	return &InstrumentService{instrumentRepo: instrumentRepo}
}

func (i *InstrumentService) List(ctx context.Context) []instrument.Instrument {
	return i.instrumentRepo.List(&ctx)
}
