package arraystack

func (s *ArrStack[T]) Push(values ...T)   { s.data.Append(values...) }
func (s *ArrStack[T]) Pop() (T, bool)     { return s.data.Delete(s.data.Length() - 1) }
func (s *ArrStack[T]) Peek() (T, bool)    { return s.data.Get(s.data.Length() - 1) }
func (s *ArrStack[T]) Size() int          { return s.data.Length() }
func (s *ArrStack[T]) IsEmpty() bool      { return s.data.IsEmpty() }
func (s *ArrStack[T]) Clear()             { s.data.Clear() }
func (s *ArrStack[T]) ToSlice() []T       { return s.data.Reverse() }
func (s *ArrStack[T]) Copy() *ArrStack[T] { return &ArrStack[T]{data: s.data.Copy()} }
