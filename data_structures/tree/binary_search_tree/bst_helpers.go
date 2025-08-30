package bst

type stack[T any] struct {
	data []T
}

func newStack[T any](values ...T) *stack[T] {
	return &stack[T]{data: values}
}

func (s *stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

func (s *stack[T]) Pop() T {
	var zero T
	if len(s.data) == 0 {
		return zero
	}
	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value
}

func (s *stack[T]) Peek() T {
	var zero T
	if len(s.data) == 0 {
		return zero
	}
	return s.data[len(s.data)-1]
}

func (s stack[T]) Size() int {
	return len(s.data)
}

func (s stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}
