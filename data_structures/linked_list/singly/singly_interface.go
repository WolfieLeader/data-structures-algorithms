package singly

import "cmp"

type SinglyLinkedList[T cmp.Ordered] interface {
	Size() int
	IsEmpty() bool
	Clear()
	Copy() *Singly[T]

	AddFirst(values ...T)
	AddLast(values ...T)
	GetFirst() (T, bool)
	GetLast() (T, bool)
	DeleteFirst() (T, bool)
	DeleteLast() (T, bool)

	SetAt(i int, value T) bool
	SetAtNode(node *Node[T], value T) bool
	InsertAfter(i int, value T) bool
	InsertAfterNode(node *Node[T], value T) bool

	DeleteAt(i int) (T, bool)
	DeleteAfter(i int) (T, bool)
	DeleteAtNode(node *Node[T]) (T, bool)
	DeleteAfterNode(node *Node[T]) (T, bool)
	DeleteValue(value T) bool

	Get(i int) (T, bool)
	ToSlice() []T
	Search(value T) int
	Contains(value T) bool

	ForEach(func(i int, value T) bool)
	Reverse() *Singly[T]
	IsSorted() bool
	Swap(i, j int) error
}

var _ SinglyLinkedList[int] = (*Singly[int])(nil)
