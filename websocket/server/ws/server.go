package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/gunsluo/go-example/websocket/server/hub"
)

const (
	SUBSCRIBE = "subscribe"
	PUSHLISH  = "pushlish"
)

// Option is option
type Option struct {
	Address string
}

// Server is a websocket server
type Server struct {
	address string

	upgrader websocket.Upgrader
	h        *hub.Hub
}

// NewServer return a websocket server
func NewServer(option Option) *Server {
	return &Server{
		address: option.Address,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		h: hub.New(),
	}
}

// Handler is handler responds to handles websocket requests
func (s *Server) subscribe(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	var client *hub.Client
	client = hub.NewClient(conn, hub.DefaultClientReader)
	s.h.Register(client)

	ctx := hub.NewContext(s.h, client)
	go client.Write(ctx)
	go client.Read(ctx)
}

// Handler is handler responds to handles websocket requests
func (s *Server) pushlish(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	cr := &hub.PushlisherClientReader{}
	client := hub.NewClient(conn, cr)

	ctx := hub.NewContext(s.h, client)
	go client.Write(ctx)
	go client.Read(ctx)
}

func (s *Server) Run() {
	http.HandleFunc("/"+SUBSCRIBE, func(w http.ResponseWriter, r *http.Request) {
		s.subscribe(w, r)
	})

	http.HandleFunc("/"+PUSHLISH, func(w http.ResponseWriter, r *http.Request) {
		s.pushlish(w, r)
	})

	go s.h.Run()

	log.Println("start up websocket server, listen on", s.address)
	err := http.ListenAndServe(s.address, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
