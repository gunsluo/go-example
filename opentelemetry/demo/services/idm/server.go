package idm

import (
	"context"
	"fmt"
	"net"

	"github.com/gunsluo/go-example/opentelemetry/demo/pkg/otlp/trace"
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
	grpcSrv  *grpc.Server
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

	// trace
	traceConfig, err := trace.FromEnv()
	if err != nil {
		return nil, fmt.Errorf("failed to loading trace config, %w", err)
	}
	traceConfig.ServiceName = "idm"

	tracer, err := traceConfig.NewTracer(trace.ServiceName("idm"))
	if err != nil {
		return nil, fmt.Errorf("failed to create tracer, %w", err)
	}

	s.grpcSrv = grpc.NewServer(
		grpc.UnaryInterceptor(trace.UnaryServerInterceptor(tracer, "Identity GRPC server")))

	db, err := trace.OpenDB(tracer, options.DSN)
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
	identitypb.RegisterIdmServer(s.grpcSrv, s)
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	s.logger.With(zap.String("address", s.address)).Info("Starting GRPC Service")
	return s.grpcSrv.Serve(listener)
}

func withContxt(ctx context.Context) []zap.Field {
	return []zap.Field{trace.TraceId(ctx)}
}

func (s *Server) UserIdentity(ctx context.Context, req *identitypb.UserIdentityRequest) (*identitypb.UserIdentityReply, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "Missing required 'id' parameter")
	}

	// query user from account server
	s.logger.With(withContxt(ctx)...).With(zap.String("userId", req.Id)).Info("loading identity from database")
	identity, err := s.database.GetIdentity(ctx, req.Id)
	if err != nil {
		s.logger.With(withContxt(ctx)...).With(zap.Error(err), zap.String("userId", req.Id)).Warn("failed to loading identity from database")
		return nil, err
	}

	return &identitypb.UserIdentityReply{Identity: identity}, nil
}
