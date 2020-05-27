package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/graph-gophers/dataloader"
)

func main() {
	// setup batch function
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result
		// do some async work to get data for specified keys
		// append to this list resolved values
		for _, key := range keys {
			fmt.Printf("->%s %v\n", key.String(), key.Raw())
			if "key" == key.String() {
				id, ok := key.Raw().(int)
				if ok {
					v := getData(id)
					results = append(results, &dataloader.Result{Data: v})
				} else {
					results = append(results, &dataloader.Result{Error: errors.New("invalid parameter")})
				}
			}
		}
		return results
	}

	cache := NewCache()
	// create Loader with an in-memory cache
	loader := dataloader.NewBatchedLoader(batchFn, dataloader.WithCache(cache))
	/**
	 * Use loader
	 *
	 * A thunk is a function returned from a function that is a
	 * closure over a value (in this case an interface value and error).
	 * When called, it will block until the value is resolved.
	 */

	ctx := context.Background()
	go func() {
		thunk := loader.Load(ctx, ValueKey{K: "key", V: 100}) // ValueKey is a convenience method that make wraps string to implement `Key` interface
		result, err := thunk()
		if err != nil {
			panic(err)
		}
		log.Printf("value: %#v", result)
	}()

	go func() {
		thunk := loader.Load(ctx, ValueKey{K: "key", V: 200}) // ValueKey is a convenience method that make wraps string to implement `Key` interface
		result, err := thunk()
		if err != nil {
			panic(err)
		}
		log.Printf("value: %#v", result)
	}()

	go func() {
		thunk := loader.Load(ctx, ValueKey{K: "key", V: 100}) // ValueKey is a convenience method that make wraps string to implement `Key` interface
		result, err := thunk()
		if err != nil {
			panic(err)
		}
		log.Printf("value: %#v", result)
	}()

	select {}
}

// ValueKey implements the Key interface for a string
type ValueKey struct {
	K string
	V interface{}
}

// String is an identity method. Used to implement String interface
func (v ValueKey) String() string { return v.K }

// Raw is an identity method. Used to implement Key Raw
func (v ValueKey) Raw() interface{} { return v.V }

func getData(id int) int {
	return id * 2
}

// Cache is an in memory implementation of Cache interface.
// This simple implementation is well suited for
// a "per-request" dataloader (i.e. one that only lives
// for the life of an http request) but it's not well suited
// for long lived cached items.
type Cache struct {
	items *sync.Map
}

// NewCache constructs a new Cache
func NewCache() *Cache {
	return &Cache{
		items: &sync.Map{},
	}
}

// Set sets the `value` at `key` in the cache
func (c *Cache) Set(_ context.Context, key dataloader.Key, value dataloader.Thunk) {
	c.items.Store(key, value)
}

// Get gets the value at `key` if it exsits, returns value (or nil) and bool
// indicating of value was found
func (c *Cache) Get(_ context.Context, key dataloader.Key) (dataloader.Thunk, bool) {
	item, found := c.items.Load(key)
	if !found {
		return nil, false
	}

	return item.(dataloader.Thunk), true
}

// Delete deletes item at `key` from cache
func (c *Cache) Delete(_ context.Context, key dataloader.Key) bool {
	if _, found := c.items.Load(key); found {
		c.items.Delete(key)
		return true
	}
	return false
}

// Clear clears the entire cache
func (c *Cache) Clear() {
	c.items.Range(func(key, _ interface{}) bool {
		c.items.Delete(key)
		return true
	})
}
