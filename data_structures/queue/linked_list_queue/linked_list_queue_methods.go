package listqueue

func (q *LinkedListQueue[T]) Enqueue(values ...T) { q.data.AddLast(values...) }
func (q *LinkedListQueue[T]) Dequeue() (T, bool)  { return q.data.RemoveFirst() }
func (q LinkedListQueue[T]) Peek() (T, bool)      { return q.data.GetFirst() }
func (q LinkedListQueue[T]) Size() int            { return q.data.Size() }
func (q LinkedListQueue[T]) IsEmpty() bool        { return q.data.IsEmpty() }
func (q *LinkedListQueue[T]) Clear()              { q.data.Clear() }
func (q LinkedListQueue[T]) ToSlice() []T         { return q.data.ToSlice() }
func (q LinkedListQueue[T]) Copy() *LinkedListQueue[T] {
	return &LinkedListQueue[T]{data: q.data.Copy()}
}
