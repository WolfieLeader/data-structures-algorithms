package listdeque

func (q *LinkedListDeque[T]) EnqueueFirst(values ...T) { q.data.AddFirst(values...) }
func (q *LinkedListDeque[T]) EnqueueLast(values ...T)  { q.data.AddLast(values...) }
func (q *LinkedListDeque[T]) DequeueFirst() (T, bool)  { return q.data.RemoveFirst() }
func (q *LinkedListDeque[T]) DequeueLast() (T, bool)   { return q.data.RemoveLast() }
func (q LinkedListDeque[T]) PeekFirst() (T, bool)      { return q.data.GetFirst() }
func (q LinkedListDeque[T]) PeekLast() (T, bool)       { return q.data.GetLast() }
func (q LinkedListDeque[T]) Size() int                 { return q.data.Size() }
func (q LinkedListDeque[T]) IsEmpty() bool             { return q.data.IsEmpty() }
func (q *LinkedListDeque[T]) Clear()                   { q.data.Clear() }
func (q LinkedListDeque[T]) ToSlice() []T              { return q.data.ToSlice() }
func (q LinkedListDeque[T]) Copy() *LinkedListDeque[T] {
	return &LinkedListDeque[T]{data: q.data.Copy()}
}
