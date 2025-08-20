package arraydeque

func (q *ArrayDeque[T]) EnqueueFirst(values ...T) { q.data.Prepend(values...) }
func (q *ArrayDeque[T]) EnqueueLast(values ...T)  { q.data.Append(values...) }
func (q *ArrayDeque[T]) DequeueFirst() (T, bool)  { return q.data.Delete(0) }
func (q *ArrayDeque[T]) DequeueLast() (T, bool)   { return q.data.Delete(q.data.Length() - 1) }
func (q ArrayDeque[T]) PeekFirst() (T, bool)      { return q.data.Get(0) }
func (q ArrayDeque[T]) PeekLast() (T, bool)       { return q.data.Get(q.data.Length() - 1) }
func (q ArrayDeque[T]) Size() int                 { return q.data.Length() }
func (q ArrayDeque[T]) IsEmpty() bool             { return q.data.Length() == 0 }
func (q *ArrayDeque[T]) Clear()                   { q.data.Clear() }
func (q ArrayDeque[T]) ToSlice() []T              { return q.data.Copy() }
func (q ArrayDeque[T]) Copy() *ArrayDeque[T]      { return &ArrayDeque[T]{data: q.data.Copy()} }
