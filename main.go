package main

import (
	"github.com/ananto30/gache/pkg/cache"
	"github.com/ananto30/gache/pkg/store"
)

type User struct {
	Name string
	Age  int
}

func add(x User) (User, error) {
	return x, nil
}

func main() {
	println("Hello, world!")
	memoryStore := store.MemoryStore[User, User]{Map: map[User]User{}}

	// cache wrapper
	addCached := cache.Cached(memoryStore, add)

	user := User{Name: "test", Age: 10}
	user2 := User{Name: "test", Age: 10}

	addCached(user)
	addCached(user2)
	addCached(user2)
}
