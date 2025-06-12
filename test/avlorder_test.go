package test

import (
	orderDomain "jollej/db-scout/internal/domain/order"
	"testing"

	"github.com/shopspring/decimal"
)

func TestAVLOrder(t *testing.T) {
	order := createDummyOrders()

	avl := createAvlTree(order)

	if avl == nil {
		t.Error("Expected AVL tree to be initialized, but it is nil")
	}

	if avl.Height() == 0 {
		t.Error("Expected AVL tree height to be greater than 0, but it is 0")
	}

}

func TestAVLFindMinOrder(t *testing.T) {
	orders := createDummyOrders()
	avl := createAvlTree(orders)

	if avl == nil {
		t.Fatal("Expected AVL tree to be initialized, but it is nil")
	}

	minPriceOrder := avl.FindMinOrder()

	if minPriceOrder == nil {
		t.Fatal("Expected to find an order with minimum price, but got nil")
	}

	expectedMinPrice := decimal.NewFromFloat(50.00)
	if minPriceOrder.Price.Cmp(expectedMinPrice) != 0 {
		t.Errorf("Expected minimum price to be %s, but got %s", expectedMinPrice.String(), minPriceOrder.Price.String())
	}
}

func TestAVLFindMaxOrder(t *testing.T) {
	orders := createDummyOrders()
	avl := createAvlTree(orders)

	if avl == nil {
		t.Fatal("Expected AVL tree to be initialized, but it is nil")
	}

	maxPriceOrder := avl.FindMaxOrder()

	if maxPriceOrder == nil {
		t.Fatal("Expected to find an order with maximum price, but got nil")
	}

	expectedMaxPrice := decimal.NewFromFloat(350.00)
	if maxPriceOrder.Price.Cmp(expectedMaxPrice) != 0 {
		t.Errorf("Expected maximum price to be %s, but got %s", expectedMaxPrice.String(), maxPriceOrder.Price.String())
	}
}

func createAvlTree(orders []orderDomain.Order) *orderDomain.AvlOrderTreeNode {
	if len(orders) == 0 {
		return nil
	}

	avl := orderDomain.NewAvlOrderTree(orders[0])
	for _, ord := range orders[1:] {
		avl.Insert(ord)
	}

	return avl
}

func createDummyOrders() []orderDomain.Order {
	return []orderDomain.Order{
		{Id: 1, Price: decimal.NewFromFloat(100.00), Side: "buy", CreatedAt: "2023-10-01T12:00:00Z"},
		{Id: 2, Price: decimal.NewFromFloat(200.00), Side: "sell"},
		{Id: 3, Price: decimal.NewFromFloat(150.00), Side: "buy"},
		{Id: 4, Price: decimal.NewFromFloat(250.00), Side: "sell"},
		{Id: 5, Price: decimal.NewFromFloat(300.00), Side: "buy"},
		{Id: 6, Price: decimal.NewFromFloat(50.00), Side: "sell", CreatedAt: "2023-10-01T12:00:00Z"},
		{Id: 6, Price: decimal.NewFromFloat(50.00), Side: "sell", CreatedAt: "2023-10-01T13:00:00Z"},
		{Id: 7, Price: decimal.NewFromFloat(75.00), Side: "buy"},
		{Id: 8, Price: decimal.NewFromFloat(125.00), Side: "sell"},
		{Id: 9, Price: decimal.NewFromFloat(175.00), Side: "buy"},
		{Id: 10, Price: decimal.NewFromFloat(225.00), Side: "sell"},
		{Id: 11, Price: decimal.NewFromFloat(275.00), Side: "buy"},
		{Id: 12, Price: decimal.NewFromFloat(350.00), Side: "sell"},
	}
}
