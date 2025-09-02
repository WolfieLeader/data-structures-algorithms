package hashset

import (
	"fmt"
	"maps"
	"sort"
	"strings"
)

func (m *Set[T]) Size() int     { return len(m.data) }
func (m *Set[T]) IsEmpty() bool { return len(m.data) == 0 }
func (m *Set[T]) Clear()        { m.data = make(map[T]struct{}) }

func (m *Set[T]) Add(values ...T) {
	for _, value := range values {
		m.data[value] = struct{}{}
	}
}

func (m *Set[T]) Delete(value T) bool {
	if _, ok := m.data[value]; ok {
		delete(m.data, value)
		return true
	}
	return false
}

func (m *Set[T]) Contains(value T) bool {
	_, ok := m.data[value]
	return ok
}

func (m *Set[T]) ToSlice() []T {
	out := make([]T, 0, len(m.data))
	for k := range m.data {
		out = append(out, k)
	}
	return out
}

func (m *Set[T]) Equal(other *Set[T]) bool {
	if m == other {
		return true
	}
	if m == nil || other == nil {
		return false
	}
	return maps.Equal(m.data, other.data)
}

func (m *Set[T]) Traverse(fn func(value T) bool) {
	for k := range m.data {
		if !fn(k) {
			return
		}
	}
}

func (m *Set[T]) Copy() *Set[T] {
	out := New[T]()
	maps.Copy(out.data, m.data)
	return out
}

func (m *Set[T]) String() string {
	if m.Size() == 0 {
		return "[]"
	}

	values := make([]string, 0, len(m.data))
	for k := range m.data {
		values = append(values, fmt.Sprintf("%v", k))
	}

	sort.Strings(values)
	return "[" + strings.Join(values, " ") + "]"
}
