package bst

import "math"

func (t *BST[T]) ContainsR(value T) bool { return t.root.contains(value) }
func (n *Node[T]) contains(value T) bool {
	if n == nil {
		return false
	}

	switch {
	case value < n.Value:
		return (n.left).contains(value)
	case value > n.Value:
		return (n.right).contains(value)
	default: // Equal
		return true
	}
}

func (t *BST[T]) InsertR(value T) bool {
	newRoot, ok := t.root.insert(value)
	t.root = newRoot
	if ok {
		t.size++
	}
	return ok
}

func (n *Node[T]) insert(value T) (*Node[T], bool) {
	if n == nil {
		return &Node[T]{Value: value}, true
	}

	switch {
	case value < n.Value:
		newLeft, ok := (n.left).insert(value)
		n.left = newLeft
		return n, ok

	case value > n.Value:
		newRight, ok := (n.right).insert(value)
		n.right = newRight
		return n, ok

	default: // Equal
		return n, false
	}
}

func (t *BST[T]) DeleteR(value T) bool {
	newRoot, ok := t.root.delete(value)
	t.root = newRoot
	if ok {
		t.size--
	}
	return ok
}

func (n *Node[T]) delete(value T) (*Node[T], bool) {
	if n == nil {
		return nil, false
	}

	switch {
	case value < n.Value:
		newLeft, ok := (n.left).delete(value)
		n.left = newLeft
		return n, ok

	case value > n.Value:
		newRight, ok := (n.right).delete(value)
		n.right = newRight
		return n, ok

	default: //Equal
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
		newRight, ok := (n.right).delete(succ.Value)
		n.right = newRight
		return n, ok
	}
}

func (t *BST[T]) MinR() (T, bool) {
	var zero T
	if t.root == nil {
		return zero, false
	}
	return t.root.min(), true
}

func (n *Node[T]) min() T {
	if n.left == nil {
		return n.Value
	}
	return n.left.min()
}

func (t *BST[T]) MaxR() (T, bool) {
	var zero T
	if t.root == nil {
		return zero, false
	}
	return t.root.max(), true
}

func (n *Node[T]) max() T {
	if n.right == nil {
		return n.Value
	}
	return n.right.max()
}

func (t *BST[T]) TraverseInOrderR(fn func(value T)) { t.root.inOrder(fn) }
func (n *Node[T]) inOrder(fn func(value T)) {
	if n == nil {
		return
	}
	(n.left).inOrder(fn)
	fn(n.Value)
	(n.right).inOrder(fn)
}

func (t *BST[T]) TraversePreOrderR(fn func(value T)) { t.root.preOrder(fn) }
func (n *Node[T]) preOrder(fn func(value T)) {
	if n == nil {
		return
	}
	fn(n.Value)
	(n.left).preOrder(fn)
	(n.right).preOrder(fn)
}

func (t *BST[T]) TraversePostOrderR(fn func(value T)) { t.root.postOrder(fn) }
func (n *Node[T]) postOrder(fn func(value T)) {
	if n == nil {
		return
	}
	(n.left).postOrder(fn)
	(n.right).postOrder(fn)
	fn(n.Value)
}

func (t *BST[T]) HeightR() int { return t.root.height() }
func (n *Node[T]) height() int {
	if n == nil {
		return 0
	}
	return max((n.left).height(), (n.right).height()) + 1
}

func (t *BST[T]) IsBalanced() bool { _, ok := t.root.balanced(); return ok }
func (n *Node[T]) balanced() (int, bool) {
	if n == nil {
		return 0, true
	}
	leftHeight, leftBalanced := (n.left).balanced()
	if !leftBalanced {
		return 0, false
	}
	rightHeight, rightBalanced := (n.right).balanced()
	if !rightBalanced {
		return 0, false
	}
	return max(leftHeight, rightHeight) + 1, math.Abs(float64(leftHeight-rightHeight)) <= 1
}

func (t *BST[T]) CopyR() *BST[T] {
	if t.root == nil {
		return New[T]()
	}
	newTree := New[T]()
	newTree.root = t.root.copy()
	newTree.size = t.size
	return newTree
}

func (n *Node[T]) copy() *Node[T] {
	if n == nil {
		return nil
	}
	return &Node[T]{Value: n.Value, left: (n.left).copy(), right: (n.right).copy()}
}

func (t *BST[T]) IsSymmetric() bool {
	if t.root == nil {
		return true
	}
	return t.root.symmetric(t.root)
}

func (n *Node[T]) symmetric(other *Node[T]) bool {
	if n == nil || other == nil {
		return n == nil && other == nil
	}
	if n.Value != other.Value {
		return false
	}
	return (n.left).symmetric(other.right) && (n.right).symmetric(other.left)
}
