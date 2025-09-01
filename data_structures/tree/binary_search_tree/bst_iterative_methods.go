package bst

func (t *BST[T]) ContainsI(value T) bool {
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

func (t *BST[T]) InsertI(value T) bool {
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

func (t *BST[T]) DeleteI(value T) bool {
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

func (t *BST[T]) MinI() (T, bool) {
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

func (t *BST[T]) MaxI() (T, bool) {
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

func (t *BST[T]) TraverseInOrderI(fn func(value T)) {
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

func (t *BST[T]) TraversePreOrderI(fn func(value T)) {
	if t.root == nil {
		return
	}
	stack := newStack[*Node[T]](t.root)
	for !stack.IsEmpty() {
		n := stack.Pop()
		fn(n.Value)
		if n.right != nil {
			stack.Push(n.right)
		}
		if n.left != nil {
			stack.Push(n.left)
		}
	}
}

func (t *BST[T]) TraversePostOrderI(fn func(value T)) {
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

		n := stack.Peek()
		if n.right != nil && n.right != visited {
			curr = n.right
			continue
		}

		fn(n.Value)
		visited = n
		stack.Pop()
	}
}

func (t *BST[T]) TraverseBreadthFirst(fn func(value T)) {
	if t.root == nil {
		return
	}
	queue := newQueue[*Node[T]](t.root)
	for !queue.IsEmpty() {
		n := queue.Dequeue()
		fn(n.Value)
		if n.left != nil {
			queue.Enqueue(n.left)
		}
		if n.right != nil {
			queue.Enqueue(n.right)
		}
	}
}

func (t *BST[T]) HeightI() int {
	if t.root == nil {
		return 0
	}
	queue := newQueue[*Node[T]](t.root)
	height := 0
	for !queue.IsEmpty() {
		for range queue.Size() {
			n := queue.Dequeue()
			if n.left != nil {
				queue.Enqueue(n.left)
			}
			if n.right != nil {
				queue.Enqueue(n.right)
			}
		}
		height++
	}
	return height
}

func (t *BST[T]) CopyI() *BST[T] {
	if t.root == nil {
		return New[T]()
	}

	type pair struct {
		src *Node[T]
		dst *Node[T]
	}

	newTree := New[T]()
	newTree.size = t.size
	newTree.root = &Node[T]{Value: t.root.Value}

	stack := newStack[pair](pair{src: t.root, dst: newTree.root})
	for !stack.IsEmpty() {
		n := stack.Pop()
		if r := n.src.right; r != nil {
			n.dst.right = &Node[T]{Value: r.Value}
			stack.Push(pair{src: r, dst: n.dst.right})
		}
		if l := n.src.left; l != nil {
			n.dst.left = &Node[T]{Value: l.Value}
			stack.Push(pair{src: l, dst: n.dst.left})
		}
	}
	return newTree
}
