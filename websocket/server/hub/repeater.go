package hub

// Repeater is a repeater
type Repeater struct {
	h *Hub
}

// Broadcast broadcast messages to the clients.
func (r *Repeater) Broadcast(message []byte) error {
	r.h.broadcast <- message
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

func (r *Repeater) Hub() *Hub {
	return r.h
}
