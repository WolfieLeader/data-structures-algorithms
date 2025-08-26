package bst

func (t BST[T]) Contains(value T) bool {
	cur := t.root
	for cur != nil {
		if value == cur.Value {
			return true
		}
		if value < cur.Value {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	return false
}

func (t *BST[T]) Insert(value T) bool {
	if t.root == nil {
		t.root = &Node[T]{Value: value}
		t.size = 1
		return true
	}

	cur := t.root
	for {
		if value == cur.Value {
			return false
		}
		if value < cur.Value {
			if cur.left == nil {
				cur.left = &Node[T]{Value: value}
				t.size++
				return true
			}
			cur = cur.left
		} else {
			if cur.right == nil {
				cur.right = &Node[T]{Value: value}
				t.size++
				return true
			}
			cur = cur.right
		}
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

	if parent == nil {
		t.root = child
	} else if parent.left == cur {
		parent.left = child
	} else {
		parent.right = child
	}

	t.size--
	return true
}

func (t BST[T]) Size() int {
	return t.size
}

func (t BST[T]) IsEmpty() bool {
	return t.size == 0
}

func (t BST[T]) Root() T {
	if t.root == nil {
		var zero T
		return zero
	}
	return t.root.Value
}

func (t BST[T]) Min() T {
	if t.root == nil {
		var zero T
		return zero
	}
	cur := t.root
	for cur.left != nil {
		cur = cur.left
	}
	return cur.Value
}

func (t BST[T]) Max() T {
	if t.root == nil {
		var zero T
		return zero
	}
	cur := t.root
	for cur.right != nil {
		cur = cur.right
	}
	return cur.Value
}
