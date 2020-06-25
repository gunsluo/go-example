package idm

import (
	"context"
	"fmt"
	"net"

	identitypb "github.com/gunsluo/go-example/opentelemetry/demo/pkg/proto/identity"
	"github.com/gunsluo/go-example/opentelemetry/demo/pkg/storage"
	"github.com/jmoiron/sqlx"
	"github.com/xo/dburl"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server implements jaeger-demo-frontend service
type Server struct {
	address string
	//tracer  opentracing.Tracer
	logger *zap.Logger

	database *storage.Database
}

// ConfigOptions used to make sure service clients
// can find correct server ports
type ConfigOptions struct {
	Address string
	DSN     string
}

// NewServer creates a new frontend.Server
func NewServer(options ConfigOptions, logger *zap.Logger) (*Server, error) {
	logger = logger.Named("idm")
	s := &Server{
		address: options.Address,
		logger:  logger,
	}

	u, err := dburl.Parse(options.DSN)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse database address, %w", err)
	}
	db, err := sqlx.Open(u.Driver, u.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to connect db, %w", err)
	}

	database, err := storage.NewDatabase(logger, db)
	if err != nil {
		return nil, fmt.Errorf("failed to create database, %w", err)
	}
	s.database = database

	return s, nil
}

// Run starts the frontend server
func (s *Server) Run() error {
	srv := grpc.NewServer()
	identitypb.RegisterIdmServer(srv, s)
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	s.logger.With(zap.String("address", s.address)).Info("Starting GRPC Service")
	return srv.Serve(listener)
}

func (s *Server) UserIdentity(ctx context.Context, req *identitypb.UserIdentityRequest) (*identitypb.UserIdentityReply, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "Missing required 'id' parameter")
	}

	// query user from account server
	s.logger.With(zap.String("userId", req.Id)).Info("loading identity from database")
	identity, err := s.database.GetIdentity(ctx, req.Id)
	if err != nil {
		s.logger.With(zap.Error(err), zap.String("userId", req.Id)).Warn("failed to loading identity from database")
		return nil, err
	}

	return &identitypb.UserIdentityReply{Identity: identity}, nil
}
