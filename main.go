package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/ananto30/gache/pkg/cache"
	"github.com/ananto30/gache/pkg/retry"
	"github.com/ananto30/gache/pkg/store"
)

type User struct {
	Name string
	Age  int
}

func add(x User) (User, error) {
	return x, nil
}

func fail(ctx context.Context, x User) (User, error) {
	return x, fmt.Errorf("failed")
}

func main() {
	memoryStore := store.MemoryStore[User, User]{Map: sync.Map{}}

	// cache wrapper
	addCached := cache.Cached(memoryStore, add)

	user := User{Name: "test", Age: 10}
	user2 := User{Name: "test", Age: 10}

	fmt.Println(addCached(user))
	fmt.Println(addCached(user2))
	fmt.Println(addCached(user))
	fmt.Println(addCached(user2))

	fmt.Println(retry.Func(fail).Do(user))
}
