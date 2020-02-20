package ws

// Repeater is a repeater
type Repeater struct {
	h *Hub

	// Inbound messages from the clients.
	delivery chan []byte
}

// NewRepeater return a repeater
func NewRepeater(h *Hub) *Repeater {
	return &Repeater{h: h, delivery: make(chan []byte)}
}

// Broadcast broadcast messages to the clients.
func (r *Repeater) Broadcast(message []byte) error {
	r.delivery <- message
	return nil
}

// Unicast send messages to a client.
func (r *Repeater) Unicast(message []byte) error {
	return nil
}

// Multicast send messages to the clients of a group.
func (r *Repeater) Multicast(message []byte) error {
	return nil
}

// Hub get clients hub
func (r *Repeater) Hub() *Hub {
	return r.h
}

// Run start receiving channel
func (r *Repeater) Run() {
	go r.h.Run()

	for {
		select {
		case message := <-r.delivery:
			r.broadcast(message)
		}
	}
}

// broadcast send the message to all clients
func (r *Repeater) broadcast(message []byte) error {
	// TODO: 1. multithreading 2. asynchronous sending
	h := r.Hub()
	for client := range h.clients {
		select {
		case client.send <- message:
		default:
			close(client.send)
			delete(h.clients, client)
		}
	}

	return nil
}
