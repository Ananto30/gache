package store

import (
	"fmt"
	"sync"
)

type MemoryStore[K comparable, V any] struct {
	Map sync.Map
}

func (m *MemoryStore[K, V]) Get(key K) (V, error) {
	var result V
	if v, ok := m.Map.Load(key); ok {
		result = v.(V)
		return result, nil
	}
	return result, fmt.Errorf("key not found")
}

func (m *MemoryStore[K, V]) Set(key K, value V) {
	m.Map.Store(key, value)
}

func (m *MemoryStore[K, V]) Delete(key K) {
	m.Map.Delete(key)
}

func (m *MemoryStore[K, V]) Clear() {
	m.Map.Range(func(key, value interface{}) bool {
		m.Map.Delete(key)
		return true
	})
}

func (m *MemoryStore[K, V]) Size() int {
	var size int
	m.Map.Range(func(key, value interface{}) bool {
		size++
		return true
	})
	return size
}

func (m *MemoryStore[K, V]) Keys() []K {
	var keys []K
	m.Map.Range(func(key, value interface{}) bool {
		keys = append(keys, key.(K))
		return true
	})
	return keys
}

func (m *MemoryStore[K, V]) Values() []V {
	var values []V
	m.Map.Range(func(key, value interface{}) bool {
		values = append(values, value.(V))
		return true
	})
	return values
}

func (m *MemoryStore[K, V]) Entries() []struct {
	K K
	V V
} {
	var entries []struct {
		K K
		V V
	}
	m.Map.Range(func(key, value interface{}) bool {
		entries = append(entries, struct {
			K K
			V V
		}{key.(K), value.(V)})
		return true
	})

	return entries
}

func (m *MemoryStore[K, V]) Has(key K) bool {
	_, ok := m.Map.Load(key)
	return ok
}
