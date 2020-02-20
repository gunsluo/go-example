package server

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/gunsluo/go-example/websocket/server2/ws"
	"github.com/sirupsen/logrus"
	"nhooyr.io/websocket"
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

	repeater *ws.Repeater

	logger logrus.FieldLogger
}

// New return a websocket server
func New(cfg Config) *Server {
	return &Server{
		address:  cfg.Address,
		repeater: ws.NewRepeater(ws.NewHub()),
		logger:   cfg.Logger,
	}
}

// subscribe is handler responds to handles websocket requests
func (s *Server) subscribe(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, nil)
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

// publish is handler responds to handles websocket requests
func (s *Server) publish(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, nil)
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
	l, err := net.Listen("tcp", s.address)
	if err != nil {
		s.logger.WithError(err).Fatalln("unable to listen on", s.address)
	}
	s.logger.Infoln("start up websocket server, listen on", l.Addr())

	m := http.NewServeMux()
	m.HandleFunc("/"+SUBSCRIBE, s.subscribe)
	m.HandleFunc("/"+PUBLISH, s.publish)

	httpServer := http.Server{
		Handler:      m,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	go s.repeater.Run()

	err = httpServer.Serve(l)
	if err != nil {
		s.logger.WithError(err).Fatalln("unable to start up websocket server")
	}
}
