package storage

import (
	"sync"
)

const (
	GraphqlContextKey = "graphqlKey"
)

type Context struct {
	Values  map[string]interface{}
	Loaders DataLoaders
	lock    sync.RWMutex
}

func (c *Context) WithValue(key string, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.Values[key] = value
}

func (c *Context) Value(key string) interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.Values[key]
}
