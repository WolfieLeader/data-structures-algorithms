package listqueue

func (q *ListQueue[T]) Enqueue(values ...T) { q.data.AddLast(values...) }
func (q *ListQueue[T]) Dequeue() (T, bool)  { return q.data.RemoveFirst() }
func (q ListQueue[T]) Peek() (T, bool)      { return q.data.GetFirst() }
func (q ListQueue[T]) Size() int            { return q.data.Size() }
func (q ListQueue[T]) IsEmpty() bool        { return q.data.IsEmpty() }
func (q *ListQueue[T]) Clear()              { q.data.Clear() }
func (q ListQueue[T]) ToSlice() []T         { return q.data.GetAll() }
func (q ListQueue[T]) Copy() *ListQueue[T]  { return &ListQueue[T]{data: q.data.Copy()} }
