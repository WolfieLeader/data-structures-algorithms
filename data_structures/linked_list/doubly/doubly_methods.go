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
	l.Traverse(func(index int, value T) { out.AddLast(value) })
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

func (l *Doubly[T]) DeleteFirst() (T, bool) {
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

func (l *Doubly[T]) DeleteLast() (T, bool) {
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

func (l Doubly[T]) shortestPath(index int) *Node[T] {
	var cur *Node[T]
	if index <= (l.size-1)/2 {
		cur = l.head
		for cur != nil && index > 0 {
			cur = cur.next
			index--
		}
		return cur
	}

	cur = l.tail
	for cur != nil && index < l.size-1 {
		cur = cur.prev
		index++
	}
	return cur
}

func (l *Doubly[T]) SetAt(index int, value T) bool {
	if index < 0 || index >= l.size {
		return false
	}

	cur := l.shortestPath(index)
	if cur == nil {
		return false
	}

	cur.Value = value
	return true
}

func (l *Doubly[T]) SetAtNode(node *Node[T], value T) bool {
	if node == nil {
		return false
	}

	cur := l.head
	for cur != nil && cur != node {
		cur = cur.next
	}

	if cur == nil {
		return false
	}

	cur.Value = value
	return true
}

func (l *Doubly[T]) InsertAt(index int, value T) bool {
	if index < 0 || index > l.size {
		return false
	}

	if index == 0 {
		l.AddFirst(value)
		return true
	}
	if index == l.size {
		l.AddLast(value)
		return true
	}

	cur := l.shortestPath(index)
	if cur == nil {
		return false
	}

	n := &Node[T]{Value: value, prev: cur.prev, next: cur}
	cur.prev.next = n
	cur.prev = n

	l.size++
	return true
}

func (l *Doubly[T]) InsertAtNode(node *Node[T], value T) bool {
	if node == nil {
		return false
	}

	cur := l.head
	for cur != nil && cur != node {
		cur = cur.next
	}

	if cur == nil {
		return false
	}

	n := &Node[T]{Value: value, prev: cur.prev, next: cur}
	cur.prev.next = n
	cur.prev = n

	l.size++
	return true
}

func (l *Doubly[T]) InsertAfter(index int, value T) bool {
	if index < 0 || index >= l.size {
		return false
	}

	if index == l.size-1 {
		l.AddLast(value)
		return true
	}

	cur := l.shortestPath(index)
	if cur == nil {
		return false
	}

	n := &Node[T]{Value: value, prev: cur, next: cur.next}
	cur.next = n

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

	cur := l.head
	for cur != nil && cur != node {
		cur = cur.next
	}

	if cur == nil {
		return false
	}

	n := &Node[T]{Value: value, prev: cur, next: cur.next}
	cur.next = n

	if n.next == nil {
		l.tail = n
	} else {
		n.next.prev = n
	}

	l.size++
	return true
}

func (l *Doubly[T]) unlink(cur *Node[T]) (T, bool) {
	var zero T
	if cur == nil {
		return zero, false
	}

	if cur.prev == nil {
		l.head = cur.next
	} else {
		cur.prev.next = cur.next
	}

	if cur.next == nil {
		l.tail = cur.prev
	} else {
		cur.next.prev = cur.prev
	}

	l.size--
	return cur.Value, true
}

func (l *Doubly[T]) DeleteAt(index int) (T, bool) {
	var zero T
	if index < 0 || index >= l.size {
		return zero, false
	}

	if index == 0 {
		return l.DeleteFirst()
	}
	if index == l.size-1 {
		return l.DeleteLast()
	}

	cur := l.shortestPath(index)
	return l.unlink(cur)
}

func (l *Doubly[T]) DeleteAfter(index int) (T, bool) {
	var zero T
	if index < 0 || index >= l.size {
		return zero, false
	}

	cur := l.shortestPath(index)
	if cur == nil {
		return zero, false
	}

	return l.unlink(cur.next)
}

func (l *Doubly[T]) DeleteAtNode(node *Node[T]) (T, bool) {
	var zero T
	if node == nil {
		return zero, false
	}

	if l.head == node {
		return l.DeleteFirst()
	}
	if l.tail == node {
		return l.DeleteLast()
	}

	cur := l.head
	for cur != nil && cur != node {
		cur = cur.next
	}

	return l.unlink(node)
}

func (l *Doubly[T]) DeleteAfterNode(node *Node[T]) (T, bool) {
	var zero T

	if node == nil {
		return zero, false
	}

	cur := l.head
	for cur != nil && cur != node {
		cur = cur.next
	}

	if cur == nil {
		return zero, false
	}

	return l.unlink(cur.next)
}

func (l *Doubly[T]) DeleteValue(value T) bool {
	if l.head == nil {
		return false
	}

	if utils.Is(l.head.Value, utils.EqualTo, value) {
		l.DeleteFirst()
		return true
	}
	if utils.Is(l.tail.Value, utils.EqualTo, value) {
		l.DeleteLast()
		return true
	}

	cur := l.head
	for cur != nil {
		if utils.Is(cur.Value, utils.EqualTo, value) {
			l.unlink(cur)
			return true
		}

		cur = cur.next
	}
	return false
}

func (l Doubly[T]) Get(i int) (T, bool) {
	var zero T
	if i < 0 || i >= l.size {
		return zero, false
	}

	cur := l.shortestPath(i)
	if cur == nil {
		return zero, false
	}

	return cur.Value, true
}

func (l Doubly[T]) ToSlice() []T {
	out := make([]T, 0, l.size)
	l.Traverse(func(i int, value T) { out = append(out, value) })
	return out
}

func (l Doubly[T]) Search(value T) int {
	i := 0
	for cur := l.head; cur != nil; cur = cur.next {
		if utils.Is(cur.Value, utils.EqualTo, value) {
			return i
		}
		i++
	}
	return -1
}

func (l Doubly[T]) Contains(value T) bool {
	return l.Search(value) != -1
}

func (l Doubly[T]) Traverse(fn func(index int, value T)) {
	i := 0
	for cur := l.head; cur != nil; cur = cur.next {
		fn(i, cur.Value)
		i++
	}
}

func (l Doubly[T]) Reverse() *Doubly[T] {
	out := l.Copy()
	cur := out.head
	for cur != nil {
		cur.prev, cur.next = cur.next, cur.prev
		cur = cur.prev
	}
	out.head, out.tail = out.tail, out.head
	return out
}

func (l Doubly[T]) IsSorted() bool {
	if l.size <= 1 {
		return true
	}

	cur := l.head
	for cur != nil && cur.next != nil {
		if utils.Is(cur.Value, utils.GreaterThan, cur.next.Value) {
			return false
		}
		cur = cur.next
	}
	return true
}

func (l *Doubly[T]) Swap(index1, index2 int) error {
	if index1 < 0 || index2 < 0 || index1 >= l.size || index2 >= l.size {
		return fmt.Errorf("index out of bounds: i=%d, j=%d", index1, index2)
	}

	if index1 == index2 {
		return nil
	}

	if index1 > index2 {
		index1, index2 = index2, index1
	}

	i := 0
	var n1, n2 *Node[T]
	for cur := l.head; cur != nil; cur = cur.next {
		if i == index1 {
			n1 = cur
		}
		if i == index2 {
			n2 = cur
			break
		}

		i++
	}

	if n1 == nil || n2 == nil {
		return fmt.Errorf("nodes not found at i=%d, j=%d", index1, index2)
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

	realNext1, realNext2 := n1.next, n2.next

	if realNext1 == nil {
		l.tail = n2
	} else {
		realNext1.prev = n2
	}

	if realNext2 == nil {
		l.tail = n1
	} else {
		realNext2.prev = n1
	}

	n1.next, n2.next = realNext2, realNext1

	return nil
}
