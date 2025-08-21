package arrayqueue

func (q *ArrayQueue[T]) Enqueue(values ...T) { q.data.Append(values...) }
func (q *ArrayQueue[T]) Dequeue() (T, bool)  { return q.data.Delete(0) }
func (q ArrayQueue[T]) Peek() (T, bool)      { return q.data.Get(0) }
func (q ArrayQueue[T]) Size() int            { return q.data.Length() }
func (q ArrayQueue[T]) IsEmpty() bool        { return q.data.IsEmpty() }
func (q *ArrayQueue[T]) Clear()              { q.data.Clear() }
func (q ArrayQueue[T]) ToSlice() []T         { return q.data.ToSlice() }
func (q ArrayQueue[T]) Copy() *ArrayQueue[T] { return &ArrayQueue[T]{data: q.data.Copy()} }
