package hashmap

type Map[K comparable, V comparable] struct {
	data map[K]V
}

func New[K comparable, V comparable]() *Map[K, V] {
	return &Map[K, V]{data: make(map[K]V)}
}
