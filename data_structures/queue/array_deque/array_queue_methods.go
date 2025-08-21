package arraydeque

func (d *ArrayDeque[T]) EnqueueFirst(values ...T) { d.data.Prepend(values...) }
func (d *ArrayDeque[T]) EnqueueLast(values ...T)  { d.data.Append(values...) }
func (d *ArrayDeque[T]) DequeueFirst() (T, bool)  { return d.data.Delete(0) }
func (d *ArrayDeque[T]) DequeueLast() (T, bool)   { return d.data.Delete(d.data.Length() - 1) }
func (d ArrayDeque[T]) PeekFirst() (T, bool)      { return d.data.Get(0) }
func (d ArrayDeque[T]) PeekLast() (T, bool)       { return d.data.Get(d.data.Length() - 1) }
func (d ArrayDeque[T]) Size() int                 { return d.data.Length() }
func (d ArrayDeque[T]) IsEmpty() bool             { return d.data.Length() == 0 }
func (d *ArrayDeque[T]) Clear()                   { d.data.Clear() }
func (d ArrayDeque[T]) ToSlice() []T              { return d.data.Copy() }
func (d ArrayDeque[T]) Copy() *ArrayDeque[T]      { return &ArrayDeque[T]{data: d.data.Copy()} }
