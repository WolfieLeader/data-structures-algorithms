package singly

import (
	"fmt"

	"github.com/WolfieLeader/data-structures-algorithms/utils"
)

func (l Singly[T]) Size() int {
	return l.size
}

func (l Singly[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *Singly[T]) Clear() {
	l.size = 0
	l.head = nil
	l.tail = nil
}

func (l Singly[T]) Copy() *Singly[T] {
	if l.size == 0 {
		return New[T]()
	}

	out := New[T]()
	l.ForEach(func(i int, value T) bool {
		out.AddLast(value)
		return true
	})
	return out
}

func (l *Singly[T]) AddFirst(values ...T) {
	for _, value := range values {
		n := &Node[T]{Value: value, next: l.head}

		if l.head == nil {
			l.tail = n
		}

		l.head = n
		l.size++
	}
}

func (l *Singly[T]) AddLast(values ...T) {
	for _, value := range values {
		n := &Node[T]{Value: value, next: nil}

		if l.tail == nil {
			l.head, l.tail = n, n
			l.size++
			continue
		}

		l.tail.next = n
		l.tail = n
		l.size++
	}
}

func (l Singly[T]) GetFirst() (T, bool) {
	var zero T
	if l.head == nil {
		return zero, false
	}
	return l.head.Value, true
}

func (l Singly[T]) GetLast() (T, bool) {
	var zero T
	if l.tail == nil {
		return zero, false
	}
	return l.tail.Value, true
}

func (l *Singly[T]) RemoveFirst() (T, bool) {
	var zero T
	if l.head == nil {
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

func (l *Singly[T]) RemoveLast() (T, bool) {
	var zero T
	if l.tail == nil {
		return zero, false
	}

	value := l.tail.Value

	if l.head == l.tail {
		l.Clear()
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

func (l *Singly[T]) SetAt(i int, value T) bool {
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

func (l *Singly[T]) SetAtNode(node *Node[T], value T) bool {
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

func (l *Singly[T]) InsertAfter(i int, value T) bool {
	if i < 0 || i >= l.size {
		return false
	}

	if i == l.size-1 {
		l.AddLast(value)
		return true
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

func (l *Singly[T]) InsertAfterNode(node *Node[T], value T) bool {
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

func (l *Singly[T]) RemoveAt(i int) (T, bool) {
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
	prev.next = prev.next.next

	l.size--
	return value, true
}

func (l *Singly[T]) RemoveAfter(i int) (T, bool) {
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
	curr.next = curr.next.next

	if curr.next == nil {
		l.tail = curr
	}

	l.size--
	return value, true
}

func (l *Singly[T]) RemoveAtNode(node *Node[T]) (T, bool) {
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

	prev := l.head
	for prev != nil && prev.next != node {
		prev = prev.next
	}

	if prev == nil || prev.next == nil {
		return zero, false
	}

	value := prev.next.Value
	prev.next = prev.next.next

	l.size--
	return value, true
}

func (l *Singly[T]) RemoveAfterNode(node *Node[T]) (T, bool) {
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
	curr.next = curr.next.next

	if curr.next == nil {
		l.tail = curr
	}

	l.size--
	return value, true
}

func (l *Singly[T]) RemoveValue(value T) bool {
	if l.head == nil {
		return false
	}

	if utils.Is(l.head.Value, utils.EqualTo, value) {
		l.RemoveFirst()
		return true
	}

	prev := l.head
	for prev != nil && prev.next != nil {
		if utils.Is(prev.next.Value, utils.EqualTo, value) {
			prev.next = prev.next.next

			if prev.next == nil {
				l.tail = prev
			}

			l.size--
			return true
		}

		prev = prev.next
	}
	return false
}

func (l Singly[T]) Get(i int) (T, bool) {
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

func (l Singly[T]) ToSlice() []T {
	out := make([]T, 0, l.size)

	l.ForEach(func(i int, value T) bool {
		out = append(out, value)
		return true
	})

	return out
}

func (l Singly[T]) Search(value T) int {
	i := 0
	for curr := l.head; curr != nil; curr = curr.next {
		if utils.Is(curr.Value, utils.EqualTo, value) {
			return i
		}
		i++
	}
	return -1
}

func (l Singly[T]) Contains(value T) bool {
	return l.Search(value) != -1
}

func (l Singly[T]) ForEach(fn func(i int, value T) bool) {
	i := 0
	for curr := l.head; curr != nil; curr = curr.next {
		if !fn(i, curr.Value) {
			break
		}
		i++
	}
}

func (l Singly[T]) Reverse() *Singly[T] {
	var prev, next *Node[T]
	out := l.Copy()
	curr := out.head
	for curr != nil {
		next = curr.next
		curr.next = prev

		prev = curr
		curr = next
	}
	out.head, out.tail = out.tail, out.head
	return out
}

func (l Singly[T]) IsSorted() bool {
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

func (l *Singly[T]) Swap(i, j int) error {
	if i < 0 || j < 0 || i >= l.size || j >= l.size {
		return fmt.Errorf("index out of bounds: i=%d, j=%d", i, j)
	}

	if i == j {
		return nil
	}

	if i > j {
		i, j = j, i
	}

	index := 0
	var prev, prev1, prev2, n1, n2 *Node[T]
	for curr := l.head; curr != nil; curr = curr.next {
		if index == i {
			prev1, n1 = prev, curr
		}

		if index == j {
			prev2, n2 = prev, curr
			break
		}

		index++
		prev = curr
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
		n1.next = n2.next
		n2.next = n1

		if n1.next == nil {
			l.tail = n1
		}

		return nil
	}

	prev2.next = n1
	n1.next, n2.next = n2.next, n1.next

	if n1.next == nil {
		l.tail = n1
	} else if n2.next == nil {
		l.tail = n2
	}

	return nil
}
