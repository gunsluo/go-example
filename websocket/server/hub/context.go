package hub

import (
	"context"
	"time"
)

// Context is the context about websocket request
type Context struct {
	ctx context.Context

	CurrentClient *Client
	Repeater      *Repeater
}

// NewContext return a websocket context
func NewContext(h *Hub, client *Client) *Context {
	return &Context{
		ctx:           context.Background(),
		Repeater:      &Repeater{h: h},
		CurrentClient: client,
	}
}

// Deadline implment deadline from context.Context
func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.ctx.Deadline()
}

// Done implment done from context.Context
func (c *Context) Done() <-chan struct{} {
	//Done() <-chan struct{}
	return c.ctx.Done()
}

// Err implment err from context.Context
func (c *Context) Err() error {
	return c.ctx.Err()
}

func (c *Context) Value(key interface{}) interface{} {
	return c.ctx.Value(key)
}

// String implment string
func (*Context) String() string {
	return "hub.Context"
}
