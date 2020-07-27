package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gunsluo/go-example/opentelemetry/demo/pkg/client"
	"github.com/gunsluo/go-example/opentelemetry/demo/pkg/otlp/trace"
	"github.com/gunsluo/go-example/opentelemetry/demo/pkg/otlp/trace/amqptrace"
	identitypb "github.com/gunsluo/go-example/opentelemetry/demo/pkg/proto/identity"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Server implements jaeger-demo-frontend service
type Server struct {
	address        string
	idmAddress     string
	logger         *zap.Logger
	accountClient  *client.AccountClient
	identityClient identitypb.IdmClient
	traceConfig    *trace.Configuration
	rabbitmqConn   *amqp.Connection
	rabbitmqCh     *amqptrace.Channel
}

// ConfigOptions used to make sure service clients
// can find correct server ports
type ConfigOptions struct {
	Address     string
	AccountURL  string
	IdmAddress  string
	RecordMQUrl string
}

// NewServer creates a new frontend.Server
func NewServer(options ConfigOptions, logger *zap.Logger) (*Server, error) {
	logger = logger.Named("backend")

	s := &Server{
		address:    options.Address,
		idmAddress: options.IdmAddress,
		logger:     logger,
	}

	// trace
	traceConfig, err := trace.FromEnv()
	if err != nil {
		return nil, fmt.Errorf("failed to loading trace config, %w", err)
	}
	traceConfig.ServiceName = "backend"
	s.traceConfig = traceConfig

	tracer, err := traceConfig.NewTracer(trace.ServiceName("backend"), trace.Logger(logger))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace, %w", err)
	}
	conn, err := grpc.Dial(s.idmAddress, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(trace.UnaryClientInterceptor(tracer, "Identity GRPC Client")),
	)
	if err != nil {
		return nil, err
	}
	s.identityClient = identitypb.NewIdmClient(conn)

	transport, err := traceConfig.NewTransport(
		trace.WithTransportComponentName("Account Http Client"),
		trace.WithTransportLogger(logger),
	)
	s.accountClient = client.NewAccountClient(
		logger,
		&http.Client{Transport: transport},
		options.AccountURL)

	// rabbitmq
	rabbitmqConn, err := amqp.Dial(options.RecordMQUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to rabbitmq, %s: %w", options.RecordMQUrl, err)
	}
	s.rabbitmqConn = rabbitmqConn
	s.logger.With(zap.String("mq-url", options.RecordMQUrl)).Info("success to connect to rabbitmq")

	rabbitmqCh, err := s.rabbitmqConn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to create to channel, %w", err)
	}
	s.rabbitmqCh = amqptrace.NewChannel(tracer, rabbitmqCh,
		amqptrace.ChannelComopenentName("Record Rabbitmq Publisher"))

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
		trace.WithHttpComponentName("Backend Http Server"),
		trace.WithHttpLogger(s.logger))
	if err != nil {
		s.logger.With(zap.Error(err)).Warn("failed to create trace http middleware")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/profile", traceMiddleware.Handle(s.profile))
	return mux
}

func withContxt(ctx context.Context) []zap.Field {
	return []zap.Field{trace.TraceId(ctx)}
}

func (s *Server) profile(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "Missing required 'id' parameter", http.StatusBadRequest)
		return
	}

	// query user from account server
	ctx := r.Context()
	s.logger.With(withContxt(ctx)...).With(zap.String("accountId", userID)).Info("query account by user id from account")
	account, err := s.accountClient.GetAccount(ctx, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.logger.With(withContxt(ctx)...).With(zap.Error(err)).With(zap.String("accountId", userID)).Warn("failed to query account")
		return
	}

	s.logger.With(withContxt(ctx)...).With(zap.String("userId", userID)).Info("query user by user id from identity")
	reply, err := s.identityClient.UserIdentity(ctx, &identitypb.UserIdentityRequest{Id: userID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.logger.With(withContxt(ctx)...).With(zap.Error(err)).With(zap.String("userId", userID)).Warn("failed to query user")
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

	err = s.rabbitmqCh.Publish(ctx, "", "test-record", false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		})
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
