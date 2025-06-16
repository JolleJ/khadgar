package order

import (
	"container/list"
	"log"

	"github.com/shopspring/decimal"
)

type AvlOrderTreeNode struct {
	Key    decimal.Decimal
	Data   *list.List
	left   *AvlOrderTreeNode
	right  *AvlOrderTreeNode
	height int
}

func NewAvlOrderTree(order Order) *AvlOrderTreeNode {
	l := list.New()
	l.PushBack(&order)
	return &AvlOrderTreeNode{
		Key:    order.Price,
		Data:   l,
		left:   nil,
		right:  nil,
		height: 1,
	}
}

func (t *AvlOrderTreeNode) Height() int {
	if t == nil {
		return 0
	}
	return t.height
}

func (t *AvlOrderTreeNode) GetBalance() int {
	if t == nil {
		return 0
	}
	return t.left.Height() - t.right.Height()
}

func (t *AvlOrderTreeNode) RotateRight() *AvlOrderTreeNode {
	newRoot := t.left
	t.left = newRoot.right
	newRoot.right = t

	t.height = max(t.left.Height(), t.right.Height()) + 1
	newRoot.height = max(newRoot.left.Height(), newRoot.right.Height()) + 1

	return newRoot
}

func (t *AvlOrderTreeNode) RotateLeft() *AvlOrderTreeNode {
	newRoot := t.right
	t.right = newRoot.left
	newRoot.left = t

	t.height = max(t.left.Height(), t.right.Height()) + 1
	newRoot.height = max(newRoot.left.Height(), newRoot.right.Height()) + 1

	return newRoot
}

func (t *AvlOrderTreeNode) Insert(order Order) *AvlOrderTreeNode {
	if t == nil {
		return nil
	}
	orderPrice, _ := order.Price.Float64()
	dataPrice, _ := t.Key.Float64()
	if orderPrice < dataPrice {
		if t.left == nil {
			t.left = NewAvlOrderTree(order)
		} else {
			t.left = t.left.Insert(order)
		}
	} else if orderPrice > dataPrice {
		if t.right == nil {
			t.right = NewAvlOrderTree(order)
		} else {
			t.right = t.right.Insert(order)
		}
	} else {
		// If the price is the same, append the order to the existing data slice
		// The *t.data will have to be sorted according to create date where the first order is the newest
		t.Data.PushBack(&order)
		return t
	}

	t.height = max(t.left.Height(), t.right.Height()) + 1

	// Run rebalance here
	balance := t.GetBalance()
	if balance > 1 {
		if t.left != nil {
			leftValue, _ := t.left.Key.Float64()
			if orderPrice < leftValue {
				return t.RotateRight()
			} else if orderPrice > leftValue {
				t.left = t.left.RotateLeft()
				return t.left.RotateRight()
			}
		}
	}

	if balance < -1 {
		if t.right != nil {
			rightValue, _ := t.right.Key.Float64()
			if orderPrice > rightValue {
				return t.RotateLeft()
			} else if orderPrice < rightValue {
				t.right = t.right.RotateRight()
				return t.right.RotateLeft()
			}
		}
	}

	return t
}

func (t *AvlOrderTreeNode) FindLowestOrderFromPrice(price decimal.Decimal) *Order {

	if t == nil {
		return nil
	}

	current := t
	for current.left != nil {
		current = current.left

	}
	oldest := current.Data.Front()
	if oldest == nil {
		log.Println("No orders found in the AVL tree")
		return nil
	}
	return oldest.Value.(*Order)
}
func (t *AvlOrderTreeNode) FindMinOrder() *Order {
	if t == nil {
		return nil
	}
	current := t
	for current.left != nil {
		current = current.left

	}
	oldest := current.Data.Front()
	if oldest == nil {
		log.Println("No orders found in the AVL tree")
		return nil
	}
	return oldest.Value.(*Order)
}

func (t *AvlOrderTreeNode) FindMinOrderNode() *AvlOrderTreeNode {
	log.Println("Finding minimum order node in AVL tree")
	if t == nil {
		return nil
	}
	current := t
	log.Printf("Starting from root node with key: %s", current.Key.String())
	for current.left != nil {
		log.Printf("Current node key: %s", current.Key.String())
		current = current.left
	}
	log.Printf("Reached leftmost node with key: %s", current.Key.String())
	if current.Data == nil || current.Data.Len() == 0 {
		log.Println("No orders found in the AVL tree")
		return nil
	}
	oldest := current.Data.Front()
	log.Printf("Oldest order found with price: %s", current.Key.String())
	if oldest == nil {
		log.Println("No orders found in the AVL tree")
		return nil
	}
	log.Printf("Found minimum order with price: %s", current.Key.String())
	return t
}

func (t *AvlOrderTreeNode) UpdateOrder(order Order) *AvlOrderTreeNode {
	if t == nil {
		return nil
	}

	switch order.Price.Cmp(t.Key) {
	case -1:
		if t.left != nil {
			t.left.UpdateOrder(order)
		}
	case 1:
		if t.right != nil {
			t.right.UpdateOrder(order)
		}
	case 0:
		if t.Data.Len() > 1 {

		}

	}
	return nil
}

func (t *AvlOrderTreeNode) FindMaxOrder() *Order {
	if t == nil {
		return nil
	}

	current := t
	for current.right != nil {
		current = current.right
	}
	oldest := current.Data.Front()
	if oldest == nil {
		return nil
	}
	return oldest.Value.(*Order)
}
