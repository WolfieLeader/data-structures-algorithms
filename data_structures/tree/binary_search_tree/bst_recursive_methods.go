package bst

import (
	"fmt"
	"math"
	"strings"
)

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

func (t *BST[T]) InsertR(values ...T) {
	for _, v := range values {
		t.insertR(v)
	}
}

func (t *BST[T]) insertR(value T) {
	newRoot, ok := t.root.insert(value)
	t.root = newRoot
	if ok {
		t.size++
	}
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

func (t *BST[T]) SuccessorR(value T) (T, bool) { return t.root.successor(value, nil) }
func (n *Node[T]) successor(value T, succ *Node[T]) (T, bool) {
	var zero T
	if n == nil {
		if succ != nil {
			return succ.Value, true
		}
		return zero, false
	}

	switch {
	case value < n.Value:
		return (n.left).successor(value, n)
	case value > n.Value:
		return (n.right).successor(value, succ)
	default: // Equal
		if r := n.right; r != nil {
			for r.left != nil {
				r = r.left
			}
			return r.Value, true
		}
		if succ != nil {
			return succ.Value, true
		}
		return zero, false
	}
}

func (t *BST[T]) PredecessorR(value T) (T, bool) { return t.root.predecessor(value, nil) }
func (n *Node[T]) predecessor(value T, pred *Node[T]) (T, bool) {
	var zero T
	if n == nil {
		if pred != nil {
			return pred.Value, true
		}
		return zero, false
	}

	switch {
	case value < n.Value:
		return (n.left).predecessor(value, pred)
	case value > n.Value:
		return (n.right).predecessor(value, n)
	default: // Equal
		if l := n.left; l != nil {
			for l.right != nil {
				l = l.right
			}
			return l.Value, true
		}
		if pred != nil {
			return pred.Value, true
		}
		return zero, false
	}
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

// Cannot be used in BST only in Binary Tree
// func (t *BST[T]) IsSymmetric() bool {
// 	if t.root == nil {
// 		return true
// 	}
// 	return t.root.left.symmetric(t.root.right)
// }

// func (n *Node[T]) symmetric(other *Node[T]) bool {
// 	if n == nil || other == nil {
// 		return n == nil && other == nil
// 	}
// 	if n.Value != other.Value {
// 		return false
// 	}
// 	return (n.left).symmetric(other.right) && (n.right).symmetric(other.left)
// }

func (t *BST[T]) String() string {
	var sb strings.Builder
	if t.root == nil {
		return "BST{size=0}\n"
	}
	fmt.Fprintf(&sb, "BST{size=%d}\n", t.size)
	t.root.draw(&sb, "", 0, "", true)
	return sb.String()
}

func (n *Node[T]) draw(sb *strings.Builder, prefix string, level int, label string, isTail bool) {
	if n == nil {
		return
	}

	// current line with branch
	sb.WriteString(prefix)
	switch {
	case prefix == "":
		sb.WriteString("└── ")
	case isTail:
		sb.WriteString("└── ")
	default:
		sb.WriteString("├── ")
	}

	if prefix == "" {
		fmt.Fprintf(sb, "%v\n", n.Value)
	} else {
		fmt.Fprintf(sb, "%v(%s%d)\n", n.Value, label, level)
	}

	type labeled struct {
		node  *Node[T]
		label string
	}

	children := make([]labeled, 0, 2)
	if n.right != nil {
		children = append(children, labeled{n.right, "R"})
	}
	if n.left != nil {
		children = append(children, labeled{n.left, "L"})
	}

	for i, child := range children {
		var nextPrefix string
		switch {
		case prefix == "":
			nextPrefix = "    "
		case isTail:
			nextPrefix = prefix + "    "
		default:
			nextPrefix = prefix + "│   "
		}
		child.node.draw(sb, nextPrefix, level+1, child.label, i == len(children)-1)
	}
}
