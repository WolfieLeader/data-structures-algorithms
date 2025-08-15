package singly

type SinglyLinkedList[T comparable] interface {
	Size() int
	IsEmpty() bool
	Clear()
	Copy() SinglyLinkedList[T]

	AddFirst(values ...T)
	AddLast(values ...T)
	GetFirst() (T, bool)
	GetLast() (T, bool)
	RemoveFirst() (T, bool)
	RemoveLast() (T, bool)

	SetAt(position int, value T) bool
	SetAtNode(node *Node[T], value T) bool
	InsertAfter(position int, value T) bool
	InsertAfterNode(node *Node[T], value T) bool

	RemoveAt(position int) (T, bool)
	RemoveAfter(position int) (T, bool)
	RemoveAtNode(node *Node[T]) (T, bool)
	RemoveAfterNode(node *Node[T]) (T, bool)
	RemoveValue(value T) (T, bool)

	Get(position int) (T, bool)
	GetAll() []T
	LinearSearch(value T) (int, bool)
	Contains(value T) bool

	Traverse(func(position int, value T) bool)
	Reverse() SinglyLinkedList[T]
	IsSorted() bool
	Swap(i, j int) error
}

var _ SinglyLinkedList[int] = (*Singly[int])(nil)
