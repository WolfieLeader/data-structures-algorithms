package hashset

type HashSet[T comparable] interface {
	Add(values ...T)
	Delete(value T) bool
	Contains(value T) bool
	Size() int
	IsEmpty() bool
	Clear()
	ToSlice() []T
	ForEach(func(value T) bool)
	Copy() *Set[T]
}

var _ HashSet[string] = (*Set[string])(nil)
