package idm

import (
	"context"
	"net"

	identitypb "github.com/gunsluo/go-example/opentelemetry/demo/pkg/proto/identity"
	"github.com/gunsluo/go-example/opentelemetry/demo/pkg/storage"
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
}

// NewServer creates a new frontend.Server
func NewServer(options ConfigOptions, logger *zap.Logger) *Server {
	//tracer := trace.Init("idm", logger, nil)
	logger = logger.Named("idm")
	return &Server{
		address:  options.Address,
		logger:   logger,
		database: storage.NewDatabase(logger),
		//tracer:   tracer,
	}
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
