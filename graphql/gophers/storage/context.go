package storage

import "github.com/graph-gophers/dataloader"

const (
	GraphqlContextKey = "graphqlKey"
)

type Context struct {
	Values map[string]interface{}
	Loader *dataloader.Loader
}

func (c *Context) WithValue(key string, value interface{}) {
	c.Values[key] = value
}

func (c *Context) Value(key string) interface{} {
	return c.Values[key]
}
