package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/graph-gophers/dataloader"
)

func main() {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result
		// do some async work to get data for specified keys
		// append to this list resolved values
		var ids []int
		for _, key := range keys {
			//fmt.Printf("->%s %v\n", key.String(), key.Raw())
			id, ok := key.Raw().(int)
			if ok {
				ids = append(ids, id)
			}
		}

		users := getUserByIds(ids)
		// should be sort by keys
		for _, u := range users {
			results = append(results, &dataloader.Result{Data: u})
		}
		return results
	}

	cache := NewCache()
	// create Loader with an in-memory cache
	loader := dataloader.NewBatchedLoader(batchFn, dataloader.WithCache(cache))
	//loader := dataloader.NewBatchedLoader(batchFn)

	var wg sync.WaitGroup
	wg.Add(3)
	ctx := context.Background()
	go func() {
		thunk := loader.Load(ctx, IntKey{1})
		data, err := thunk()
		if err != nil {
			panic(err)
		}

		user, ok := data.(*User)
		if !ok {
			panic("invalid type")
		}

		//user := getUserById(1)
		log.Printf("key: %d value: %#v", 1, user)
		wg.Done()
	}()

	go func() {
		thunk := loader.Load(ctx, IntKey{2})
		data, err := thunk()
		if err != nil {
			panic(err)
		}

		user, ok := data.(*User)
		if !ok {
			panic("invalid type")
		}

		//user := getUserById(2)
		log.Printf("key: %d value: %#v", 2, user)
		wg.Done()
	}()

	go func() {
		thunk := loader.Load(ctx, IntKey{3})
		data, err := thunk()
		if err != nil {
			panic(err)
		}

		user, ok := data.(*User)
		if !ok {
			panic("invalid type")
		}
		//user := getUserById(3)
		log.Printf("key: %d value: %#v", 3, user)
		wg.Done()
	}()

	wg.Wait()
	fmt.Printf("call times: %d\n", callTimes)
}

type User struct {
	Id   int
	Name string
}

var mockUsers = []*User{
	&User{Id: 1, Name: "luoji"},
	&User{Id: 2, Name: "jerry"},
	&User{Id: 3, Name: "mary"},
	&User{Id: 4, Name: "lili"},
	&User{Id: 5, Name: "mark"},
}

var callTimes int

func getUserById(id int) *User {
	callTimes++
	for _, u := range mockUsers {
		if u.Id == id {
			return u
		}
	}

	return nil
}

func getUserByIds(ids []int) []*User {
	callTimes++
	var users []*User
	for _, id := range ids {
		for _, u := range mockUsers {
			if u.Id == id {
				users = append(users, u)
				break
			}
		}
	}

	return users
}

// IntKey implements the Key interface for a int
type IntKey struct {
	I int
}

func (i IntKey) String() string { return "ik" }

func (i IntKey) Raw() interface{} { return i.I }

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
