package hashset

func (s *Set[T]) Add(values ...T) {
	for _, value := range values {
		s.data[value] = struct{}{}
	}
}

func (s *Set[T]) Delete(value T) bool {
	if _, ok := s.data[value]; ok {
		delete(s.data, value)
		return true
	}
	return false
}

func (s Set[T]) Contains(value T) bool {
	_, ok := s.data[value]
	return ok
}

func (s Set[T]) Size() int {
	return len(s.data)
}

func (s Set[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Set[T]) Clear() {
	s.data = make(map[T]struct{})
}

func (s Set[T]) ToSlice() []T {
	out := make([]T, 0, len(s.data))
	for k := range s.data {
		out = append(out, k)
	}
	return out
}

func (s Set[T]) ForEach(fn func(value T) bool) {
	for k := range s.data {
		if !fn(k) {
			return
		}
	}
}

func (s Set[T]) Copy() *Set[T] {
	out := New[T]()
	for k := range s.data {
		out.data[k] = struct{}{}
	}
	return out
}
