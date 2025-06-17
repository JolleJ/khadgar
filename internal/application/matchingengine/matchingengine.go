package matchingengine

import (
	"context"
	"jollej/db-scout/internal/application/instrument"
	orderApp "jollej/db-scout/internal/application/order"
	"jollej/db-scout/internal/application/trade"
	"jollej/db-scout/internal/domain/order"
	"jollej/db-scout/lib/prettylog"
	"log"
	"sync"
)

type MatchingEngine struct {
	instrumentService *instrument.InstrumentService
	tradeService      *trade.TradeService
	orderService      *orderApp.OrderService
	orderChannel      *sync.Map
}

func NewMatchingEngine(instrumentService *instrument.InstrumentService, orderService *orderApp.OrderService, tradeService *trade.TradeService) *MatchingEngine {

	orderChannels := sync.Map{}
	ctx := context.TODO()
	instrumentsList := instrumentService.List(ctx)
	for _, inst := range instrumentsList {
		orderChannels.Store(inst.Ticker, make(chan order.Order))
	}

	orderChannels.Range(func(key, value any) bool {
		// log.Infof("Starting matching worker for instrument: %v", key)
		symbol := key.(string)
		ch := value.(chan order.Order)
		go func() {
			// Start the matching loop for each instrumentService
			// This will block until the channel is closed
			RunMatchingLoop(symbol, ch, orderService, tradeService)
		}()
		return true
	})

	return &MatchingEngine{
		instrumentService: instrumentService,
		orderService:      orderService,
		orderChannel:      &orderChannels,
		tradeService:      tradeService,
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
		//For testing puropse we wait for the channel to be processed, in production this should be removed
		// this should be replaced with a more robust mechanism, for example, using a message queue or a more sophisticated event system TODO
		// send event notification to the order service
		<-ch.(chan order.Order) // Ensure the channel is not blocked
	} else {
		newChan := make(chan order.Order)
		me.orderChannel.Store(instrument, newChan)
		go RunMatchingLoop(instrument, newChan, me.orderService, me.tradeService)
		newChan <- ord
		<-newChan // Ensure the channel is not blocked
	}

}

func RunMatchingLoop(symbol string, ch chan order.Order, orderService *orderApp.OrderService, tradeService *trade.TradeService) {
	orderbook := orderApp.NewOrderBook(symbol, orderService, tradeService)
	for ord := range ch {
		if ord.Side == "buy" {
			orderbook.AddBuyOrder(&ord)
			ch <- ord // Echo back the order to the channel for testing purposes this has to be revisited later
		} else if ord.Side == "sell" {

		}
	}
}
