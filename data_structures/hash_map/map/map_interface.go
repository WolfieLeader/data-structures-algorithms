package hashmap

type HashMap[K comparable, V comparable] interface {
	Set(key K, value V)
	Get(key K) (V, bool)
	Delete(key K) (V, bool)
	Contains(key K) bool
	ContainsValue(value V) bool
	Size() int
	IsEmpty() bool
	Clear()
	ToMap() map[K]V
	Keys() []K
	Values() []V
	Traverse(func(key K, value V) bool)
	Copy() *Map[K, V]
	String() string
}

var _ HashMap[string, int] = (*Map[string, int])(nil)
