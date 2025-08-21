package arraystack

func (s *ArrayStack[T]) Push(values ...T)     { s.data.Append(values...) }
func (s *ArrayStack[T]) Pop() (T, bool)       { return s.data.Delete(s.data.Length() - 1) }
func (s *ArrayStack[T]) Peek() (T, bool)      { return s.data.Get(s.data.Length() - 1) }
func (s *ArrayStack[T]) Size() int            { return s.data.Length() }
func (s *ArrayStack[T]) IsEmpty() bool        { return s.data.IsEmpty() }
func (s *ArrayStack[T]) Clear()               { s.data.Clear() }
func (s *ArrayStack[T]) ToSlice() []T         { return s.data.Reverse().ToSlice() }
func (s *ArrayStack[T]) Copy() *ArrayStack[T] { return &ArrayStack[T]{data: s.data.Copy()} }
