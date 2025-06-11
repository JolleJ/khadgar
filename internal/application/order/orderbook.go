package order

import (
	"context"
	"jollej/db-scout/internal/domain/order"
	"log"
)

type SetSlice struct {
	items    []float64
	indexMap map[float64]int
}

type OrderBook struct {
	orderService     *OrderService
	Instrument       string
	buyOrdersValues  []float64
	sellOrdersValues []float64
	buyOrders        map[float64]order.Order
	sellOrders       map[float64]order.Order
}

func NewOrderBook(symbol string, orderService *OrderService) *OrderBook {

	ctx := context.TODO()
	var buyOrdersValues []float64
	buyOrders := make(map[float64]order.Order)
	sellOrders := make(map[float64]order.Order)
	var sellOrdersValues []float64

	orders, err := orderService.ListOrdersByInstrument(ctx, symbol)
	if err != nil {
		panic("Failed to load orders for instrument: " + symbol + " - " + err.Error())
	}

	for _, ord := range orders {
		if ord.Side == "buy" {
			p, exact := ord.Price.Float64()
			if !exact {
				log.Default().Println("Failed to convert price to float64 for order:", ord)
				continue
			}
			buyOrdersValues = append(buyOrdersValues, p)
			buyOrders[p] = ord
		} else if ord.Side == "sell" {
		}
	}

	return &OrderBook{
		Instrument:       symbol,
		orderService:     orderService,
		buyOrdersValues:  buyOrdersValues,
		sellOrdersValues: sellOrdersValues,
		buyOrders:        buyOrders,
		sellOrders:       sellOrders,
	}
}
func (ob *OrderBook) AddBuyOrder(ord order.Order) {
}

func (ob *OrderBook) AddSellOrder(ord order.Order) {
}
