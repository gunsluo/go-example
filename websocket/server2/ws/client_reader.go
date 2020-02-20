package ws

import (
	"context"
)

// ClientReader is a client reader
type ClientReader interface {
	// call the method when the user send a request
	Read(ctx context.Context, chain *Chain, meesage []byte) error
}

var DefaultClientReader ClientReader = &TODOClientReader{}

// TODOClientReader is a todo client reader, nothing to do
type TODOClientReader struct {
}

// Read read message
func (r *TODOClientReader) Read(ctx context.Context, chain *Chain, message []byte) error {
	return nil
}

// PublisherClientReader is publisher ClientReader, broadcast messages to clients.
type PublisherClientReader struct {
}

func (r *PublisherClientReader) Read(ctx context.Context, chain *Chain, message []byte) error {
	// publish the message to subscribe hub
	chain.Repeater.Broadcast(message)

	// send response to client
	chain.CurrentClient.Send(ctx, []byte("publish done"))

	return nil
}
