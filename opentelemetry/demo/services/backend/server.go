package backend

import (
	"encoding/json"
	"net/http"

	"github.com/gunsluo/go-example/opentelemetry/demo/pkg/client"
	identitypb "github.com/gunsluo/go-example/opentelemetry/demo/pkg/proto/identity"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Server implements jaeger-demo-frontend service
type Server struct {
	address        string
	idmAddress     string
	logger         *zap.SugaredLogger
	accountClient  *client.AccountClient
	identityClient identitypb.IdmClient
	//tracer         opentracing.Tracer
}

// ConfigOptions used to make sure service clients
// can find correct server ports
type ConfigOptions struct {
	Address    string
	AccountURL string
	IdmAddress string
}

// NewServer creates a new frontend.Server
func NewServer(options ConfigOptions, logger *zap.Logger) *Server {
	//tracer := trace.Init("backend", logger, nil)
	return &Server{
		address:       options.Address,
		idmAddress:    options.IdmAddress,
		logger:        logger.Sugar(),
		accountClient: client.NewAccountClient(logger, options.AccountURL),
		//tracer:        tracer,
	}
}

// Run starts the frontend server
func (s *Server) Run() error {
	identityClient, err := createIdmClient(s.idmAddress)
	if err != nil {
		return err
	}
	s.identityClient = identityClient

	mux := s.createServeMux()
	s.logger.With("address", s.address).Info("Starting Service")
	return http.ListenAndServe(s.address, mux)
}

func (s *Server) createServeMux() http.Handler {
	mux := http.NewServeMux()
	/*
		traceMiddleware := trace.NewHttpMiddleware(s.tracer,
			trace.WithHttpComponentName("Backend Server"),
			trace.WithHttpLogger(s.logger))
		mux.HandleFunc("/profile", traceMiddleware.Handle(s.profile))
	*/
	mux.HandleFunc("/profile", s.profile)
	return mux
}

func (s *Server) profile(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "Missing required 'id' parameter", http.StatusBadRequest)
		return
	}

	// query user from account server
	ctx := r.Context()
	s.logger.Info("this is a test")
	account, err := s.accountClient.GetAccount(ctx, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reply, err := s.identityClient.UserIdentity(ctx, &identitypb.UserIdentityRequest{Id: userID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u := user{
		Id:     account.Id,
		Name:   account.Name,
		Email:  account.Email,
		CertId: reply.Identity.CertId,
	}
	data, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

type user struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	CertId string `json:"cert_id"`
}

func createIdmClient(address string) (identitypb.IdmClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return identitypb.NewIdmClient(conn), nil
}
