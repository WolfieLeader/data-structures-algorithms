package bst

func (t BST[T]) IterContains(value T) bool {
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

func (t BST[T]) RecContains(value T) bool {
	return t.root.contains(value)
}

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

func (t *BST[T]) IterInsert(value T) bool {
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

func (t *BST[T]) RecInsert(value T) bool {
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

func (t *BST[T]) IterDelete(value T) bool {
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

func (t *BST[T]) RecDelete(value T) bool {
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

func (t BST[T]) RecTraverseInOrder(fn func(value T))   { t.root.inOrder(fn) }
func (t BST[T]) RecTraversePreOrder(fn func(value T))  { t.root.preOrder(fn) }
func (t BST[T]) RecTraversePostOrder(fn func(value T)) { t.root.postOrder(fn) }

func (n *Node[T]) inOrder(fn func(value T)) {
	if n == nil {
		return
	}
	(n.left).inOrder(fn)
	fn(n.Value)
	(n.right).inOrder(fn)
}

func (n *Node[T]) preOrder(fn func(value T)) {
	if n == nil {
		return
	}
	fn(n.Value)
	(n.left).preOrder(fn)
	(n.right).preOrder(fn)
}

func (n *Node[T]) postOrder(fn func(value T)) {
	if n == nil {
		return
	}
	(n.left).postOrder(fn)
	(n.right).postOrder(fn)
	fn(n.Value)
}

func (t BST[T]) IterTraverseInOrder(fn func(value T)) {
	if t.root == nil {
		return
	}
	stack := newStack[*Node[T]]()
	curr := t.root
	for curr != nil || !stack.IsEmpty() {
		for curr != nil {
			stack.Push(curr)
			curr = curr.left
		}

		curr = stack.Pop()
		fn(curr.Value)
		curr = curr.right
	}
}

func (t BST[T]) IterTraversePreOrder(fn func(value T)) {
	if t.root == nil {
		return
	}
	stack := newStack[*Node[T]](t.root)
	for !stack.IsEmpty() {
		curr := stack.Pop()
		fn(curr.Value)
		if curr.right != nil {
			stack.Push(curr.right)
		}
		if curr.left != nil {
			stack.Push(curr.left)
		}
	}
}

func (t BST[T]) IterTraversePostOrder(fn func(value T)) {
	if t.root == nil {
		return
	}
	var visited *Node[T]
	stack := newStack[*Node[T]]()
	curr := t.root
	for curr != nil || !stack.IsEmpty() {
		if curr != nil {
			stack.Push(curr)
			curr = curr.left
			continue
		}

		top := stack.Peek()
		if top.right != nil && top.right != visited {
			curr = top.right
			continue
		}

		fn(top.Value)
		visited = top
		stack.Pop()
	}
}
