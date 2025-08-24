package singly

import (
	"fmt"
	"strings"

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
	l.Traverse(func(index int, value T) { out.AddLast(value) })
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

func (l Singly[T]) GetFirstNode() *Node[T] {
	return l.head
}

func (l Singly[T]) GetLast() (T, bool) {
	var zero T
	if l.tail == nil {
		return zero, false
	}
	return l.tail.Value, true
}

func (l Singly[T]) GetLastNode() *Node[T] {
	return l.tail
}

func (l *Singly[T]) DeleteFirst() (T, bool) {
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

func (l *Singly[T]) DeleteLast() (T, bool) {
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

func (l *Singly[T]) SetAt(index int, value T) bool {
	if index < 0 || index >= l.size {
		return false
	}

	cur := l.head
	for cur != nil && index > 0 {
		cur = cur.next
		index--
	}

	if cur == nil {
		return false
	}

	cur.Value = value
	return true
}

func (l *Singly[T]) SetAtNode(node *Node[T], value T) bool {
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

func (l *Singly[T]) InsertAt(index int, value T) bool {
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

	prev := l.head
	for prev != nil && index > 1 {
		prev = prev.next
		index--
	}

	if prev == nil {
		return false
	}

	n := &Node[T]{Value: value, next: prev.next}
	prev.next = n

	if n.next == nil {
		l.tail = n
	}

	l.size++
	return true
}

func (l *Singly[T]) InsertAtNode(node *Node[T], value T) bool {
	if node == nil {
		return false
	}

	if l.head == node {
		l.AddFirst(value)
		return true
	}
	if l.tail == node {
		l.AddLast(value)
		return true
	}

	prev := l.head
	for prev != nil && prev.next != node {
		prev = prev.next
	}

	if prev == nil {
		return false
	}

	n := &Node[T]{Value: value, next: prev.next}
	prev.next = n

	if n.next == nil {
		l.tail = n
	}

	l.size++
	return true
}

func (l *Singly[T]) InsertAfter(index int, value T) bool {
	if index < 0 || index >= l.size {
		return false
	}

	if index == l.size-1 {
		l.AddLast(value)
		return true
	}

	cur := l.head
	for cur != nil && index > 0 {
		cur = cur.next
		index--
	}

	if cur == nil {
		return false
	}

	n := &Node[T]{Value: value, next: cur.next}
	cur.next = n

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

	cur := l.head
	for cur != nil && cur != node {
		cur = cur.next
	}

	if cur == nil {
		return false
	}

	n := &Node[T]{Value: value, next: cur.next}
	cur.next = n

	if n.next == nil {
		l.tail = n
	}

	l.size++
	return true
}

func (l *Singly[T]) DeleteAt(index int) (T, bool) {
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

	prev := l.head
	for prev != nil && index > 1 {
		prev = prev.next
		index--
	}

	if prev == nil || prev.next == nil {
		return zero, false
	}

	value := prev.next.Value
	prev.next = prev.next.next

	l.size--
	return value, true
}

func (l *Singly[T]) DeleteAfter(index int) (T, bool) {
	var zero T
	if index < 0 || index >= l.size {
		return zero, false
	}

	cur := l.head
	for cur != nil && index > 0 {
		cur = cur.next
		index--
	}

	if cur == nil || cur.next == nil {
		return zero, false
	}

	value := cur.next.Value
	cur.next = cur.next.next

	if cur.next == nil {
		l.tail = cur
	}

	l.size--
	return value, true
}

func (l *Singly[T]) DeleteAtNode(node *Node[T]) (T, bool) {
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

func (l *Singly[T]) DeleteAfterNode(node *Node[T]) (T, bool) {
	var zero T

	if node == nil {
		return zero, false
	}

	cur := l.head
	for cur != nil && cur != node {
		cur = cur.next
	}

	if cur == nil || cur.next == nil {
		return zero, false
	}

	value := cur.next.Value
	cur.next = cur.next.next

	if cur.next == nil {
		l.tail = cur
	}

	l.size--
	return value, true
}

func (l *Singly[T]) DeleteValue(value T) bool {
	if l.head == nil {
		return false
	}

	if utils.Is(l.head.Value, utils.EqualTo, value) {
		l.DeleteFirst()
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

func (l Singly[T]) Get(index int) (T, bool) {
	var zero T
	if index < 0 || index >= l.size {
		return zero, false
	}

	cur := l.head
	for cur != nil && index > 0 {
		cur = cur.next
		index--
	}

	if cur == nil {
		return zero, false
	}

	return cur.Value, true
}

func (l Singly[T]) GetNode(index int) *Node[T] {
	if index < 0 || index >= l.size {
		return nil
	}

	cur := l.head
	for cur != nil && index > 0 {
		cur = cur.next
		index--
	}

	return cur
}

func (l Singly[T]) ToSlice() []T {
	out := make([]T, 0, l.size)
	l.Traverse(func(index int, value T) { out = append(out, value) })
	return out
}

func (l Singly[T]) Search(value T) int {
	i := 0
	for cur := l.head; cur != nil; cur = cur.next {
		if utils.Is(cur.Value, utils.EqualTo, value) {
			return i
		}
		i++
	}
	return -1
}

func (l Singly[T]) Contains(value T) bool {
	return l.Search(value) != -1
}

func (l Singly[T]) Traverse(fn func(index int, value T)) {
	i := 0
	for cur := l.head; cur != nil; cur = cur.next {
		fn(i, cur.Value)
		i++
	}
}

func (l Singly[T]) Reverse() *Singly[T] {
	var prev, next *Node[T]
	out := l.Copy()
	cur := out.head
	for cur != nil {
		next = cur.next
		cur.next = prev

		prev = cur
		cur = next
	}
	out.head, out.tail = out.tail, out.head
	return out
}

func (l Singly[T]) IsSorted() bool {
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

func (l *Singly[T]) Swap(index1, index2 int) bool {
	if index1 < 0 || index2 < 0 || index1 >= l.size || index2 >= l.size {
		return false
	}

	if index1 == index2 {
		return true
	}

	if index1 > index2 {
		index1, index2 = index2, index1
	}

	i := 0
	var prev, prev1, prev2, n1, n2 *Node[T]
	for cur := l.head; cur != nil; cur = cur.next {
		if i == index1 {
			prev1, n1 = prev, cur
		}

		if i == index2 {
			prev2, n2 = prev, cur
			break
		}

		i++
		prev = cur
	}

	if n1 == nil || n2 == nil {
		return false
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

		return true
	}

	prev2.next = n1
	n1.next, n2.next = n2.next, n1.next

	if n1.next == nil {
		l.tail = n1
	} else if n2.next == nil {
		l.tail = n2
	}

	return true
}

func (l Singly[T]) String() string {
	if l.head == nil {
		return "[nil]"
	}

	var sb strings.Builder
	sb.WriteString("[(head) ")

	for cur := l.head; cur != nil; cur = cur.next {
		fmt.Fprintf(&sb, "%v", cur.Value)
		if cur.next != nil {
			sb.WriteString(" -> ")
		} else {
			sb.WriteString(" (tail) -> nil]")
		}
	}
	return sb.String()
}
