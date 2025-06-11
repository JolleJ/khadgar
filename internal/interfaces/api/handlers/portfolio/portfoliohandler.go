package portfolio

import (
	"encoding/json"
	"jollej/db-scout/internal/application/balance"
	"jollej/db-scout/internal/application/portfolio"
	balanceDto "jollej/db-scout/internal/interfaces/api/dto/balance"
	portfolioDto "jollej/db-scout/internal/interfaces/api/dto/portfolio"
	"net/http"
	"strconv"
)

type PortfolioHandler struct {
	portfolioService *portfolio.PortfolioService
	balanceService   *balance.BalanceService
}

func NewPortfolioHandler(portfolioService *portfolio.PortfolioService, balanceService *balance.BalanceService) *PortfolioHandler {
	return &PortfolioHandler{portfolioService: portfolioService, balanceService: balanceService}
}

func (a *PortfolioHandler) GetPortfolio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var portfolioResponse portfolioDto.Portfolio
	portfolioId := r.PathValue("id")
	if portfolioId == "" {
		http.Error(w, "No id value found in path", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(portfolioId)
	if err != nil {
		http.Error(w, "Invalid id given in path", http.StatusBadRequest)
		return
	}

	portfolio := a.portfolioService.Get(r.Context(), id)
	if portfolio.Id == 0 {
		http.Error(w, "No portfolio found", http.StatusNotFound)
	}
	portfolioResponse = portfolioDto.ToPortfolioDto(portfolio)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(portfolioResponse); err != nil {
		http.Error(w, "Error encoding the response", http.StatusInternalServerError)
		return
	}
}

func (a *PortfolioHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")
	portfolioId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return

	}
	balance, err := a.balanceService.GetBalanceByPortfolioId(portfolioId)
	if err != nil {
		http.Error(w, "Error retrieving balance", http.StatusInternalServerError)
		return
	}
	balanceResponse := balanceDto.GetBalanceResponse{
		Balance: balanceDto.Balance{Amount: balance, Currency: "USD"},
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(balanceResponse); err != nil {
		http.Error(w, "Error encoding the response", http.StatusInternalServerError)
		return
	}
}
