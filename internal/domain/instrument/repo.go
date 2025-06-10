package instrument

import (
	"context"
)

type InstrumentRepo interface {
	List(ctx *context.Context) []Instrument
}
