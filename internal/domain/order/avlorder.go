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

func (t *AvlOrderTree) RotateRight() *AvlOrderTree {
	newRoot := t.left
	t.left = newRoot.right
	newRoot.right = t

	t.height = max(t.left.Height(), t.right.Height()) + 1
	newRoot.height = max(newRoot.left.Height(), newRoot.right.Height()) + 1

	return newRoot
}

func (t *AvlOrderTree) RotateLeft() *AvlOrderTree {
	newRoot := t.right
	t.right = newRoot.left
	newRoot.left = t

	t.height = max(t.left.Height(), t.right.Height()) + 1
	newRoot.height = max(newRoot.left.Height(), newRoot.right.Height()) + 1

	return newRoot
}

func (t *AvlOrderTree) Insert(order Order) {
	if t == nil {
		return
	}
	log.Println("Inserting order:", order)
	orderPrice, _ := order.Price.Float64()
	dataPrice, _ := t.data.Price.Float64()
	if orderPrice < dataPrice {
		if t.left == nil {
			t.left = NewAvlOrderTree(order)
		} else {
			t.left.Insert(order)
		}
	} else if orderPrice > dataPrice {
		if t.right == nil {
			t.right = NewAvlOrderTree(order)
		} else {
			t.right.Insert(order)
		}
	} else {
		return // Duplicate keys are not allowed
	}

	t.height = max(t.left.Height(), t.right.Height()) + 1

	balance := t.GetBalance()
	if balance > 1 {
		if t.left != nil {
			leftValue, _ := t.left.data.Price.Float64()
			if orderPrice < leftValue {
				t = t.RotateRight()
				return
			} else if orderPrice > leftValue {
				t = t.left.RotateLeft()
				t = t.RotateRight()
				return
			}
		}
	}

	if balance < -1 {
		if t.right != nil {
			rightValue, _ := t.right.data.Price.Float64()
			if orderPrice > rightValue {
				t = t.RotateLeft()
				return
			} else if orderPrice < rightValue {
				t = t.right.RotateRight()
				t = t.RotateLeft()
				return
			}
		}
	}

}
