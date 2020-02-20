package ws

import (
	"context"

	"github.com/sirupsen/logrus"
	"nhooyr.io/websocket"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	// Maximum message size allowed from peer.
	maxMessageSize int64

	reader ClientReader

	logger logrus.FieldLogger
}

// NewClient return a new client
func NewClient(logger logrus.FieldLogger, conn *websocket.Conn, reader ClientReader) *Client {
	return &Client{
		conn:           conn,
		send:           make(chan []byte, 1024),
		maxMessageSize: 2048,
		reader:         reader,
		logger:         logger,
	}
}

// Read pumps messages from the websocket connection to the hub.
//
// The application runs Read in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) Read(ctx context.Context, chain *Chain) {
	defer func() {
		h := chain.Repeater.Hub()
		h.Unregister(c)
		c.conn.CloseRead(ctx)
	}()
	c.conn.SetReadLimit(c.maxMessageSize)
	for {
		//err := wsjson.Read(ctx, c.conn, &v)
		_, message, err := c.conn.Read(ctx)
		if err != nil {
			if websocket.CloseStatus(err) != websocket.StatusNormalClosure {
				c.logger.WithError(err).Warnln("unable to read message")
			} else {
				c.logger.Infoln("xx client is closed")
			}
			break
		}
		// TODO: unmarshaml message
		// update group information of client to hub
		c.reader.Read(ctx, chain, message)
	}
}

// Write pumps messages from the hub to the websocket connection.
//
// A goroutine running Write is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) Write(ctx context.Context, chain *Chain) {
	defer func() {
		c.conn.Close(websocket.StatusNoStatusRcvd, "client is closed")
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				// The hub closed the channel.
				return
			}

			err := c.conn.Write(ctx, websocket.MessageText, message)
			if err != nil {
				return
			}

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				err := c.conn.Write(ctx, websocket.MessageText, <-c.send)
				if err != nil {
					return
				}
			}
		}
	}
}

// Send send messages to client
func (c *Client) Send(ctx context.Context, message []byte) {
	c.send <- message
}
