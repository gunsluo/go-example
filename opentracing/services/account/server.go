package account

import (
	"encoding/json"
	"net/http"

	"github.com/gunsluo/go-example/opentracing/pkg/storage"
	"github.com/gunsluo/go-example/opentracing/pkg/trace"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Server implements jaeger-demo-frontend service
type Server struct {
	address string
	tracer  opentracing.Tracer
	logger  logrus.FieldLogger

	database *storage.Database
}

// ConfigOptions used to make sure service clients
// can find correct server ports
type ConfigOptions struct {
	Address string
}

// NewServer creates a new frontend.Server
func NewServer(options ConfigOptions, logger logrus.FieldLogger) *Server {
	tracer := trace.Init("account", logger, nil)
	return &Server{
		address:  options.Address,
		logger:   logger,
		database: storage.NewDatabase(logger),
		tracer:   tracer,
	}
}

// Run starts the frontend server
func (s *Server) Run() error {
	mux := s.createServeMux()
	s.logger.WithField("address", s.address).Info("Starting Service")
	return http.ListenAndServe(s.address, mux)
}

func (s *Server) createServeMux() http.Handler {
	mux := http.NewServeMux()
	traceMiddleware := trace.NewHttpMiddleware(s.tracer, trace.WithHttpComponentName("Account Server"))
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
	account, err := s.database.GetAccount(ctx, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
