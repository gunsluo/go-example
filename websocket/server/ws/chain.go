package ws

// Chain is request chain of a websocket request
type Chain struct {
	CurrentClient *Client
	Repeater      *Repeater
}

// NewChain return a chain of websocket context
func NewChain(r *Repeater, client *Client) *Chain {
	return &Chain{
		Repeater:      r,
		CurrentClient: client,
	}
}
