package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gunsluo/go-example/opentracing/pkg/internal"
	"github.com/sirupsen/logrus"
)

// AccountClient is a remote client that implements customer.Interface
type AccountClient struct {
	logger     logrus.FieldLogger
	accountURL string
	client     *http.Client
	//tracer   opentracing.Tracer
	//client   *tracing.HTTPAccountClient
}

// NewAccountClient creates a new customer.AccountClient
func NewAccountClient(logger logrus.FieldLogger, accountURL string) *AccountClient {
	return &AccountClient{
		//tracer: tracer,
		logger:     logger,
		client:     &http.Client{},
		accountURL: accountURL,
	}
}

// Get implements customer.Interface#Get as an RPC
func (c *AccountClient) GetAccount(ctx context.Context, id string) (*internal.Account, error) {
	c.logger.WithField("account id", id).Info("Getting customer")

	url := fmt.Sprintf(c.accountURL+"/account?id=%s", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	//req = req.WithContext(ctx)

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
