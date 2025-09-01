package bst

func (t *BST[T]) Size() int     { return t.size }
func (t *BST[T]) IsEmpty() bool { return t.size == 0 }
func (t *BST[T]) Clear()        { *t = BST[T]{} }
func (t *BST[T]) Root() (T, bool) {
	var zero T
	if t.root == nil {
		return zero, false
	}
	return t.root.Value, true
}
