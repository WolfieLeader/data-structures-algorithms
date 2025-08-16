package singly

import (
	"fmt"

	"github.com/WolfieLeader/data-structures-algorithms/utils"
)

func (l singlyLinkedList[T]) Size() int {
	return l.size
}

func (l singlyLinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *singlyLinkedList[T]) Clear() {
	l.size = 0
	l.head = nil
	l.tail = nil
}

func (l singlyLinkedList[T]) Copy() *singlyLinkedList[T] {
	if l.size == 0 {
		return New[T]()
	}

	out := New[T]()
	curr := l.head
	for curr != nil {
		out.AddLast(curr.Value)
		curr = curr.next
	}
	return out
}

func (l *singlyLinkedList[T]) AddFirst(values ...T) {
	for _, value := range values {
		n := &Node[T]{Value: value, next: l.head}
		l.head = n
		if l.tail == nil {
			l.tail = n
		}
		l.size++
	}
}

func (l *singlyLinkedList[T]) AddLast(values ...T) {
	for _, value := range values {
		n := &Node[T]{Value: value}
		if l.tail == nil {
			l.head, l.tail = n, n
		} else {
			l.tail.next = n
			l.tail = n
		}
		l.size++
	}
}

func (l singlyLinkedList[T]) GetFirst() (T, bool) {
	if l.head == nil {
		var zero T
		return zero, false
	}
	return l.head.Value, true
}

func (l singlyLinkedList[T]) GetLast() (T, bool) {
	if l.tail == nil {
		var zero T
		return zero, false
	}
	return l.tail.Value, true
}

func (l *singlyLinkedList[T]) RemoveFirst() (T, bool) {
	if l.head == nil {
		var zero T
		return zero, false
	}

	value := l.head.Value
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	}

	l.size--
	return value, true
}

func (l *singlyLinkedList[T]) RemoveLast() (T, bool) {
	if l.tail == nil {
		var zero T
		return zero, false
	}

	value := l.tail.Value
	if l.head == l.tail {
		l.head, l.tail = nil, nil
		l.size = 0
		return value, true
	}

	prev := l.head
	for prev.next != l.tail {
		prev = prev.next
	}

	prev.next = nil
	l.tail = prev

	l.size--
	return value, true
}

func (l *singlyLinkedList[T]) SetAt(i int, value T) bool {
	if i < 0 || i >= l.size {
		return false
	}

	curr := l.head
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

func (l *singlyLinkedList[T]) SetAtNode(node *Node[T], value T) bool {
	if node == nil {
		return false
	}

	curr := l.head
	for curr != nil && curr != node {
		curr = curr.next
	}

	if curr == nil {
		return false
	}

	curr.Value = value
	return true
}

func (l *singlyLinkedList[T]) InsertAfter(i int, value T) bool {
	if i < 0 || i >= l.size {
		return false
	}

	curr := l.head
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
		l.tail = n
	}

	l.size++
	return true
}

func (l *singlyLinkedList[T]) InsertAfterNode(node *Node[T], value T) bool {
	if node == nil {
		return false
	}

	curr := l.head
	for curr != nil && curr != node {
		curr = curr.next
	}

	if curr == nil {
		return false
	}

	n := &Node[T]{Value: value, next: curr.next}
	curr.next = n

	if n.next == nil {
		l.tail = n
	}

	l.size++
	return true
}

func (l *singlyLinkedList[T]) RemoveAt(i int) (T, bool) {
	var zero T

	if i < 0 || i >= l.size {
		return zero, false
	}

	if i == 0 {
		return l.RemoveFirst()
	}

	if i == l.size-1 {
		return l.RemoveLast()
	}

	prev := l.head
	for prev != nil && i > 1 {
		prev = prev.next
		i--
	}

	if prev == nil || prev.next == nil {
		return zero, false
	}

	value := prev.next.Value
	prev.next = (prev.next).next

	l.size--
	return value, true
}

func (l *singlyLinkedList[T]) RemoveAfter(i int) (T, bool) {
	var zero T

	if i < 0 || i >= l.size {
		return zero, false
	}

	curr := l.head
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
		l.tail = curr
	}

	l.size--
	return value, true
}

func (l *singlyLinkedList[T]) RemoveAtNode(node *Node[T]) (T, bool) {
	var zero T

	if node == nil {
		return zero, false
	}

	if l.head == node {
		return l.RemoveFirst()
	}

	if l.tail == node {
		return l.RemoveLast()
	}

	curr := l.head
	for curr != nil && curr.next != node {
		curr = curr.next
	}

	if curr == nil {
		return zero, false
	}

	value := curr.next.Value
	curr.next = (curr.next).next
	if curr.next == nil {
		l.tail = curr // Already covered but keeping for clarity
	}

	l.size--
	return value, true
}

func (l *singlyLinkedList[T]) RemoveAfterNode(node *Node[T]) (T, bool) {
	var zero T

	if node == nil {
		return zero, false
	}

	curr := l.head
	for curr != nil && curr != node {
		curr = curr.next
	}

	if curr == nil || curr.next == nil {
		return zero, false
	}

	value := curr.next.Value
	curr.next = (curr.next).next
	if curr.next == nil {
		l.tail = curr
	}

	l.size--
	return value, true
}

func (l *singlyLinkedList[T]) RemoveValue(value T) bool {
	if l.head == nil {
		return false
	}

	if l.head.Value == value {
		_, found := l.RemoveFirst()
		return found
	}

	curr := l.head
	for curr != nil && curr.next != nil && curr.next.Value != value {
		curr = curr.next
	}

	if curr == nil || curr.next == nil {
		return false
	}

	curr.next = (curr.next).next
	if curr.next == nil {
		l.tail = curr
	}

	l.size--
	return true
}

func (l singlyLinkedList[T]) Get(i int) (T, bool) {
	var zero T

	if i < 0 || i >= l.size {
		return zero, false
	}

	curr := l.head
	for curr != nil && i > 0 {
		curr = curr.next
		i--
	}

	if curr == nil {
		return zero, false
	}

	return curr.Value, true
}

func (l singlyLinkedList[T]) GetAll() []T {
	out := make([]T, 0, l.size)
	curr := l.head
	for curr != nil {
		out = append(out, curr.Value)
		curr = curr.next
	}

	return out
}

func (l singlyLinkedList[T]) LinearSearch(value T) int {
	i := 0
	curr := l.head
	for curr != nil && curr.Value != value {
		curr = curr.next
		i++
	}

	if curr == nil {
		return -1
	}

	return i
}

func (l singlyLinkedList[T]) Contains(value T) bool {
	return l.LinearSearch(value) != -1
}

func (l singlyLinkedList[T]) Traverse(fn func(i int, value T) bool) {
	i := 0
	curr := l.head
	for curr != nil {
		if !fn(i, curr.Value) {
			break
		}
		curr = curr.next
		i++
	}
}

func (l singlyLinkedList[T]) Reverse() *singlyLinkedList[T] {
	out := New[T]()
	curr := l.head
	for curr != nil {
		out.AddFirst(curr.Value)
		curr = curr.next
	}
	return out
}

func (l singlyLinkedList[T]) IsSorted() bool {
	if l.size <= 1 {
		return true
	}

	curr := l.head
	for curr != nil && curr.next != nil {
		if utils.Is(curr.Value, utils.GreaterThan, curr.next.Value) {
			return false
		}
		curr = curr.next
	}
	return true
}

func (l *singlyLinkedList[T]) Swap(i, j int) error {
	if i < 0 || j < 0 || i >= l.size || j >= l.size {
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
	curr := l.head
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
		l.head = n2
	} else {
		prev1.next = n2
	}

	if n1.next == n2 {
		n1.next, n2.next = n2.next, n1

		if n1.next == nil {
			l.tail = n1
		}

		return nil
	}

	if prev2 == nil {
		l.head = n1
	} else {
		prev2.next = n1
	}

	n1.next, n2.next = n2.next, n1.next

	if n1.next == nil {
		l.tail = n1
	} else if n2.next == nil {
		l.tail = n2
	}

	return nil
}
