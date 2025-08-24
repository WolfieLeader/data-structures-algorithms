package doubly

import "cmp"

type DoublyLinkedList[T cmp.Ordered] interface {
	Size() int
	IsEmpty() bool
	Clear()
	Copy() *Doubly[T]

	AddFirst(values ...T)
	AddLast(values ...T)
	GetFirst() (T, bool)
	GetLast() (T, bool)
	DeleteFirst() (T, bool)
	DeleteLast() (T, bool)

	SetAt(index int, value T) bool
	SetAtNode(node *Node[T], value T) bool
	InsertAt(index int, value T) bool
	InsertAtNode(node *Node[T], value T) bool
	InsertAfter(index int, value T) bool
	InsertAfterNode(node *Node[T], value T) bool

	DeleteAt(index int) (T, bool)
	DeleteAfter(index int) (T, bool)
	DeleteAtNode(node *Node[T]) (T, bool)
	DeleteAfterNode(node *Node[T]) (T, bool)
	DeleteValue(value T) bool

	Get(index int) (T, bool)
	ToSlice() []T
	Search(value T) int
	Contains(value T) bool

	Traverse(fn func(index int, value T))
	Reverse() *Doubly[T]
	IsSorted() bool
	Swap(index1, index2 int) error
	String() string
}

var _ DoublyLinkedList[int] = (*Doubly[int])(nil)
