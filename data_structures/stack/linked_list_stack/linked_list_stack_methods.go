package liststack

func (s *LinkedListStack[T]) Push(values ...T) { s.data.AddFirst(values...) }
func (s *LinkedListStack[T]) Pop() (T, bool)   { return s.data.DeleteFirst() }
func (s *LinkedListStack[T]) Peek() (T, bool)  { return s.data.GetFirst() }
func (s *LinkedListStack[T]) Size() int        { return s.data.Size() }
func (s *LinkedListStack[T]) IsEmpty() bool    { return s.data.IsEmpty() }
func (s *LinkedListStack[T]) Clear()           { s.data.Clear() }
func (s *LinkedListStack[T]) ToSlice() []T     { return s.data.ToSlice() }
func (s *LinkedListStack[T]) Copy() *LinkedListStack[T] {
	return &LinkedListStack[T]{data: s.data.Copy()}
}
