package doubly

import "cmp"

type DoublyLinkedList[T cmp.Ordered] interface {
	Size() int
	IsEmpty() bool
	Clear()
	Copy() *doublyLinkedList[T]

	AddFirst(values ...T)
	AddLast(values ...T)
	GetFirst() (T, bool)
	GetLast() (T, bool)
	RemoveFirst() (T, bool)
	RemoveLast() (T, bool)

	SetAt(i int, value T) bool
	SetAtNode(node *Node[T], value T) bool
	InsertAfter(i int, value T) bool
	InsertAfterNode(node *Node[T], value T) bool

	RemoveAt(i int) (T, bool)
	RemoveAfter(i int) (T, bool)
	RemoveAtNode(node *Node[T]) (T, bool)
	RemoveAfterNode(node *Node[T]) (T, bool)
	RemoveValue(value T) bool

	Get(i int) (T, bool)
	GetAll() []T
	LinearSearch(value T) int
	Contains(value T) bool

	Traverse(func(i int, value T) bool)
	Reverse() *doublyLinkedList[T]
	IsSorted() bool
	Swap(i, j int) error
}

var _ DoublyLinkedList[int] = (*doublyLinkedList[int])(nil)
