package instrument

import (
	"context"
	"database/sql"
	"jollej/db-scout/internal/domain/instrument"
	"log"
)

type InstrumentRepo struct {
	db *sql.DB
}

func NewInstrumentrepo(db *sql.DB) instrument.InstrumentRepo {
	return &InstrumentRepo{db: db}
}

func (i *InstrumentRepo) List(ctx *context.Context) []instrument.Instrument {
	var res []instrument.Instrument
	query := `SELECT * FROM instruments`

	rows, err := i.db.QueryContext(*ctx, query)
	if err != nil {
		log.Fatalf("Error fetching instruments: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var instrument instrument.Instrument
		if err := rows.Scan(&instrument.Id, &instrument.Name, &instrument.Type, &instrument.Ticker, &instrument.Exchange, &instrument.Created_at, &instrument.Updated_at); err != nil {
			log.Fatalf("Error mapping instrument: %v", err)
		}
		res = append(res, instrument)
	}

	return res
}

// func Get(ctx *context.Context, id int) instrument.Instrument {
// 	var instrument instrument.Instrument
// 	query := `SELECT * FROM instruments WHERE id = ?`
//
// 	row := i.db.QueryRowContext(*ctx, query, id)
//
// 	if err := row.Scan(&instrument.Id, &instrument.Name, &instrument.Type, &instrument.Ticker, &instrument.Exchange, &instrument.Created_at, &instrument.Updated_at); err != nil {
//
// 		log.Fatalf("Could not parse the instrument: %v", err)
// 	}
//
// 	return instrument
// }
