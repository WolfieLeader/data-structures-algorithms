package avl

func (t *AVLTree[T]) Contains(value T) bool {
	curr := t.root
	for curr != nil {
		switch {
		case value < curr.Value:
			curr = curr.left
		case value > curr.Value:
			curr = curr.right
		default: // Equal
			return true
		}
	}
	return false
}

func (n *Node[T]) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *Node[T]) updateHeight() {
	if n == nil {
		return
	}
	n.height = 1 + max(n.left.getHeight(), n.right.getHeight())
}

func (n *Node[T]) balanceFactor() int {
	if n == nil {
		return 0
	}
	return n.left.getHeight() - n.right.getHeight()
}

//     X(2)                               3
//     / \                               / \
//    1   3       Left Rotation        X(2) 5
//       / \      ------------>        / \   \
//      4   5                         1   4   6
//           \
//            6

func (n *Node[T]) leftRotate() *Node[T] {
	if n == nil || n.right == nil {
		return n
	}

	pivot := n.right
	t2 := pivot.left

	pivot.left = n
	n.right = t2

	n.updateHeight()
	pivot.updateHeight()
	return pivot
}

//	       5                         4
//	      / \                       / \
//	     4   6   Right Rotation    3   5
//	    / \      ------------>    /   / \
//	   3   2                    X(1) 2   6
//	  /
//  X(1)

func (n *Node[T]) rightRotate() *Node[T] {
	if n == nil || n.left == nil {
		return n
	}

	pivot := n.left
	t2 := pivot.right

	pivot.right = n
	n.left = t2

	n.updateHeight()
	pivot.updateHeight()
	return pivot
}

func (n *Node[T]) rebalance() *Node[T] {
	if n == nil {
		return nil
	}

	bf := n.balanceFactor()
	if bf > 1 {
		if n.left.balanceFactor() < 0 {
			n.left = n.left.leftRotate()
		}
		return n.rightRotate()
	}

	if bf < -1 {
		if n.right.balanceFactor() > 0 {
			n.right = n.right.rightRotate()
		}
		return n.leftRotate()
	}

	n.updateHeight()
	return n
}

func (t *AVLTree[T]) Insert(values ...T) int {
	inserts := 0
	for _, v := range values {
		var inserted bool
		if t.root, inserted = t.root.insert(v); inserted {
			t.size++
			inserts++
		}
	}
	return inserts
}

func (n *Node[T]) insert(value T) (*Node[T], bool) {
	if n == nil {
		return &Node[T]{Value: value, height: 1}, true
	}

	var ok bool
	switch {
	case value < n.Value:
		n.left, ok = n.left.insert(value)
	case value > n.Value:
		n.right, ok = n.right.insert(value)
	default: // Equal
		return n, false
	}
	return n.rebalance(), ok
}

func (t *AVLTree[T]) Delete(values ...T) int {
	deletes := 0
	for _, v := range values {
		var deleted bool
		if t.root, deleted = t.root.delete(v); deleted {
			t.size--
			deletes++
		}
	}
	return deletes
}

func (n *Node[T]) delete(value T) (*Node[T], bool) {
	if n == nil {
		return nil, false
	}

	var ok bool
	switch {
	case value < n.Value:
		if n.left, ok = n.left.delete(value); !ok {
			return n, false
		}

	case value > n.Value:
		if n.right, ok = n.right.delete(value); !ok {
			return n, false
		}

	default: // Equal
		if n.left == nil {
			return n.right, true
		}
		if n.right == nil {
			return n.left, true
		}

		succ := n.right
		for succ.left != nil {
			succ = succ.left
		}

		n.Value = succ.Value
		n.right, _ = n.right.delete(succ.Value)
	}
	return n.rebalance(), true
}
