package order

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

func (t *AvlOrderTree) Insert(order Order) *AvlOrderTree {
	if t == nil {
		return nil
	}
	orderPrice, _ := order.Price.Float64()
	dataPrice, _ := t.data.Price.Float64()
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
		return nil
	}

	t.height = max(t.left.Height(), t.right.Height()) + 1

	balance := t.GetBalance()
	if balance > 1 {
		if t.left != nil {
			leftValue, _ := t.left.data.Price.Float64()
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
			rightValue, _ := t.right.data.Price.Float64()
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

func (t *AvlOrderTree) FindMinOrder() *Order {
	if t == nil {
		return nil
	}
	current := t
	for current.left != nil {
		current = current.left
	}
	return &current.data
}

func (t *AvlOrderTree) FindMaxOrder() *Order {
	if t == nil {
		return nil
	}

	current := t
	for current.right != nil {
		current = current.right
	}
	return &current.data
}
