package hub

// ClientReader is a client reader
type ClientReader interface {
	// call the method when the user send a request
	Read(ctx *Context, meesage []byte) error
}

var DefaultClientReader ClientReader = &TODOClientReader{}

// TODOClientReader is a todo client reader, nothing to do
type TODOClientReader struct {
}

// Read read message
func (r *TODOClientReader) Read(ctx *Context, message []byte) error {
	return nil
}

// PushlisherClientReader is pushlisher ClientReader, broadcast messages to clients.
type PushlisherClientReader struct {
}

func (r *PushlisherClientReader) Read(ctx *Context, message []byte) error {
	// pushlish the message to subscribe hub
	ctx.Repeater.Broadcast(message)

	// send response to client
	ctx.CurrentClient.Send(ctx, []byte("pushlish done"))

	return nil
}
