package heap

func (h *Heap[T]) Size() int     { return len(h.data) }
func (h *Heap[T]) IsEmpty() bool { return len(h.data) == 0 }
func (h *Heap[T]) Clear()        { *h = Heap[T]{less: h.less, heapType: h.heapType} }

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

func (h *Heap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap[T]) siftUp(index int) {
	if index <= 0 || index >= len(h.data) {
		return
	}
	for {
		parent := h.parentIndex(index)
		indexLessParent := h.less(h.data[index], h.data[parent])
		if parent == -1 || (h.heapType == minHeap && !indexLessParent) || (h.heapType == maxHeap && indexLessParent) {
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
		if right != -1 {
			leftLessRight := h.less(h.data[left], h.data[right])
			if (h.heapType == minHeap && !leftLessRight) || (h.heapType == maxHeap && leftLessRight) {
				child = right
			}
		}

		indexLessChild := h.less(h.data[index], h.data[child])
		if (h.heapType == minHeap && indexLessChild) || (h.heapType == maxHeap && !indexLessChild) {
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

func (h *Heap[T]) Contain(value T) bool {
	for _, v := range h.data {
		if !h.less(v, value) && !h.less(value, v) {
			return true
		}
	}
	return false
}
