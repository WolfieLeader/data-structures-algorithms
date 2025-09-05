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

//     X(2)                               3
//     / \                               / \
//    1   3       Left Rotation        X(2) 5
//       / \      ------------>        / \   \
//      4   5                         1   4   6
//           \
//            6

func (t *AVLTree[T]) leftRotate(n *Node[T]) *Node[T] {
	if t.root == nil {
		return nil
	}

	if n == nil || n.right == nil {
		return n
	}

	pivot := n.right
	moved := pivot.left

	pivot.left = n
	n.right = moved

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
	moved := pivot.right

	pivot.right = n
	n.left = moved

	n.updateHeight()
	pivot.updateHeight()
	return pivot
}
