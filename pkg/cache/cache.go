package cache

import (
	"github.com/ananto30/gache/pkg/store"
)

func Cached[T comparable, R any](cache store.MemoryStore[T, R], f func(T) (R, error)) func(T) (R, error) {
	return func(x T) (R, error) {
		if v, err := cache.Get(x); err == nil {
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
