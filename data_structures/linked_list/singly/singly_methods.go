package singly

import (
	"fmt"

	"github.com/WolfieLeader/data-structures-algorithms/utils"
)

func (list singly[T]) Size() int {
	return list.size
}

func (list singly[T]) IsEmpty() bool {
	return list.head == nil
}

func (list *singly[T]) Clear() {
	list.size = 0
	list.head = nil
	list.tail = nil
}

func (list singly[T]) Copy() *singly[T] {
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

func (list *singly[T]) AddFirst(values ...T) {
	for _, value := range values {
		n := &Node[T]{Value: value, next: list.head}
		list.head = n
		if list.head == nil {
			list.tail = n
		}
		list.size++
	}
}

func (list *singly[T]) AddLast(values ...T) {
	for _, value := range values {
		n := &Node[T]{Value: value}
		if list.tail == nil {
			list.head = n
		} else {
			list.tail.next = n
		}
		list.tail = n
		list.size++
	}
}

func (list singly[T]) GetFirst() (T, bool) {
	if list.head == nil {
		var zero T
		return zero, false
	}
	return list.head.Value, true
}

func (list singly[T]) GetLast() (T, bool) {
	if list.tail == nil {
		var zero T
		return zero, false
	}
	return list.tail.Value, true
}

func (list *singly[T]) RemoveFirst() (T, bool) {
	if list.head == nil {
		var zero T
		return zero, false
	}

	value := list.head.Value
	list.head = list.head.next
	if list.head == nil {
		list.tail = nil
	}

	list.size--
	return value, true
}

func (list *singly[T]) RemoveLast() (T, bool) {
	if list.tail == nil {
		var zero T
		return zero, false
	}

	value := list.tail.Value
	if list.head == list.tail {
		list.head, list.tail = nil, nil
		list.size = 0
		return value, true
	}

	prev := list.head
	for prev.next != list.tail {
		prev = prev.next
	}

	prev.next = nil
	list.tail = prev

	list.size--
	return value, true
}

func (list *singly[T]) SetAt(i int, value T) bool {
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

func (list *singly[T]) SetAtNode(node *Node[T], value T) bool {
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

func (list *singly[T]) InsertAfter(i int, value T) bool {
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

	n := &Node[T]{Value: value, next: curr.next}
	curr.next = n

	if n.next == nil {
		list.tail = n
	}

	list.size++
	return true
}

func (list *singly[T]) InsertAfterNode(node *Node[T], value T) bool {
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

	n := &Node[T]{Value: value, next: curr.next}
	curr.next = n

	if n.next == nil {
		list.tail = n
	}

	list.size++
	return true
}

func (list *singly[T]) RemoveAt(i int) (T, bool) {
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

	prev := list.head
	for prev != nil && i > 1 {
		prev = prev.next
		i--
	}

	if prev == nil || prev.next == nil {
		return zero, false
	}

	value := prev.next.Value
	prev.next = (prev.next).next

	list.size--
	return value, true
}

func (list *singly[T]) RemoveAfter(i int) (T, bool) {
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

	value := curr.next.Value
	curr.next = (curr.next).next

	if curr.next == nil {
		list.tail = curr
	}

	list.size--
	return value, true
}

func (list *singly[T]) RemoveAtNode(node *Node[T]) (T, bool) {
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

	if curr == nil {
		return zero, false
	}

	value := curr.next.Value
	curr.next = (curr.next).next
	if curr.next == nil {
		list.tail = curr // Already covered but keeping for clarity
	}

	list.size--
	return value, true
}

func (list *singly[T]) RemoveAfterNode(node *Node[T]) (T, bool) {
	var zero T

	if node == nil {
		return zero, false
	}

	curr := list.head
	for curr != nil && curr != node {
		curr = curr.next
	}

	if curr == nil || curr.next == nil {
		return zero, false
	}

	value := curr.next.Value
	curr.next = (curr.next).next
	if curr.next == nil {
		list.tail = curr
	}

	list.size--
	return value, true
}

func (list *singly[T]) RemoveValue(value T) bool {
	if list.head == nil {
		return false
	}

	if list.head.Value == value {
		_, found := list.RemoveFirst()
		return found
	}

	curr := list.head
	for curr != nil && curr.next != nil && curr.next.Value != value {
		curr = curr.next
	}

	if curr == nil || curr.next == nil {
		return false
	}

	curr.next = (curr.next).next
	if curr.next == nil {
		list.tail = curr
	}

	list.size--
	return true
}

func (list singly[T]) Get(i int) (T, bool) {
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

func (list singly[T]) GetAll() []T {
	out := make([]T, 0, list.size)
	curr := list.head
	for curr != nil {
		out = append(out, curr.Value)
		curr = curr.next
	}

	return out
}

func (list singly[T]) Search(value T) int {
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

func (list singly[T]) Contains(value T) bool {
	return list.Search(value) != -1
}

func (list singly[T]) Traverse(fn func(i int, value T) bool) {
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

func (list singly[T]) Reverse() *singly[T] {
	var prev, next *Node[T]

	out := list.Copy()
	curr := out.head
	for curr != nil {
		next = curr.next
		curr.next = prev

		prev = curr
		curr = next
	}
	return out
}

func (list singly[T]) IsSorted() bool {
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

func (list *singly[T]) Swap(i, j int) error {
	if i < 0 || j < 0 || i >= list.size || j >= list.size {
		return fmt.Errorf("index out of bounds: i=%d, j=%d", i, j)
	}

	if i == j {
		return nil
	}

	if i > j {
		i, j = j, i
	}

	var prev1, n1, prev2, n2 *Node[T]

	index := 0
	curr := list.head
	var prev *Node[T]

	for curr != nil {
		if index == i {
			prev1, n1 = prev, curr
		}
		if index == j {
			prev2, n2 = prev, curr
			break
		}

		index++
		prev = curr
		curr = curr.next
	}

	if n1 == nil || n2 == nil {
		return fmt.Errorf("nodes not found at i=%d, j=%d", i, j)
	}

	if prev1 == nil {
		list.head = n2
	} else {
		prev1.next = n2
	}

	if n1.next == n2 {
		n1.next, n2.next = n2.next, n1

		if n1.next == nil {
			list.tail = n1
		}

		return nil
	}

	if prev2 == nil {
		list.head = n1
	} else {
		prev2.next = n1
	}

	n1.next, n2.next = n2.next, n1.next

	if n1.next == nil {
		list.tail = n1
	} else if n2.next == nil {
		list.tail = n2
	}

	return nil
}
