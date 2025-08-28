package bst

func (t BST[T]) Contains(value T) bool {
	cur := t.root
	for cur != nil {
		switch {
		case value < cur.Value:
			cur = cur.left
		case value > cur.Value:
			cur = cur.right
		default: // Equal
			return true
		}
	}
	return false
}

func (t BST[T]) RecursiveContains(value T) bool {
	return t.root.nodeContains(value)
}

func (n *Node[T]) nodeContains(value T) bool {
	if n == nil {
		return false
	}

	switch {
	case value < n.Value:
		return (n.left).nodeContains(value)
	case value > n.Value:
		return (n.right).nodeContains(value)
	default: // Equal
		return true
	}
}

func (t *BST[T]) Insert(value T) bool {
	if t.root == nil {
		t.root = &Node[T]{Value: value}
		t.size = 1
		return true
	}

	cur := t.root
	for {
		switch {
		case value < cur.Value:
			if cur.left == nil {
				cur.left = &Node[T]{Value: value}
				t.size++
				return true
			}
			cur = cur.left

		case value > cur.Value:
			if cur.right == nil {
				cur.right = &Node[T]{Value: value}
				t.size++
				return true
			}
			cur = cur.right

		default: // Equal
			return false
		}
	}
}

func (t *BST[T]) RecursiveInsert(value T) bool {
	newRoot, ok := t.root.nodeInsert(value)
	t.root = newRoot
	if ok {
		t.size++
	}
	return ok
}

func (n *Node[T]) nodeInsert(value T) (*Node[T], bool) {
	if n == nil {
		return &Node[T]{Value: value}, true
	}

	switch {
	case value < n.Value:
		newLeft, ok := (n.left).nodeInsert(value)
		n.left = newLeft
		return n, ok

	case value > n.Value:
		newRight, ok := (n.right).nodeInsert(value)
		n.right = newRight
		return n, ok

	default: // Equal
		return n, false
	}
}

func (t *BST[T]) Delete(value T) bool {
	var parent *Node[T]
	cur := t.root
	for cur != nil && cur.Value != value {
		parent = cur
		if value < cur.Value {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}

	if cur == nil {
		return false
	}

	if cur.left != nil && cur.right != nil {
		succParent, succ := cur, cur.right
		for succ.left != nil {
			succParent = succ
			succ = succ.left
		}

		cur.Value = succ.Value
		parent, cur = succParent, succ
	}

	var child *Node[T]
	if cur.left != nil {
		child = cur.left
	} else {
		child = cur.right
	}

	switch {
	case parent == nil:
		t.root = child
	case parent.left == cur:
		parent.left = child
	default:
		parent.right = child
	}

	t.size--
	return true
}

func (t *BST[T]) RecursiveDelete(value T) bool {
	newRoot, ok := t.root.nodeDelete(value)
	t.root = newRoot
	if ok {
		t.size--
	}
	return ok
}

func (n *Node[T]) nodeDelete(value T) (*Node[T], bool) {
	if n == nil {
		return nil, false
	}

	switch {
	case value < n.Value:
		newLeft, ok := (n.left).nodeDelete(value)
		n.left = newLeft
		return n, ok

	case value > n.Value:
		newRight, ok := (n.right).nodeDelete(value)
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
		newRight, ok := (n.right).nodeDelete(succ.Value)
		n.right = newRight
		return n, ok
	}
}

func (t BST[T]) Size() int {
	return t.size
}

func (t BST[T]) IsEmpty() bool {
	return t.size == 0
}

func (t BST[T]) Root() (T, bool) {
	var zero T
	if t.root == nil {
		return zero, false
	}
	return t.root.Value, true
}

func (t BST[T]) Min() (T, bool) {
	var zero T
	if t.root == nil {
		return zero, false
	}
	cur := t.root
	for cur.left != nil {
		cur = cur.left
	}
	return cur.Value, true
}

func (t BST[T]) Max() (T, bool) {
	var zero T
	if t.root == nil {
		return zero, false
	}
	cur := t.root
	for cur.right != nil {
		cur = cur.right
	}
	return cur.Value, true
}
