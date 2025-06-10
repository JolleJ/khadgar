package instrument

import (
	"encoding/json"
	"jollej/db-scout/internal/application/instrument"
	instrumentDto "jollej/db-scout/internal/interfaces/api/dto/instrument"
	"net/http"
)

type InstrumentHandler struct {
	instrumentService *instrument.InstrumentService
}

func NewInstrumentHandler(instrumentService *instrument.InstrumentService) *InstrumentHandler {
	return &InstrumentHandler{instrumentService: instrumentService}
}

func (a *InstrumentHandler) ListInstruments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var listInstrumentsResponse instrumentDto.ListInstrumentsResponse
	instruments := a.instrumentService.List(r.Context())
	for _, instrumet := range instruments {
		listInstrumentsResponse.Instruments = append(listInstrumentsResponse.Instruments, instrumentDto.ToInsturmentDto(instrumet))
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(listInstrumentsResponse); err != nil {
		http.Error(w, "Error encoding the response", http.StatusInternalServerError)
		return
	}

}
