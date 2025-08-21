package listdeque

func (d *LinkedListDeque[T]) EnqueueFirst(values ...T) { d.data.AddFirst(values...) }
func (d *LinkedListDeque[T]) EnqueueLast(values ...T)  { d.data.AddLast(values...) }
func (d *LinkedListDeque[T]) DequeueFirst() (T, bool)  { return d.data.RemoveFirst() }
func (d *LinkedListDeque[T]) DequeueLast() (T, bool)   { return d.data.RemoveLast() }
func (d LinkedListDeque[T]) PeekFirst() (T, bool)      { return d.data.GetFirst() }
func (d LinkedListDeque[T]) PeekLast() (T, bool)       { return d.data.GetLast() }
func (d LinkedListDeque[T]) Size() int                 { return d.data.Size() }
func (d LinkedListDeque[T]) IsEmpty() bool             { return d.data.IsEmpty() }
func (d *LinkedListDeque[T]) Clear()                   { d.data.Clear() }
func (d LinkedListDeque[T]) ToSlice() []T              { return d.data.ToSlice() }
func (d LinkedListDeque[T]) Copy() *LinkedListDeque[T] {
	return &LinkedListDeque[T]{data: d.data.Copy()}
}
