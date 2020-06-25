package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gunsluo/go-example/opentelemetry/demo/pkg/internal"
	"go.uber.org/zap"
)

// AccountClient is a remote client that implements customer.Interface
type AccountClient struct {
	logger     *zap.Logger
	accountURL string
	client     *http.Client
}

// NewAccountClient creates a new customer.AccountClient
func NewAccountClient(logger *zap.Logger, client *http.Client, accountURL string) *AccountClient {
	return &AccountClient{
		logger:     logger,
		client:     client,
		accountURL: accountURL,
	}
}

// Get implements customer.Interface#Get as an RPC
func (c *AccountClient) GetAccount(ctx context.Context, id string) (*internal.Account, error) {
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
		return nil, errors.New(strings.TrimSpace((string(body))))
	}

	var a internal.Account
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&a)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
