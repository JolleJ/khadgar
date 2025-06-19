package order

import (
	"context"
	"jollej/db-scout/internal/application/trade"
	"jollej/db-scout/internal/domain/order"
	tradeDomain "jollej/db-scout/internal/domain/trade"
	"jollej/db-scout/lib/prettylog"
	"slices"
	"time"

	"github.com/shopspring/decimal"
)

type SetSlice struct {
	items    []float64
	indexMap map[float64]int
}

type OrderBook struct {
	orderService     *OrderService
	tradeService     *trade.TradeService
	Instrument       string
	buyOrdersValues  []float64
	sellOrdersValues []float64
	buyOrders        *order.AvlOrderTreeNode
	sellOrders       *order.AvlOrderTreeNode
}

func NewOrderBook(symbol string, orderService *OrderService, tradeService *trade.TradeService) *OrderBook {

	ctx := context.TODO()
	var buyOrdersValues []float64
	var buyOrders *order.AvlOrderTreeNode
	var sellOrders *order.AvlOrderTreeNode
	var sellOrdersValues []float64
	log := prettylog.NewPrettyLog()
	orders, err := orderService.ListOrdersByInstrument(ctx, symbol)
	if err != nil {
		panic("Failed to load orders for instrument: " + symbol + " - " + err.Error())
	}
	for _, ord := range orders {

		if buyOrders == nil && ord.Side == "buy" {
			buyOrders = order.NewAvlOrderTree(ord)
		} else if sellOrders == nil && ord.Side == "sell" {
			sellOrders = order.NewAvlOrderTree(ord)
		}
		if ord.Side == "buy" {
			p, exact := ord.Price.Float64()
			if !exact {
				log.Infof("Failed to convert price to float64 for order: %v", ord)
				continue
			}
			buyOrders.Insert(ord)
			buyOrdersValues = append(buyOrdersValues, p)
		} else if ord.Side == "sell" {
			p, exact := ord.Price.Float64()
			if !exact {
				log.Infof("Failed to convert price to float64 for order: %v", ord)
				continue
			}
			sellOrders.Insert(ord)
			sellOrdersValues = append(sellOrdersValues, p)
		} else {
			log.Infof("Unknown order side for order: %v", ord)
			continue
		}
	}

	slices.Sort(buyOrdersValues)
	slices.Sort(sellOrdersValues)

	return &OrderBook{
		Instrument:       symbol,
		orderService:     orderService,
		tradeService:     tradeService,
		buyOrdersValues:  buyOrdersValues,
		sellOrdersValues: sellOrdersValues,
		buyOrders:        buyOrders,
		sellOrders:       sellOrders,
	}
}

func (ob *OrderBook) AddBuyOrder(ord *order.Order) {
	ctx := context.TODO()
	//I should probably add mutex to this
	log := prettylog.NewPrettyLog()
	p, exact := ord.Price.Float64()
	if !exact {
		return
	}
	qty := ord.Quantity
	for minSellOrder := ob.sellOrders.FindMinOrderNode(); minSellOrder != nil && !qty.IsZero(); {
		log.Infof("Found minimum sell order")
		sellP, exact := minSellOrder.Key.Float64()
		log.Infof("Minimum sell order price: %f", sellP)
		log.Infof("Buy order price: %f", p)
		if !exact {
			return
		} // Check if the buy order can match with the minimum sell order
		if p < sellP {
			log.Infof("Buy order %+v cannot match with sell order at price %f", ord, sellP)
			break // No more sell orders can be matched
		}
		if minSellOrder.Data.Len() > 0 && !qty.IsZero() {
			// Get the first sell order from the minimum sell order node
			sellOrder := minSellOrder.Data.Front().Value.(*order.Order)

			switch ord.Quantity.Cmp(sellOrder.Quantity) {
			case -1:
				// Buy order quantity is less than sell order Quantity
				// Complete the buy order and update the sell order
				log.Infof("Partially matching buy order %+v with sell order %+v", *ord, sellOrder)
				sellOrder.Quantity = sellOrder.Quantity.Sub(ord.Quantity)
				log.Infof("Sell order %+v updated with new quantity %+v", sellOrder, sellOrder.Quantity)
				qty = decimal.Zero
				trade := tradeDomain.Trade{
					OrderId:    sellOrder.Id,
					Price:      sellOrder.Price,
					Quantity:   sellOrder.Quantity,
					Status:     "completed",
					ExecutedAt: time.Now().Format(time.RFC3339),
				}
				_, err := ob.tradeService.CreateTrade(ctx, trade)
				if err != nil {
					log.Errorf("Failed to create trade for order %v: %v", *ord, err)
					return
				}

				// Update the buy order status to completed
				log.Infof("Done %+v", trade)
			case 1: // Buy order quantity is greater than sell order quantity
				log.Infof("Partially matching sell order %+v with buy order %+v", sellOrder, *ord)
				qty = ord.Quantity.Sub(sellOrder.Quantity)
				trade := tradeDomain.Trade{
					OrderId:    sellOrder.Id,
					Price:      sellOrder.Price,
					Quantity:   sellOrder.Quantity,
					Status:     "completed",
					ExecutedAt: time.Now().Format(time.RFC3339),
				}
				_, err := ob.tradeService.CreateTrade(ctx, trade)
				if err != nil {
					log.Errorf("Failed to create trade for order %v: %v", *ord, err)
					return
				}
				log.Infof("Created trade for buy order %+v with sell order %+v", *ord, sellOrder)

				// Remove the sell order from the order book
				minSellOrder.Data.Remove(minSellOrder.Data.Front())
			case 0: // Buy order quantity is equal to sell order Quantity
				trade := tradeDomain.Trade{
					OrderId:    ord.Id,
					Price:      ord.Price,
					Quantity:   ord.Quantity,
					Status:     "completed",
					ExecutedAt: time.Now().Format(time.RFC3339),
				}
				_, err := ob.tradeService.CreateTrade(ctx, trade)
				if err != nil {
					log.Errorf("Failed to create trade for order %v: %v", ord, err)
					return
				}
				log.Infof("Fully matching buy order %+v with sell order %+v", ord, sellOrder)

			}

			// Add cleanup logic for the sell order if minselloder.Data.Len() == 0
			if minSellOrder.Data.Len() == 0 {
				log.Infof("Removing sell order %+v from the order book as it has been fully matched", sellOrder)
				// ob.sellOrders = ob.sellOrders.DeleteNode(minSellOrder.Key)
			}
		}
			if qty.IsZero() {
				log.Infof("Buy order %+v has been fully matched", ord)
				// Add cleanup logic for the buy order if buyOrders.Data.Len() == 0
				// Lets update the buy order status to completed 
				err := ob.orderService.SetOrderStatus(ctx, ord.Id, "completed")
				if err != nil {
					log.Errorf("Failed to update order status for order %+v: %v", ord, err)
					return
				}
			} else {
			  log.Infof("Found no more sell orders to match with buy order %+v", ord)
			}
	}
}
