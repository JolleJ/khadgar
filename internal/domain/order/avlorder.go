package order

import "log"

type AvlOrderTree struct {
	data   Order
	left   *AvlOrderTree
	right  *AvlOrderTree
	height int
}

func NewAvlOrderTree(order Order) *AvlOrderTree {
	return &AvlOrderTree{
		data:   order,
		left:   nil,
		right:  nil,
		height: 1,
	}
}

func (t *AvlOrderTree) Height() int {
	if t == nil {
		return 0
	}
	return t.height
}

func (t *AvlOrderTree) GetBalance() int {
	if t == nil {
		return 0
	}
	return t.left.Height() - t.right.Height()
}

func (t *AvlOrderTree) RotateRight() {
	if t == nil || t.left == nil {
		return
	}
	newRoot := t.left
	t.left = newRoot.right
	newRoot.right = t
	t.height = max(t.left.Height(), t.right.Height()) + 1
	newRoot.height = max(newRoot.left.Height(), newRoot.right.Height()) + 1
	*t = *newRoot
}

func (t *AvlOrderTree) RotateLeft() {
	if t == nil || t.right == nil {
		return
	}
	newRoot := t.right
	t.right = newRoot.left
	newRoot.left = t
	t.height = max(t.left.Height(), t.right.Height()) + 1
	newRoot.height = max(newRoot.left.Height(), newRoot.right.Height()) + 1
	*t = *newRoot
}

func (t *AvlOrderTree) Insert(order Order) {
	if t == nil {
		return
	}
	log.Println("Inserting order:", order)
	orderPrice, _ := order.Price.Float64()
	dataPrice, _ := t.data.Price.Float64()
	log.Println("Current node price:", dataPrice, "Order price:", orderPrice)
	if orderPrice < dataPrice {
		if t.left == nil {
			t.left = NewAvlOrderTree(order)
		} else {
			t.left.Insert(order)
		}
	} else if orderPrice > dataPrice {
		log.Println("Inserting to the right subtree")
		if t.right == nil {
			log.Println("Creating new right subtree for order:", order)
			t.right = NewAvlOrderTree(order)
		} else {
			log.Println("Inserting into existing right subtree")
			t.right.Insert(order)
		}
	} else {
		return // Duplicate keys are not allowed
	}

	t.height = max(t.left.Height(), t.right.Height()) + 1

	balance := t.GetBalance()
	leftValue, _ := t.left.data.Price.Float64()
	rightValue, _ := t.right.data.Price.Float64()
	if balance > 1 && orderPrice < leftValue {
		t.RotateRight()
		return
	}

	if balance < -1 && orderPrice > rightValue {
		t.RotateLeft()
		return
	}

	if balance > 1 && orderPrice > leftValue {
		t.left.RotateLeft()
		t.RotateRight()
		return
	}

	if balance < -1 && orderPrice < rightValue {
		t.right.RotateRight()
		t.RotateLeft()
		return
	}
}
