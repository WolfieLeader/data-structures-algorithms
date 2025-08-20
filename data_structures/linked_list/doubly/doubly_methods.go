package doubly

import (
	"fmt"

	"github.com/WolfieLeader/data-structures-algorithms/utils"
)

func (list Doubly[T]) Size() int {
	return list.size
}

func (list Doubly[T]) IsEmpty() bool {
	return list.size == 0
}

func (list *Doubly[T]) Clear() {
	list.size = 0
	list.head = nil
	list.tail = nil
}

func (list Doubly[T]) Copy() *Doubly[T] {
	if list.size == 0 {
		return New[T]()
	}

	out := New[T]()
	list.Traverse(func(i int, value T) bool {
		out.AddLast(value)
		return true
	})
	return out
}

func (list *Doubly[T]) AddFirst(values ...T) {
	for _, value := range values {
		n := &Node[T]{Value: value, next: list.head}

		if list.head == nil {
			list.tail = n
		} else {
			list.head.prev = n
		}

		list.head = n
		list.size++
	}
}

func (list *Doubly[T]) AddLast(values ...T) {
	for _, value := range values {
		n := &Node[T]{Value: value, prev: list.tail}

		if list.tail == nil {
			list.head, list.tail = n, n
			list.size++
			continue
		}

		list.tail.next = n
		list.tail = n
		list.size++
	}
}

func (list Doubly[T]) GetFirst() (T, bool) {
	var zero T
	if list.head == nil {
		return zero, false
	}
	return list.head.Value, true
}

func (list Doubly[T]) GetLast() (T, bool) {
	var zero T
	if list.tail == nil {
		return zero, false
	}
	return list.tail.Value, true
}

func (list *Doubly[T]) RemoveFirst() (T, bool) {
	var zero T
	if list.head == nil {
		return zero, false
	}

	value := list.head.Value
	list.head = list.head.next

	if list.head == nil {
		list.tail = nil
	} else {
		list.head.prev = nil
	}

	list.size--
	return value, true
}

func (list *Doubly[T]) RemoveLast() (T, bool) {
	var zero T
	if list.tail == nil {
		return zero, false
	}

	value := list.tail.Value
	list.tail = list.tail.prev

	if list.tail == nil {
		list.head = nil
	} else {
		list.tail.next = nil
	}

	list.size--
	return value, true
}

func (list Doubly[T]) shortestPath(i int) *Node[T] {
	var curr *Node[T]
	if i <= (list.size-1)/2 {
		curr = list.head
		for curr != nil && i > 0 {
			curr = curr.next
			i--
		}
		return curr
	}

	curr = list.tail
	for curr != nil && i < list.size-1 {
		curr = curr.prev
		i++
	}
	return curr
}

func (list *Doubly[T]) SetAt(i int, value T) bool {
	if i < 0 || i >= list.size {
		return false
	}

	curr := list.shortestPath(i)
	if curr == nil {
		return false
	}

	curr.Value = value
	return true
}

func (list *Doubly[T]) SetAtNode(node *Node[T], value T) bool {
	if node == nil {
		return false
	}

	curr := list.head
	for curr != nil && curr != node {
		curr = curr.next
	}

	if curr == nil {
		return false
	}

	curr.Value = value
	return true
}

func (list *Doubly[T]) InsertAfter(i int, value T) bool {
	if i < 0 || i >= list.size {
		return false
	}

	if i == list.size-1 {
		list.AddLast(value)
		return true
	}

	curr := list.shortestPath(i)
	if curr == nil {
		return false
	}

	n := &Node[T]{Value: value, prev: curr, next: curr.next}
	curr.next = n

	if n.next == nil {
		list.tail = n
	} else {
		n.next.prev = n
	}

	list.size++
	return true
}

func (list *Doubly[T]) InsertAfterNode(node *Node[T], value T) bool {
	if node == nil {
		return false
	}

	curr := list.head
	for curr != nil && curr != node {
		curr = curr.next
	}

	if curr == nil {
		return false
	}

	n := &Node[T]{Value: value, prev: curr, next: curr.next}
	curr.next = n

	if n.next == nil {
		list.tail = n
	} else {
		n.next.prev = n
	}

	list.size++
	return true
}

func (list *Doubly[T]) unlink(curr *Node[T]) (T, bool) {
	var zero T
	if curr == nil {
		return zero, false
	}

	if curr.prev == nil {
		list.head = curr.next
	} else {
		curr.prev.next = curr.next
	}

	if curr.next == nil {
		list.tail = curr.prev
	} else {
		curr.next.prev = curr.prev
	}

	list.size--
	return curr.Value, true
}

func (list *Doubly[T]) RemoveAt(i int) (T, bool) {
	var zero T
	if i < 0 || i >= list.size {
		return zero, false
	}

	if i == 0 {
		return list.RemoveFirst()
	}
	if i == list.size-1 {
		return list.RemoveLast()
	}

	curr := list.shortestPath(i)
	return list.unlink(curr)
}

func (list *Doubly[T]) RemoveAfter(i int) (T, bool) {
	var zero T
	if i < 0 || i >= list.size {
		return zero, false
	}

	curr := list.shortestPath(i)
	if curr == nil {
		return zero, false
	}

	return list.unlink(curr.next)
}

func (list *Doubly[T]) RemoveAtNode(node *Node[T]) (T, bool) {
	var zero T
	if node == nil {
		return zero, false
	}

	if list.head == node {
		return list.RemoveFirst()
	}
	if list.tail == node {
		return list.RemoveLast()
	}

	curr := list.head
	for curr != nil && curr != node {
		curr = curr.next
	}

	return list.unlink(node)
}

func (list *Doubly[T]) RemoveAfterNode(node *Node[T]) (T, bool) {
	var zero T

	if node == nil {
		return zero, false
	}

	curr := list.head
	for curr != nil && curr != node {
		curr = curr.next
	}

	if curr == nil {
		return zero, false
	}

	return list.unlink(curr.next)
}

func (list *Doubly[T]) RemoveValue(value T) bool {
	if list.head == nil {
		return false
	}

	if utils.Is(list.head.Value, utils.EqualTo, value) {
		list.RemoveFirst()
		return true
	}
	if utils.Is(list.tail.Value, utils.EqualTo, value) {
		list.RemoveLast()
		return true
	}

	curr := list.head
	for curr != nil {
		if utils.Is(curr.Value, utils.EqualTo, value) {
			list.unlink(curr)
			return true
		}

		curr = curr.next
	}
	return false
}

func (list Doubly[T]) Get(i int) (T, bool) {
	var zero T
	if i < 0 || i >= list.size {
		return zero, false
	}

	curr := list.shortestPath(i)
	if curr == nil {
		return zero, false
	}

	return curr.Value, true
}

func (list Doubly[T]) ToSlice() []T {
	out := make([]T, 0, list.size)

	list.Traverse(func(i int, value T) bool {
		out = append(out, value)
		return true
	})

	return out
}

func (list Doubly[T]) Search(value T) int {
	i := 0
	for curr := list.head; curr != nil; curr = curr.next {
		if utils.Is(curr.Value, utils.EqualTo, value) {
			return i
		}
		i++
	}
	return -1
}

func (list Doubly[T]) Contains(value T) bool {
	return list.Search(value) != -1
}

func (list Doubly[T]) Traverse(fn func(i int, value T) bool) {
	i := 0
	for curr := list.head; curr != nil; curr = curr.next {
		if !fn(i, curr.Value) {
			break
		}
		i++
	}
}

func (list Doubly[T]) Reverse() *Doubly[T] {
	out := list.Copy()
	curr := out.head
	for curr != nil {
		curr.prev, curr.next = curr.next, curr.prev
		curr = curr.prev
	}
	out.head, out.tail = out.tail, out.head
	return out
}

func (list Doubly[T]) IsSorted() bool {
	if list.size <= 1 {
		return true
	}

	curr := list.head
	for curr != nil && curr.next != nil {
		if utils.Is(curr.Value, utils.GreaterThan, curr.next.Value) {
			return false
		}
		curr = curr.next
	}
	return true
}

func (list *Doubly[T]) Swap(i, j int) error {
	if i < 0 || j < 0 || i >= list.size || j >= list.size {
		return fmt.Errorf("index out of bounds: i=%d, j=%d", i, j)
	}

	if i == j {
		return nil
	}

	if i > j {
		i, j = j, i
	}

	index := 0
	var n1, n2 *Node[T]
	for curr := list.head; curr != nil; curr = curr.next {
		if index == i {
			n1 = curr
		}
		if index == j {
			n2 = curr
			break
		}

		index++

	}

	if n1 == nil || n2 == nil {
		return fmt.Errorf("nodes not found at i=%d, j=%d", i, j)
	}

	if n1.prev == nil {
		list.head = n2
		n2.prev = nil
	} else {
		n1.prev.next = n2
		n2.prev = n1.prev
	}

	if n1.next == n2 {
		n1.next = n2.next

		if n2.next == nil {
			list.tail = n1
		} else {
			n2.next.prev = n1
		}

		n2.next = n1
		n1.prev = n2

		return nil
	}

	n2.prev.next = n1
	n1.prev = n2.prev

	ogNext1, ogNext2 := n1.next, n2.next

	if ogNext1 == nil {
		list.tail = n2
	} else {
		ogNext1.prev = n2
	}

	if ogNext2 == nil {
		list.tail = n1
	} else {
		ogNext2.prev = n1
	}

	n1.next, n2.next = ogNext2, ogNext1

	return nil
}
