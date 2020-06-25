package account

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gunsluo/go-example/opentelemetry/demo/pkg/otlp/trace"
	"github.com/gunsluo/go-example/opentelemetry/demo/pkg/storage"
	"github.com/jmoiron/sqlx"
	"github.com/xo/dburl"
	"go.uber.org/zap"
)

// Server implements jaeger-demo-frontend service
type Server struct {
	address string
	//tracer  opentracing.Tracer
	logger *zap.Logger

	database    *storage.Database
	traceConfig *trace.Configuration
}

// ConfigOptions used to make sure service clients
// can find correct server ports
type ConfigOptions struct {
	Address string
	DSN     string
}

// NewServer creates a new frontend.Server
func NewServer(options ConfigOptions, logger *zap.Logger) (*Server, error) {
	logger = logger.Named("account")
	s := &Server{
		address: options.Address,
		logger:  logger,
	}

	// trace
	traceConfig, err := trace.FromEnv()
	if err != nil {
		return nil, fmt.Errorf("failed to loading trace config, %w", err)
	}
	traceConfig.ServiceName = "account"
	s.traceConfig = traceConfig

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
	mux := s.createServeMux()
	s.logger.With(zap.String("address", s.address)).Info("Starting Service")
	return http.ListenAndServe(s.address, mux)
}

func (s *Server) createServeMux() http.Handler {
	traceMiddleware, err := s.traceConfig.NewHttpMiddleware(
		trace.WithHttpComponentName("Account Http Server"),
		trace.WithHttpLogger(s.logger))
	if err != nil {
		s.logger.With(zap.Error(err)).Warn("failed to create trace http middleware")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/account", traceMiddleware.Handle(s.account))
	return mux
}

func (s *Server) account(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "Missing required 'id' parameter", http.StatusBadRequest)
		return
	}

	// query user from account server
	ctx := r.Context()
	s.logger.With(zap.String("accountId", userID)).Info("loading account from database")
	account, err := s.database.GetAccount(ctx, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.logger.With(zap.Error(err), zap.String("accountId", userID)).Warn("failed to loading account from database")
		return
	}

	data, err := json.Marshal(account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
