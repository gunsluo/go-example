package server

import (
	"context"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/gunsluo/go-example/websocket/server/ws"
	"github.com/sirupsen/logrus"
)

const (
	SUBSCRIBE = "subscribe"
	PUBLISH   = "publish"
)

// Config is config
type Config struct {
	Address string

	Logger logrus.FieldLogger
}

// Server is a websocket server
type Server struct {
	address string

	upgrader websocket.Upgrader
	repeater *ws.Repeater

	logger logrus.FieldLogger
}

// New return a websocket server
func New(cfg Config) *Server {
	return &Server{
		address: cfg.Address,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				// allow all origin
				// TODO: it should be configurable
				return true
			},
		},
		repeater: ws.NewRepeater(ws.NewHub()),
		logger:   cfg.Logger,
	}
}

// Handler is handler responds to handles websocket requests
func (s *Server) subscribe(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.logger.WithError(err).Warnln("unable to create websocket connection")
		return
	}

	var client *ws.Client
	client = ws.NewClient(s.logger, conn, ws.DefaultClientReader)
	s.repeater.Hub().Register(client)

	ctx := context.Background()
	chain := ws.NewChain(s.repeater, client)
	go client.Write(ctx, chain)
	go client.Read(ctx, chain)
}

// Handler is handler responds to handles websocket requests
func (s *Server) publish(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.logger.WithError(err).Warnln("unable to create websocket connection")
		return
	}

	cr := &ws.PublisherClientReader{}
	client := ws.NewClient(s.logger, conn, cr)

	ctx := context.Background()
	chain := ws.NewChain(s.repeater, client)
	go client.Write(ctx, chain)
	go client.Read(ctx, chain)
}

func (s *Server) Run() {
	http.HandleFunc("/"+SUBSCRIBE, func(w http.ResponseWriter, r *http.Request) {
		s.subscribe(w, r)
	})

	http.HandleFunc("/"+PUBLISH, func(w http.ResponseWriter, r *http.Request) {
		s.publish(w, r)
	})

	go s.repeater.Run()

	s.logger.Infoln("start up websocket server, listen on", s.address)
	err := http.ListenAndServe(s.address, nil)
	if err != nil {
		s.logger.WithError(err).Warnln("unable to start up websocket server")
	}
}
