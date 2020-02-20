package server

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/gunsluo/go-example/websocket/server/ws"
	"github.com/sirupsen/logrus"
)

const (
	SUBSCRIBE = "subscribe"
	PUSHLISH  = "pushlish"
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

	ctx := ws.NewContext(s.repeater, client)
	go client.Write(ctx)
	go client.Read(ctx)
}

// Handler is handler responds to handles websocket requests
func (s *Server) pushlish(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.logger.WithError(err).Warnln("unable to create websocket connection")
		return
	}

	cr := &ws.PushlisherClientReader{}
	client := ws.NewClient(s.logger, conn, cr)

	ctx := ws.NewContext(s.repeater, client)
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

	go s.repeater.Run()

	s.logger.Infoln("start up websocket server, listen on", s.address)
	err := http.ListenAndServe(s.address, nil)
	if err != nil {
		s.logger.WithError(err).Warnln("unable to start up websocket server")
	}
}
