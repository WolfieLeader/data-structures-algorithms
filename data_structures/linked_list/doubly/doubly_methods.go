package doubly

import (
	"fmt"

	"github.com/WolfieLeader/data-structures-algorithms/utils"
)

func (list doubly[T]) Size() int {
	return list.size
}

func (list doubly[T]) IsEmpty() bool {
	return list.head == nil
}

func (list *doubly[T]) Clear() {
	list.size = 0
	list.head = nil
	list.tail = nil
}

func (list doubly[T]) Copy() *doubly[T] {
	if list.size == 0 {
		return New[T]()
	}

	out := New[T]()
	curr := list.head
	for curr != nil {
		out.AddLast(curr.Value)
		curr = curr.next
	}
	return out
}

func (list *doubly[T]) AddFirst(values ...T) {
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

func (list *doubly[T]) AddLast(values ...T) {
	for _, value := range values {
		n := &Node[T]{Value: value, prev: list.tail}
		if list.tail == nil {
			list.head = n
		} else {
			list.tail.next = n
		}
		list.tail = n
		list.size++
	}
}

func (list doubly[T]) GetFirst() (T, bool) {
	if list.head == nil {
		var zero T
		return zero, false
	}
	return list.head.Value, true
}

func (list doubly[T]) GetLast() (T, bool) {
	if list.tail == nil {
		var zero T
		return zero, false
	}
	return list.tail.Value, true
}

func (list *doubly[T]) RemoveFirst() (T, bool) {
	if list.head == nil {
		var zero T
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

func (list *doubly[T]) RemoveLast() (T, bool) {
	if list.tail == nil {
		var zero T
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

func (list *doubly[T]) SetAt(i int, value T) bool {
	if i < 0 || i >= list.size {
		return false
	}

	curr := list.head
	for curr != nil && i > 0 {
		curr = curr.next
		i--
	}

	if curr == nil {
		return false
	}

	curr.Value = value
	return true
}

func (list *doubly[T]) SetAtNode(node *Node[T], value T) bool {
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

func (list *doubly[T]) InsertAfter(i int, value T) bool {
	if i < 0 || i >= list.size {
		return false
	}

	curr := list.head
	for curr != nil && i > 0 {
		curr = curr.next
		i--
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

func (list *doubly[T]) InsertAfterNode(node *Node[T], value T) bool {
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

func (list *doubly[T]) unlink(curr *Node[T]) (T, bool) {
	if curr == nil || list.head == nil || list.tail == nil {
		var zero T
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

func (list *doubly[T]) RemoveAt(i int) (T, bool) {
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

	curr := list.head
	for curr != nil && i > 0 {
		curr = curr.next
		i--
	}

	if curr == nil {
		return zero, false
	}

	return list.unlink(curr)
}

func (list *doubly[T]) RemoveAfter(i int) (T, bool) {
	var zero T

	if i < 0 || i >= list.size {
		return zero, false
	}

	curr := list.head
	for curr != nil && i > 0 {
		curr = curr.next
		i--
	}

	if curr == nil || curr.next == nil {
		return zero, false
	}

	return list.unlink(curr.next)
}

func (list *doubly[T]) RemoveAtNode(node *Node[T]) (T, bool) {
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
	for curr != nil && curr.next != node {
		curr = curr.next
	}

	return list.unlink(node)
}

func (list *doubly[T]) RemoveAfterNode(node *Node[T]) (T, bool) {
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

func (list *doubly[T]) RemoveValue(value T) bool {
	if list.head == nil {
		return false
	}

	if list.head.Value == value {
		_, found := list.RemoveFirst()
		return found
	}

	curr := list.head
	for curr != nil {
		if curr.Value == value {
			list.unlink(curr)
			return true
		}
		curr = curr.next
	}
	return false
}

func (list doubly[T]) Get(i int) (T, bool) {
	var zero T

	if i < 0 || i >= list.size {
		return zero, false
	}

	curr := list.head
	for curr != nil && i > 0 {
		curr = curr.next
		i--
	}

	if curr == nil {
		return zero, false
	}

	return curr.Value, true
}

func (list doubly[T]) GetAll() []T {
	out := make([]T, 0, list.size)
	curr := list.head
	for curr != nil {
		out = append(out, curr.Value)
		curr = curr.next
	}

	return out
}

func (list doubly[T]) Search(value T) int {
	i := 0
	curr := list.head
	for curr != nil && curr.Value != value {
		curr = curr.next
		i++
	}

	if curr == nil {
		return -1
	}

	return i
}

func (list doubly[T]) Contains(value T) bool {
	return list.Search(value) != -1
}

func (list doubly[T]) Traverse(fn func(i int, value T) bool) {
	i := 0
	curr := list.head
	for curr != nil {
		if !fn(i, curr.Value) {
			break
		}
		curr = curr.next
		i++
	}
}

func (list doubly[T]) Reverse() *doubly[T] {

	out := list.Copy()
	curr := out.head
	for curr != nil {
		curr.prev, curr.next = curr.next, curr.prev
		curr = curr.prev
	}
	out.head, out.tail = out.tail, out.head
	return out
}

func (list doubly[T]) IsSorted() bool {
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

func (list *doubly[T]) Swap(i, j int) error {
	if i < 0 || j < 0 || i >= list.size || j >= list.size {
		return fmt.Errorf("index out of bounds: i=%d, j=%d", i, j)
	}

	if i == j {
		return nil
	}

	if i > j {
		i, j = j, i
	}

	var n1, n2 *Node[T]

	index := 0
	curr := list.head

	for curr != nil {
		if index == i {
			n1 = curr
		}
		if index == j {
			n2 = curr
			break
		}

		index++
		curr = curr.next
	}

	if n1 == nil || n2 == nil {
		return fmt.Errorf("nodes not found at i=%d, j=%d", i, j)
	}

	prev1, prev2 := n1.prev, n2.prev
	next1, next2 := n1.next, n2.next

	if next1 != n2 {
		// reconnect neighbors to opposite nodes
		if prev1 != nil {
			prev1.next = n2
		} else {
			list.head = n2
		}
		if next1 != nil {
			next1.prev = n2
		}
		if prev2 != nil {
			prev2.next = n1
		} else {
			list.head = n1
		}
		if next2 != nil {
			next2.prev = n1
		}

		// swap n1 <-> n2 links
		n1.prev, n1.next, n2.prev, n2.next = prev2, next2, prev1, next1
	} else {
		// n1 <-> n2 are adjacent: prev1 -> n1 -> n2 -> nj
		if prev1 != nil {
			prev1.next = n2
		} else {
			list.head = n2
		}
		if next2 != nil {
			next2.prev = n1
		} else {
			list.tail = n1
		}

		n2.prev = prev1
		n1.next = next2
		n2.next = n1
		n1.prev = n2
	}

	// fix tail (for non-adjacent case we may need both checks)
	if n1.next == nil {
		list.tail = n1
	}
	if n2.next == nil {
		list.tail = n2
	}
	return nil
}
