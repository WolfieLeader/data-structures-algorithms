package singly

import (
	"fmt"

	"github.com/WolfieLeader/data-structures-algorithms/utils"
)

func (s singlyLinkedList[T]) Size() int {
	return s.size
}

func (s singlyLinkedList[T]) IsEmpty() bool {
	return s.size == 0
}

func (s *singlyLinkedList[T]) Clear() {
	s.size = 0
	s.head = nil
	s.tail = nil
}

func (s singlyLinkedList[T]) Copy() *singlyLinkedList[T] {
	if s.size == 0 {
		return New[T]()
	}

	out := New[T]()
	curr := s.head
	for curr != nil {
		out.AddLast(curr.Value)
		curr = curr.next
	}
	return out
}

func (s *singlyLinkedList[T]) AddFirst(values ...T) {
	for _, value := range values {
		n := &Node[T]{Value: value, next: s.head}
		s.head = n
		if s.head == nil {
			s.tail = n
		}
		s.size++
	}
}

func (s *singlyLinkedList[T]) AddLast(values ...T) {
	for _, value := range values {
		n := &Node[T]{Value: value}
		if s.tail == nil {
			s.head = n
		} else {
			s.tail.next = n
		}
		s.tail = n
		s.size++
	}
}

func (s singlyLinkedList[T]) GetFirst() (T, bool) {
	if s.head == nil {
		var zero T
		return zero, false
	}
	return s.head.Value, true
}

func (s singlyLinkedList[T]) GetLast() (T, bool) {
	if s.tail == nil {
		var zero T
		return zero, false
	}
	return s.tail.Value, true
}

func (s *singlyLinkedList[T]) RemoveFirst() (T, bool) {
	if s.head == nil {
		var zero T
		return zero, false
	}

	value := s.head.Value
	s.head = s.head.next
	if s.head == nil {
		s.tail = nil
	}

	s.size--
	return value, true
}

func (s *singlyLinkedList[T]) RemoveLast() (T, bool) {
	if s.tail == nil {
		var zero T
		return zero, false
	}

	value := s.tail.Value
	if s.head == s.tail {
		s.head, s.tail = nil, nil
		s.size = 0
		return value, true
	}

	prev := s.head
	for prev.next != s.tail {
		prev = prev.next
	}

	prev.next = nil
	s.tail = prev

	s.size--
	return value, true
}

func (s *singlyLinkedList[T]) SetAt(i int, value T) bool {
	if i < 0 || i >= s.size {
		return false
	}

	curr := s.head
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

func (s *singlyLinkedList[T]) SetAtNode(node *Node[T], value T) bool {
	if node == nil {
		return false
	}

	curr := s.head
	for curr != nil && curr != node {
		curr = curr.next
	}

	if curr == nil {
		return false
	}

	curr.Value = value
	return true
}

func (s *singlyLinkedList[T]) InsertAfter(i int, value T) bool {
	if i < 0 || i >= s.size {
		return false
	}

	curr := s.head
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
		s.tail = n
	}

	s.size++
	return true
}

func (s *singlyLinkedList[T]) InsertAfterNode(node *Node[T], value T) bool {
	if node == nil {
		return false
	}

	curr := s.head
	for curr != nil && curr != node {
		curr = curr.next
	}

	if curr == nil {
		return false
	}

	n := &Node[T]{Value: value, next: curr.next}
	curr.next = n

	if n.next == nil {
		s.tail = n
	}

	s.size++
	return true
}

func (s *singlyLinkedList[T]) RemoveAt(i int) (T, bool) {
	var zero T

	if i < 0 || i >= s.size {
		return zero, false
	}

	if i == 0 {
		return s.RemoveFirst()
	}

	if i == s.size-1 {
		return s.RemoveLast()
	}

	prev := s.head
	for prev != nil && i > 1 {
		prev = prev.next
		i--
	}

	if prev == nil || prev.next == nil {
		return zero, false
	}

	value := prev.next.Value
	prev.next = (prev.next).next

	s.size--
	return value, true
}

func (s *singlyLinkedList[T]) RemoveAfter(i int) (T, bool) {
	var zero T

	if i < 0 || i >= s.size {
		return zero, false
	}

	curr := s.head
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
		s.tail = curr
	}

	s.size--
	return value, true
}

func (s *singlyLinkedList[T]) RemoveAtNode(node *Node[T]) (T, bool) {
	var zero T

	if node == nil {
		return zero, false
	}

	if s.head == node {
		return s.RemoveFirst()
	}

	if s.tail == node {
		return s.RemoveLast()
	}

	curr := s.head
	for curr != nil && curr.next != node {
		curr = curr.next
	}

	if curr == nil {
		return zero, false
	}

	value := curr.next.Value
	curr.next = (curr.next).next
	if curr.next == nil {
		s.tail = curr // Already covered but keeping for clarity
	}

	s.size--
	return value, true
}

func (s *singlyLinkedList[T]) RemoveAfterNode(node *Node[T]) (T, bool) {
	var zero T

	if node == nil {
		return zero, false
	}

	curr := s.head
	for curr != nil && curr != node {
		curr = curr.next
	}

	if curr == nil || curr.next == nil {
		return zero, false
	}

	value := curr.next.Value
	curr.next = (curr.next).next
	if curr.next == nil {
		s.tail = curr
	}

	s.size--
	return value, true
}

func (s *singlyLinkedList[T]) RemoveValue(value T) bool {
	if s.head == nil {
		return false
	}

	if s.head.Value == value {
		_, found := s.RemoveFirst()
		return found
	}

	curr := s.head
	for curr != nil && curr.next != nil && curr.next.Value != value {
		curr = curr.next
	}

	if curr == nil || curr.next == nil {
		return false
	}

	curr.next = (curr.next).next
	if curr.next == nil {
		s.tail = curr
	}

	s.size--
	return true
}

func (s singlyLinkedList[T]) Get(i int) (T, bool) {
	var zero T

	if i < 0 || i >= s.size {
		return zero, false
	}

	curr := s.head
	for curr != nil && i > 0 {
		curr = curr.next
		i--
	}

	if curr == nil {
		return zero, false
	}

	return curr.Value, true
}

func (s singlyLinkedList[T]) GetAll() []T {
	out := make([]T, 0, s.size)
	curr := s.head
	for curr != nil {
		out = append(out, curr.Value)
		curr = curr.next
	}

	return out
}

func (s singlyLinkedList[T]) LinearSearch(value T) int {
	i := 0
	curr := s.head
	for curr != nil && curr.Value != value {
		curr = curr.next
		i++
	}

	if curr == nil {
		return -1
	}

	return i
}

func (s singlyLinkedList[T]) Contains(value T) bool {
	return s.LinearSearch(value) != -1
}

func (s singlyLinkedList[T]) Traverse(fn func(i int, value T) bool) {
	i := 0
	curr := s.head
	for curr != nil {
		if !fn(i, curr.Value) {
			break
		}
		curr = curr.next
		i++
	}
}

func (s singlyLinkedList[T]) Reverse() *singlyLinkedList[T] {
	out := New[T]()
	curr := s.head
	for curr != nil {
		out.AddFirst(curr.Value)
		curr = curr.next
	}
	return out
}

func (s singlyLinkedList[T]) IsSorted() bool {
	if s.size <= 1 {
		return true
	}

	curr := s.head
	for curr != nil && curr.next != nil {
		if utils.Is(curr.Value, utils.GreaterThan, curr.next.Value) {
			return false
		}
		curr = curr.next
	}
	return true
}

func (s *singlyLinkedList[T]) Swap(i, j int) error {
	if i < 0 || j < 0 || i >= s.size || j >= s.size {
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
	curr := s.head
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
		s.head = n2
	} else {
		prev1.next = n2
	}

	if n1.next == n2 {
		n1.next, n2.next = n2.next, n1

		if n1.next == nil {
			s.tail = n1
		}

		return nil
	}

	if prev2 == nil {
		s.head = n1
	} else {
		prev2.next = n1
	}

	n1.next, n2.next = n2.next, n1.next

	if n1.next == nil {
		s.tail = n1
	} else if n2.next == nil {
		s.tail = n2
	}

	return nil
}
