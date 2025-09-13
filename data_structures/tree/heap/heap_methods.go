package heap

import (
	"fmt"
	"strings"
)

func (h *Heap[T]) Size() int     { return len(h.data) }
func (h *Heap[T]) IsEmpty() bool { return len(h.data) == 0 }
func (h *Heap[T]) Clear()        { *h = Heap[T]{aboveFn: h.aboveFn} }

func (h *Heap[T]) Peek() (T, bool) {
	var zero T
	if len(h.data) == 0 {
		return zero, false
	}
	return h.data[0], true
}

func (h *Heap[T]) parentIndex(index int) int {
	if index <= 0 || index >= len(h.data) {
		return -1
	}
	return (index - 1) / 2
}

func (h *Heap[T]) leftChildIndex(index int) int {
	if index < 0 || index >= len(h.data) {
		return -1
	}
	left := (2 * index) + 1
	if left >= len(h.data) {
		return -1
	}
	return left
}

func (h *Heap[T]) rightChildIndex(index int) int {
	if index < 0 || index >= len(h.data) {
		return -1
	}
	right := (2 * index) + 2
	if right >= len(h.data) {
		return -1
	}
	return right
}

func (h *Heap[T]) lastParentIndex() int {
	if len(h.data) == 0 {
		return -1
	}
	return (len(h.data) / 2) - 1
}

func (h *Heap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap[T]) siftUp(index int) {
	if index <= 0 || index >= len(h.data) {
		return
	}
	for {
		parent := h.parentIndex(index)
		if parent == -1 || h.aboveFn(h.data[parent], h.data[index]) {
			break
		}
		h.swap(index, parent)
		index = parent
	}
}

func (h *Heap[T]) siftDown(index int) {
	if index < 0 || index >= len(h.data) {
		return
	}
	for {
		left := h.leftChildIndex(index)
		if left == -1 {
			break
		}
		child := left

		right := h.rightChildIndex(index)
		if right != -1 && h.aboveFn(h.data[right], h.data[left]) {
			child = right
		}
		if h.aboveFn(h.data[index], h.data[child]) {
			break
		}
		h.swap(index, child)
		index = child
	}
}

func (h *Heap[T]) Pop() (T, bool) {
	var zero T
	if len(h.data) == 0 {
		return zero, false
	}
	value := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.siftDown(0)
	return value, true
}

func (h *Heap[T]) Push(values ...T) {
	for _, value := range values {
		h.data = append(h.data, value)
		h.siftUp(len(h.data) - 1)
	}
}

func (h *Heap[T]) Contains(value T) bool {
	for _, v := range h.data {
		if !h.aboveFn(v, value) && !h.aboveFn(value, v) {
			return true
		}
	}
	return false
}

func (h *Heap[T]) Heapify(values ...T) {
	h.data = append(h.data[:0], values...)
	for i := h.lastParentIndex(); i >= 0; i-- {
		h.siftDown(i)
	}
}

func (h *Heap[T]) Copy() *Heap[T] {
	out := &Heap[T]{aboveFn: h.aboveFn, data: make([]T, len(h.data))}
	copy(out.data, h.data)
	return out
}

func (h *Heap[T]) ToSlice() []T {
	out := make([]T, len(h.data))
	copy(out, h.data)
	return out
}

func (h *Heap[T]) TraverseBreadthFirst(fn func(value T)) {
	for _, v := range h.data {
		fn(v)
	}
}

func (h *Heap[T]) Equal(other *Heap[T]) bool {
	if h == other {
		return true
	}
	if h == nil || other == nil {
		return false
	}
	if len(h.data) != len(other.data) {
		return false
	}
	for i := 0; i < len(h.data); i++ {
		if !h.aboveFn(h.data[i], other.data[i]) && !h.aboveFn(other.data[i], h.data[i]) {
			return false
		}
	}
	return true
}

func (h *Heap[T]) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "Heap{len=%d}\n", len(h.data))
	if len(h.data) == 0 {
		return sb.String()
	}
	h.drawIndex(&sb, "", 0, true, 0)
	return sb.String()
}

func (h *Heap[T]) drawIndex(sb *strings.Builder, prefix string, level int, isTail bool, i int) {
	if i < 0 || i >= len(h.data) {
		return
	}

	sb.WriteString(prefix)
	switch {
	case prefix == "":
		sb.WriteString("└── ")
	case isTail:
		sb.WriteString("└── ")
	default:
		sb.WriteString("├── ")
	}

	if prefix == "" {
		fmt.Fprintf(sb, "%v[%d]\n", h.data[i], i)
	} else {
		fmt.Fprintf(sb, "%v[%d](%d)\n", h.data[i], i, level)
	}

	right := 2*i + 2
	left := 2*i + 1

	children := make([]int, 0, 2)
	if right < len(h.data) {
		children = append(children, right)
	}
	if left < len(h.data) {
		children = append(children, left)
	}

	for k, c := range children {
		var nextPrefix string
		switch {
		case prefix == "":
			nextPrefix = "    "
		case isTail:
			nextPrefix = prefix + "    "
		default:
			nextPrefix = prefix + "│   "
		}
		h.drawIndex(sb, nextPrefix, level+1, k == len(children)-1, c)
	}
}
