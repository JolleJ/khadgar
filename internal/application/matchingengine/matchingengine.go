package matchingengine

import (
	"context"
	"jollej/db-scout/internal/application/instrument"
	orderApp "jollej/db-scout/internal/application/order"
	"jollej/db-scout/internal/domain/order"
	"jollej/db-scout/lib/prettylog"
	"log"
	"sync"
)

type MatchingEngine struct {
	instrumentService *instrument.InstrumentService
	orderService      *orderApp.OrderService
	orderChannel      *sync.Map
}

func NewMatchingEngine(instrumentService *instrument.InstrumentService, orderService *orderApp.OrderService) *MatchingEngine {

	log := prettylog.NewPrettyLog()
	orderChannels := sync.Map{}
	ctx := context.TODO()
	instrumentsList := instrumentService.List(ctx)
	for _, inst := range instrumentsList {
		orderChannels.Store(inst.Ticker, make(chan order.Order))
	}

	orderChannels.Range(func(key, value any) bool {
		log.Infof("Starting matching worker for instrument: %v", key)
		symbol := key.(string)
		ch := value.(chan order.Order)
		go func() {
			// Start the matching loop for each instrumentService
			// This will block until the channel is closed
			RunMatchingLoop(symbol, ch, orderService)
		}()
		return true
	})

	return &MatchingEngine{
		instrumentService: instrumentService,
		orderService:      orderService,
		orderChannel:      &orderChannels,
	}
}

func (me *MatchingEngine) RegisterInstrument(instrument string) {
	// Register the instrument in the matching engine
	me.orderChannel.Store(instrument, make(chan order.Order))
}

func (me *MatchingEngine) PlaceOrder(instrument string, ord order.Order) {

	ctx := context.TODO()
	id, err := me.orderService.Create(ctx, ord)
	ord.Id = id
	if err != nil {
		prettylog.NewPrettyLog().Errorf("Failed to place order: %v", err)
		return
	}
	if ch, ok := me.orderChannel.Load(instrument); ok {
		log.Println("Found existing channel for instrument:", instrument)
		ch.(chan order.Order) <- ord
	} else {
		newChan := make(chan order.Order)
		me.orderChannel.Store(instrument, newChan)
		go RunMatchingLoop(instrument, newChan, me.orderService)
		newChan <- ord
	}
}

func RunMatchingLoop(symbol string, ch chan order.Order, orderService *orderApp.OrderService) {
	log := prettylog.NewPrettyLog()
	orderbook := orderApp.NewOrderBook(symbol, orderService)
	for ord := range ch {
		log.Infof("Received order: %v", ord)
		if ord.Side == "buy" {
			orderbook.AddBuyOrder(&ord)
		} else if ord.Side == "sell" {
		}
		// Here you would implement the logic to match orders
		// For now, we just print the order
		println("Matching order:", ord.Id, "for instrument:", symbol)
	}
}
