package singly

type SinglyLinkedList[T comparable] interface {
	AddFirst(values ...T)
	AddLast(values ...T)
	SetAt(position int, value T) bool
	SetAtNode(node *Node[T], value T) bool
	InsertAfter(position int, value T) bool
	InsertAfterNode(node *Node[T], value T) bool
	RemoveFirst() (T, bool)
	RemoveLast() (T, bool)
	RemoveAt(position int) (T, bool)
	RemoveAtNode(node *Node[T]) (T, bool)
	RemoveValue(value T) (T, bool)
	Clear()
	GetFirst() (T, bool)
	GetLast() (T, bool)
	Get(position int) (T, bool)
	GetAll() []T
	Find(value T) (int, bool)
	IsEmpty() bool
	Size() int
	Traverse(func(value T))
	Reverse() SinglyLinkedList[T]
	IsSorted() bool
	Copy() SinglyLinkedList[T]
	Swap(i, j int) error
}

var _ SinglyLinkedList[int] = (*Singly[int])(nil)
