package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gunsluo/go-example/opentelemetry/demo/pkg/internal"
	"go.uber.org/zap"
)

// AccountClient is a remote client that implements customer.Interface
type AccountClient struct {
	logger     *zap.SugaredLogger
	accountURL string
	client     *http.Client
	//tracer   opentracing.Tracer
	//client   *tracing.HTTPAccountClient
}

// NewAccountClient creates a new customer.AccountClient
func NewAccountClient(logger *zap.Logger, accountURL string) *AccountClient {
	return &AccountClient{
		logger: logger.Sugar(),
		/*
			client: &http.Client{
				Transport: trace.NewTransport(tracer,
					trace.TransportComponentName("Customer Client"))},
		*/
		client:     &http.Client{},
		accountURL: accountURL,
	}
}

// Get implements customer.Interface#Get as an RPC
func (c *AccountClient) GetAccount(ctx context.Context, id string) (*internal.Account, error) {
	c.logger.With("account id", id).Info("Getting customer")

	url := fmt.Sprintf(c.accountURL+"/account?id=%s", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	// important
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(body))
	}

	var a internal.Account
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&a)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
