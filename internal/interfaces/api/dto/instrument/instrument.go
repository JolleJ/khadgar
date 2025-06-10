package instrument

import instrumentDomain "jollej/db-scout/internal/domain/instrument"

type Instrument struct {
	Id         int    `json:"Id"`
	Name       string `json:"Name"`
	Type       string `json:"Type"`
	Ticker     string `json:"Ticker"`
	Exchange   string `json:"Exchange"`
	Created_at string `json:"Created_at"`
	Updated_at string `json:"Updated_at"`
}

type ListInstrumentsResponse struct {
	Instruments []Instrument `json:"instruments"`
}

func ToInsturmentDto(instrument instrumentDomain.Instrument) Instrument {
	return Instrument{
		Id:         instrument.Id,
		Name:       instrument.Name,
		Type:       instrument.Type,
		Ticker:     instrument.Ticker,
		Exchange:   instrument.Exchange,
		Created_at: instrument.Created_at,
		Updated_at: instrument.Updated_at,
	}
}

func ToInstrumentDomain(instrument Instrument) instrumentDomain.Instrument {
	return instrumentDomain.Instrument{
		Id:         instrument.Id,
		Name:       instrument.Name,
		Type:       instrument.Type,
		Ticker:     instrument.Ticker,
		Exchange:   instrument.Exchange,
		Created_at: instrument.Created_at,
		Updated_at: instrument.Updated_at,
	}
}
