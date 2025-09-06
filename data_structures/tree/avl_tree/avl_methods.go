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

func (t *AVLTree[T]) leftRotate(n *Node[T]) *Node[T] {
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

func (t *AVLTree[T]) rightRotate(n *Node[T]) *Node[T] {
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

func (t *AVLTree[T]) rebalance(n *Node[T]) *Node[T] {
	bf := n.balanceFactor()
	if bf > 1 {
		if n.left.balanceFactor() < 0 {
			n.left = t.leftRotate(n.left)
		}
		return t.rightRotate(n)
	}

	if bf < -1 {
		if n.right.balanceFactor() > 0 {
			n.right = t.rightRotate(n.right)
		}
		return t.leftRotate(n)
	}

	n.updateHeight()
	return n
}

func (t *AVLTree[T]) Insert(values ...T) {
	for _, v := range values {
		t.root = t.insert(t.root, v)
	}
}

func (t *AVLTree[T]) insert(n *Node[T], v T) *Node[T] {
	if n == nil {
		t.size++
		return &Node[T]{Value: v, height: 1}
	}

	switch {
	case v < n.Value:
		n.left = t.insert(n.left, v)
	case v > n.Value:
		n.right = t.insert(n.right, v)
	default: // Equal
		return n
	}
	return t.rebalance(n)
}
