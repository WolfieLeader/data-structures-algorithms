package arrqueue

func (q *ArrQueue[T]) Enqueue(values ...T) { q.data.Append(values...) }
func (q *ArrQueue[T]) Dequeue() (T, bool)  { return q.data.Delete(0) }
func (q ArrQueue[T]) Peek() (T, bool)      { return q.data.Get(0) }
func (q ArrQueue[T]) Size() int            { return q.data.Length() }
func (q ArrQueue[T]) IsEmpty() bool        { return q.data.Length() == 0 }
func (q *ArrQueue[T]) Clear()              { q.data.Clear() }
func (q ArrQueue[T]) ToSlice() []T         { return q.data.Copy() }
func (q ArrQueue[T]) Copy() *ArrQueue[T]   { return &ArrQueue[T]{data: q.data.Copy()} }
