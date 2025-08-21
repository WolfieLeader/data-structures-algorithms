package doubly

import (
	"fmt"

	"github.com/WolfieLeader/data-structures-algorithms/utils"
)

func (l Doubly[T]) Size() int {
	return l.size
}

func (l Doubly[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *Doubly[T]) Clear() {
	l.size = 0
	l.head = nil
	l.tail = nil
}

func (l Doubly[T]) Copy() *Doubly[T] {
	if l.size == 0 {
		return New[T]()
	}

	out := New[T]()
	l.Traverse(func(i int, value T) bool {
		out.AddLast(value)
		return true
	})
	return out
}

func (l *Doubly[T]) AddFirst(values ...T) {
	for _, value := range values {
		n := &Node[T]{Value: value, next: l.head}

		if l.head == nil {
			l.tail = n
		} else {
			l.head.prev = n
		}

		l.head = n
		l.size++
	}
}

func (l *Doubly[T]) AddLast(values ...T) {
	for _, value := range values {
		n := &Node[T]{Value: value, prev: l.tail}

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

func (l Doubly[T]) GetFirst() (T, bool) {
	var zero T
	if l.head == nil {
		return zero, false
	}
	return l.head.Value, true
}

func (l Doubly[T]) GetLast() (T, bool) {
	var zero T
	if l.tail == nil {
		return zero, false
	}
	return l.tail.Value, true
}

func (l *Doubly[T]) RemoveFirst() (T, bool) {
	var zero T
	if l.head == nil {
		return zero, false
	}

	value := l.head.Value
	l.head = l.head.next

	if l.head == nil {
		l.tail = nil
	} else {
		l.head.prev = nil
	}

	l.size--
	return value, true
}

func (l *Doubly[T]) RemoveLast() (T, bool) {
	var zero T
	if l.tail == nil {
		return zero, false
	}

	value := l.tail.Value
	l.tail = l.tail.prev

	if l.tail == nil {
		l.head = nil
	} else {
		l.tail.next = nil
	}

	l.size--
	return value, true
}

func (l Doubly[T]) shortestPath(i int) *Node[T] {
	var curr *Node[T]
	if i <= (l.size-1)/2 {
		curr = l.head
		for curr != nil && i > 0 {
			curr = curr.next
			i--
		}
		return curr
	}

	curr = l.tail
	for curr != nil && i < l.size-1 {
		curr = curr.prev
		i++
	}
	return curr
}

func (l *Doubly[T]) SetAt(i int, value T) bool {
	if i < 0 || i >= l.size {
		return false
	}

	curr := l.shortestPath(i)
	if curr == nil {
		return false
	}

	curr.Value = value
	return true
}

func (l *Doubly[T]) SetAtNode(node *Node[T], value T) bool {
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

func (l *Doubly[T]) InsertAfter(i int, value T) bool {
	if i < 0 || i >= l.size {
		return false
	}

	if i == l.size-1 {
		l.AddLast(value)
		return true
	}

	curr := l.shortestPath(i)
	if curr == nil {
		return false
	}

	n := &Node[T]{Value: value, prev: curr, next: curr.next}
	curr.next = n

	if n.next == nil {
		l.tail = n
	} else {
		n.next.prev = n
	}

	l.size++
	return true
}

func (l *Doubly[T]) InsertAfterNode(node *Node[T], value T) bool {
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

	n := &Node[T]{Value: value, prev: curr, next: curr.next}
	curr.next = n

	if n.next == nil {
		l.tail = n
	} else {
		n.next.prev = n
	}

	l.size++
	return true
}

func (l *Doubly[T]) unlink(curr *Node[T]) (T, bool) {
	var zero T
	if curr == nil {
		return zero, false
	}

	if curr.prev == nil {
		l.head = curr.next
	} else {
		curr.prev.next = curr.next
	}

	if curr.next == nil {
		l.tail = curr.prev
	} else {
		curr.next.prev = curr.prev
	}

	l.size--
	return curr.Value, true
}

func (l *Doubly[T]) RemoveAt(i int) (T, bool) {
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

	curr := l.shortestPath(i)
	return l.unlink(curr)
}

func (l *Doubly[T]) RemoveAfter(i int) (T, bool) {
	var zero T
	if i < 0 || i >= l.size {
		return zero, false
	}

	curr := l.shortestPath(i)
	if curr == nil {
		return zero, false
	}

	return l.unlink(curr.next)
}

func (l *Doubly[T]) RemoveAtNode(node *Node[T]) (T, bool) {
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
	for curr != nil && curr != node {
		curr = curr.next
	}

	return l.unlink(node)
}

func (l *Doubly[T]) RemoveAfterNode(node *Node[T]) (T, bool) {
	var zero T

	if node == nil {
		return zero, false
	}

	curr := l.head
	for curr != nil && curr != node {
		curr = curr.next
	}

	if curr == nil {
		return zero, false
	}

	return l.unlink(curr.next)
}

func (l *Doubly[T]) RemoveValue(value T) bool {
	if l.head == nil {
		return false
	}

	if utils.Is(l.head.Value, utils.EqualTo, value) {
		l.RemoveFirst()
		return true
	}
	if utils.Is(l.tail.Value, utils.EqualTo, value) {
		l.RemoveLast()
		return true
	}

	curr := l.head
	for curr != nil {
		if utils.Is(curr.Value, utils.EqualTo, value) {
			l.unlink(curr)
			return true
		}

		curr = curr.next
	}
	return false
}

func (l Doubly[T]) Get(i int) (T, bool) {
	var zero T
	if i < 0 || i >= l.size {
		return zero, false
	}

	curr := l.shortestPath(i)
	if curr == nil {
		return zero, false
	}

	return curr.Value, true
}

func (l Doubly[T]) ToSlice() []T {
	out := make([]T, 0, l.size)

	l.Traverse(func(i int, value T) bool {
		out = append(out, value)
		return true
	})

	return out
}

func (l Doubly[T]) Search(value T) int {
	i := 0
	for curr := l.head; curr != nil; curr = curr.next {
		if utils.Is(curr.Value, utils.EqualTo, value) {
			return i
		}
		i++
	}
	return -1
}

func (l Doubly[T]) Contains(value T) bool {
	return l.Search(value) != -1
}

func (l Doubly[T]) Traverse(fn func(i int, value T) bool) {
	i := 0
	for curr := l.head; curr != nil; curr = curr.next {
		if !fn(i, curr.Value) {
			break
		}
		i++
	}
}

func (l Doubly[T]) Reverse() *Doubly[T] {
	out := l.Copy()
	curr := out.head
	for curr != nil {
		curr.prev, curr.next = curr.next, curr.prev
		curr = curr.prev
	}
	out.head, out.tail = out.tail, out.head
	return out
}

func (l Doubly[T]) IsSorted() bool {
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

func (l *Doubly[T]) Swap(i, j int) error {
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
	var n1, n2 *Node[T]
	for curr := l.head; curr != nil; curr = curr.next {
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
		l.head = n2
		n2.prev = nil
	} else {
		n1.prev.next = n2
		n2.prev = n1.prev
	}

	if n1.next == n2 {
		n1.next = n2.next

		if n2.next == nil {
			l.tail = n1
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
		l.tail = n2
	} else {
		ogNext1.prev = n2
	}

	if ogNext2 == nil {
		l.tail = n1
	} else {
		ogNext2.prev = n1
	}

	n1.next, n2.next = ogNext2, ogNext1

	return nil
}
