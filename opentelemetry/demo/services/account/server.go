package account

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	ometric "github.com/gunsluo/go-example/opentelemetry/demo/pkg/otlp/metric"
	"github.com/gunsluo/go-example/opentelemetry/demo/pkg/otlp/trace"
	"github.com/gunsluo/go-example/opentelemetry/demo/pkg/storage"
	"go.opentelemetry.io/otel/api/kv"
	"go.opentelemetry.io/otel/api/metric"
	"go.uber.org/zap"
)

// Server implements jaeger-demo-frontend service
type Server struct {
	address string
	//tracer  opentracing.Tracer
	logger *zap.Logger

	database    *storage.Database
	traceConfig *trace.Configuration

	accountReadCounter metric.BoundInt64Counter
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

	db, err := traceConfig.OpenDB(options.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to connect db, %w", err)
	}

	database, err := storage.NewDatabase(logger, db)
	if err != nil {
		return nil, fmt.Errorf("failed to create database, %w", err)
	}
	s.database = database

	// metric
	metricConfig, err := ometric.FromEnv()
	if err != nil {
		return nil, fmt.Errorf("failed to loading metric config, %w", err)
	}
	meter, err := metricConfig.NewMeter("account", ometric.Logger(logger))
	if err != nil {
		return nil, fmt.Errorf("failed to create metric, %w", err)
	}

	counter, err := meter.NewInt64Counter("account.read")
	if err != nil {
		return nil, fmt.Errorf("failed to create metric recorder, %w", err)
	}

	commonLabels := []kv.KeyValue{
		kv.String("priority", "Ultra"),
	}
	s.accountReadCounter = counter.Bind(commonLabels...)
	//defer s.accountReadCounter.Unbind()

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

func withContxt(ctx context.Context) []zap.Field {
	return []zap.Field{trace.TraceId(ctx)}
}

func (s *Server) account(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "Missing required 'id' parameter", http.StatusBadRequest)
		return
	}

	// query user from account server
	ctx := r.Context()
	s.logger.With(withContxt(ctx)...).With(zap.String("accountId", userID)).Info("loading account from database")
	account, err := s.database.GetAccount(ctx, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.logger.With(withContxt(ctx)...).With(zap.Error(err), zap.String("accountId", userID)).Warn("failed to loading account from database")
		return
	}

	data, err := json.Marshal(account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.accountReadCounter.Add(ctx, 1)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
