package position

import (
	"encoding/json"
	positionApp "jollej/db-scout/internal/application/position"
	"jollej/db-scout/internal/interfaces/api/dto/position"
	"net/http"
	"strconv"
)

type PositionHandler struct {
	positionService *positionApp.PositionService
}

func NewPositionHandler(
	positioService *positionApp.PositionService) *PositionHandler {
	return &PositionHandler{positionService: positioService}
}

func (p *PositionHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var position position.Position
	positionId := r.PathValue("id")
	if positionId == "" {
		http.Error(w, "No id value found in path", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(positionId)
	if err != nil {
		http.Error(w, "Invalid id given in path", http.StatusBadRequest)
		return
	}

	portfolio := p.positionService.Get(r.Context(), id)
	if portfolio.Id == 0 {
		http.Error(w, "No portfolio found", http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(position); err != nil {
		http.Error(w, "Error encoding the response", http.StatusInternalServerError)
		return
	}
}
