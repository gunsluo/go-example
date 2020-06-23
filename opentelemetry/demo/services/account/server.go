package account

import (
	"encoding/json"
	"net/http"

	"github.com/gunsluo/go-example/opentelemetry/demo/pkg/storage"
	"go.uber.org/zap"
)

// Server implements jaeger-demo-frontend service
type Server struct {
	address string
	//tracer  opentracing.Tracer
	logger *zap.SugaredLogger

	database *storage.Database
}

// ConfigOptions used to make sure service clients
// can find correct server ports
type ConfigOptions struct {
	Address string
}

// NewServer creates a new frontend.Server
func NewServer(options ConfigOptions, logger *zap.Logger) *Server {
	//tracer := trace.Init("account", logger, nil)
	logger = logger.Named("account")
	return &Server{
		address:  options.Address,
		logger:   logger.Sugar(),
		database: storage.NewDatabase(logger),
		//tracer:   tracer,
	}
}

// Run starts the frontend server
func (s *Server) Run() error {
	mux := s.createServeMux()
	s.logger.With("address", s.address).Info("Starting Service")
	return http.ListenAndServe(s.address, mux)
}

func (s *Server) createServeMux() http.Handler {
	mux := http.NewServeMux()
	/*
		traceMiddleware := trace.NewHttpMiddleware(s.tracer, trace.WithHttpComponentName("Account Server"))
		mux.HandleFunc("/account", traceMiddleware.Handle(s.account))
	*/
	mux.HandleFunc("/account", s.account)
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
	s.logger.With("accountId", userID).Info("loading account from database")
	account, err := s.database.GetAccount(ctx, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.logger.With(zap.Error(err), "accountId", userID).Warn("failed to loading account from database")
		return
	}

	data, err := json.Marshal(account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
