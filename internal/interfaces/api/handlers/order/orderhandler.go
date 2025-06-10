package order

import (
	"encoding/json"
	orderApplication "jollej/db-scout/internal/application/order"
	"jollej/db-scout/internal/interfaces/api/dto/order"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	orderService *orderApplication.OrderService
}

func NewOrderHandler(orderService *orderApplication.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	portfolioId := r.PathValue("id")
	if portfolioId == "" {
		http.Error(w, "Portfolio ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(portfolioId)
	if err != nil {
		http.Error(w, "Invalid Portfolio ID", http.StatusBadRequest)
		return
	}

	var orderDtoDetails order.OrderDetails
	if err := json.NewDecoder(r.Body).Decode(&orderDtoDetails); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	orderDto := order.Order{Id: id, Details: orderDtoDetails}

	orderDomain := order.ToOrderDomain(orderDto)
	orderId, err := h.orderService.Create(r.Context(), orderDomain)
	if err != nil {
		http.Error(w, "Error creating order: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := map[string]int{"id": orderId}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding the response", http.StatusInternalServerError)
		return
	}
}
