package liststack

func (s *ListStack[T]) Push(values ...T)    { s.data.AddFirst(values...) }
func (s *ListStack[T]) Pop() (T, bool)      { return s.data.RemoveFirst() }
func (s *ListStack[T]) Peek() (T, bool)     { return s.data.GetFirst() }
func (s *ListStack[T]) Size() int           { return s.data.Size() }
func (s *ListStack[T]) IsEmpty() bool       { return s.data.IsEmpty() }
func (s *ListStack[T]) Clear()              { s.data.Clear() }
func (s *ListStack[T]) ToSlice() []T        { return s.data.GetAll() }
func (s *ListStack[T]) Copy() *ListStack[T] { return &ListStack[T]{data: s.data.Copy()} }
