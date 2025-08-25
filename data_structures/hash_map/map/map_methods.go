package hashmap

import "fmt"

func (m *Map[K, V]) Set(key K, value V) {
	m.data[key] = value
}

func (m Map[K, V]) Get(key K) (V, bool) {
	v, ok := m.data[key]
	return v, ok
}

func (m *Map[K, V]) Delete(key K) (V, bool) {
	if v, ok := m.data[key]; ok {
		delete(m.data, key)
		return v, true
	}
	var zero V
	return zero, false
}

func (m Map[K, V]) Contains(key K) bool {
	_, ok := m.data[key]
	return ok
}

func (m Map[K, V]) ContainsValue(value V) bool {
	for _, v := range m.data {
		if v == value {
			return true
		}
	}
	return false
}

func (m Map[K, V]) Size() int {
	return len(m.data)
}

func (m Map[K, V]) IsEmpty() bool {
	return len(m.data) == 0
}

func (m *Map[K, V]) Clear() {
	m.data = make(map[K]V)
}

func (m Map[K, V]) ToMap() map[K]V {
	out := make(map[K]V, len(m.data))
	for k, v := range m.data {
		out[k] = v
	}
	return out
}

func (m Map[K, V]) Keys() []K {
	out := make([]K, 0, len(m.data))
	for k := range m.data {
		out = append(out, k)
	}
	return out
}

func (m Map[K, V]) Values() []V {
	out := make([]V, 0, len(m.data))
	for _, v := range m.data {
		out = append(out, v)
	}
	return out
}

func (m Map[K, V]) Copy() *Map[K, V] {
	out := New[K, V]()
	for k, v := range m.data {
		out.data[k] = v
	}
	return out
}

func (m Map[K, V]) Traverse(fn func(K, V) bool) {
	for k, v := range m.data {
		if !fn(k, v) {
			return
		}
	}
}

func (m Map[K, V]) String() string {
	return fmt.Sprintf("%v", m.data)
}
