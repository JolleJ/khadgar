package order

import (
	"container/list"
	"log"

	"github.com/shopspring/decimal"
)

type AvlOrderTreeNode struct {
	key    decimal.Decimal
	data   *list.List
	left   *AvlOrderTreeNode
	right  *AvlOrderTreeNode
	height int
}

func NewAvlOrderTree(order Order) *AvlOrderTreeNode {
	l := list.New()
	l.PushBack(&order)
	return &AvlOrderTreeNode{
		key:    order.Price,
		data:   l,
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
	dataPrice, _ := t.key.Float64()
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
		t.data.PushBack(&order)
		return t
	}

	t.height = max(t.left.Height(), t.right.Height()) + 1

	balance := t.GetBalance()
	if balance > 1 {
		if t.left != nil {
			leftValue, _ := t.left.key.Float64()
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
			rightValue, _ := t.right.key.Float64()
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

func (t *AvlOrderTreeNode) FindMinOrder() *Order {
	if t == nil {
		return nil
	}
	current := t
	for current.left != nil {
		current = current.left

	}
	oldest := current.data.Front()
	if oldest == nil {
		log.Println("No orders found in the AVL tree")
		return nil
	}
	return oldest.Value.(*Order)
}

func (t *AvlOrderTreeNode) FindMaxOrder() *Order {
	if t == nil {
		return nil
	}

	current := t
	for current.right != nil {
		current = current.right
	}
	oldest := current.data.Front()
	if oldest == nil {
		return nil
	}
	return oldest.Value.(*Order)
}
