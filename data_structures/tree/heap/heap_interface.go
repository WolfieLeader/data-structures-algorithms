package heap

type heap[T any] interface {
	Size() int
	IsEmpty() bool
	Clear()
	Peek() (T, bool)
	Pop() (T, bool)
	Push(values ...T)
	Contain(value T) bool
	Copy() *Heap[T]
	Equal(other *Heap[T]) bool
	String() string
}
