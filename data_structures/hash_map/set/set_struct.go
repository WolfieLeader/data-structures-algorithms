package hashset

type Set[T comparable] struct {
	data map[T]struct{}
}

func New[T comparable]() *Set[T] {
	return &Set[T]{data: make(map[T]struct{})}
}
