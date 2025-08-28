package bst

func (t BST[T]) Contains(value T) bool {
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

	curr := t.root
	for {
		switch {
		case value < curr.Value:
			if curr.left == nil {
				curr.left = &Node[T]{Value: value}
				t.size++
				return true
			}
			curr = curr.left

		case value > curr.Value:
			if curr.right == nil {
				curr.right = &Node[T]{Value: value}
				t.size++
				return true
			}
			curr = curr.right

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
	curr, ptr := t.root, &t.root
	for curr != nil && curr.Value != value {
		if value < curr.Value {
			ptr = &curr.left
			curr = curr.left
		} else {
			ptr = &curr.right
			curr = curr.right
		}
	}

	if curr == nil {
		return false
	}

	if curr.left != nil && curr.right != nil {
		succ, succPtr := curr.right, &curr.right
		for succ.left != nil {
			succPtr = &succ.left
			succ = succ.left
		}

		curr.Value = succ.Value
		curr, ptr = succ, succPtr
	}

	if curr.left != nil {
		*ptr = curr.left
	} else {
		*ptr = curr.right
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
	curr := t.root
	for curr.left != nil {
		curr = curr.left
	}
	return curr.Value, true
}

func (t BST[T]) Max() (T, bool) {
	var zero T
	if t.root == nil {
		return zero, false
	}
	curr := t.root
	for curr.right != nil {
		curr = curr.right
	}
	return curr.Value, true
}
