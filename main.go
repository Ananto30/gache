package main

import (
	"fmt"
)

type MemoryStore[K comparable, V any] struct {
	Map map[K]V
}

func (m *MemoryStore[K, V]) Get(key K) (V, error) {
	var result V
	if v, ok := m.Map[key]; ok {
		return v, nil
	}
	return result, fmt.Errorf("key not found")
}

func (m *MemoryStore[K, V]) Set(key K, value V) {
	m.Map[key] = value
}

func (m *MemoryStore[K, V]) Delete(key K) {
	delete(m.Map, key)
}

func (m *MemoryStore[K, V]) Clear() {
	m.Map = map[K]V{}
}

func (m *MemoryStore[K, V]) Size() int {
	return len(m.Map)
}

func (m *MemoryStore[K, V]) Keys() []K {
	keys := make([]K, 0, len(m.Map))
	for k := range m.Map {
		keys = append(keys, k)
	}
	return keys
}

func (m *MemoryStore[K, V]) Values() []V {
	values := make([]V, 0, len(m.Map))
	for _, v := range m.Map {
		values = append(values, v)
	}
	return values
}

func (m *MemoryStore[K, V]) Entries() []struct {
	K K
	V V
} {
	entries := make([]struct {
		K K
		V V
	}, 0, len(m.Map))
	for k, v := range m.Map {
		entries = append(entries, struct {
			K K
			V V
		}{k, v})
	}
	return entries
}

func (m *MemoryStore[K, V]) Has(key K) bool {
	_, ok := m.Map[key]
	return ok
}

func (m *MemoryStore[K, V]) ForEach(f func(K, V)) {
	for k, v := range m.Map {
		f(k, v)
	}
}

func cached[T comparable, R any](cache MemoryStore[T, R], f func(T) (R, error)) func(T) (R, error) {
	return func(x T) (R, error) {
		if v, err := cache.Get(x); err == nil {
			fmt.Println("cache hit")
			return v, nil
		}
		v, err := f(x)
		if err != nil {
			return v, err
		}
		cache.Set(x, v)
		return v, nil
	}
}

type User struct {
	Name string
	Age  int
}

func add(x User) (User, error) {
	return x, nil
}

func main() {
	println("Hello, world!")
	memoryStore := MemoryStore[User, User]{Map: map[User]User{}}

	// cache wrapper
	addCached := cached(memoryStore, add)

	user := User{Name: "test", Age: 10}
	user2 := User{Name: "test", Age: 10}

	addCached(user)
	addCached(user2)
	addCached(user2)

}
