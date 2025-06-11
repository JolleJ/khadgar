package test

import (
	orderDomain "jollej/db-scout/internal/domain/order"
	"log"
	"testing"

	"github.com/shopspring/decimal"
)

func TestAVLOrder(t *testing.T) {
	order := createDummyOrders()

	avl := orderDomain.NewAvlOrderTree(order[0])
	log.Println("AVL Tree initialized with first order:", avl)
	// Insert elements

	for _, ord := range order[1:] {
		avl.Insert(ord)
	}

	if avl == nil {
		t.Error("Expected AVL tree to be initialized, but it is nil")
	}

}

func createDummyOrders() []orderDomain.Order {
	return []orderDomain.Order{
		{Id: 1, Price: decimal.NewFromFloat(100.00), Side: "buy"},
		{Id: 2, Price: decimal.NewFromFloat(200.00), Side: "sell"},
		{Id: 3, Price: decimal.NewFromFloat(150.00), Side: "buy"},
		{Id: 4, Price: decimal.NewFromFloat(250.00), Side: "sell"},
		{Id: 5, Price: decimal.NewFromFloat(300.00), Side: "buy"},
		{Id: 6, Price: decimal.NewFromFloat(50.00), Side: "sell"},
		{Id: 7, Price: decimal.NewFromFloat(75.00), Side: "buy"},
		{Id: 8, Price: decimal.NewFromFloat(125.00), Side: "sell"},
		{Id: 9, Price: decimal.NewFromFloat(175.00), Side: "buy"},
		{Id: 10, Price: decimal.NewFromFloat(225.00), Side: "sell"},
		{Id: 11, Price: decimal.NewFromFloat(275.00), Side: "buy"},
		{Id: 12, Price: decimal.NewFromFloat(350.00), Side: "sell"},
	}
}
