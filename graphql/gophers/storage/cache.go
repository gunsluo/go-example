package storage

import (
	"context"
	"sync"

	"github.com/graph-gophers/dataloader"
)

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
