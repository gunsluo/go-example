package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	c, err := NewClient(
		Config{
			Endpoint: "",
			ClientId: "",
			ApiKey:   "password",
			SkipTLS:  false,
		},
	)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	ids, err := c.Send(ctx, []string{}, "test")
	if err != nil {
		panic(err)
	}
	fmt.Println("--->", ids, err)
}

type Config struct {
	Endpoint string
	ClientId string
	ApiKey   string
	Sender   string

	SkipTLS bool
	CaPath  string
}

func NewClient(config Config) (*oiaOwnerClient, error) {
	httpClient := &http.Client{}

	transport := &http.Transport{
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	if config.CaPath != "" {
		caCert, err := ioutil.ReadFile(config.CaPath)
		if err != nil {
			return nil, err
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		transport.TLSClientConfig = &tls.Config{
			RootCAs: caCertPool,
		}
	} else if config.SkipTLS {
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	httpClient.Transport = transport

	return &oiaOwnerClient{
		httpClient: httpClient,
		config:     config,
	}, nil
}

type oiaOwnerClient struct {
	httpClient *http.Client
	config     Config
}

type oiaOwnerSendRequest struct {
	SenderId      string `json:"SenderId"`
	IsUnicode     bool   `json:"Is_Unicode"`
	IsFlash       bool   `json:"Is_Flash"`
	SchedTime     string `json:"SchedTime"`
	GroupId       string `json:"GroupId"`
	Message       string `json:"Message"`
	MobileNumbers string `json:"MobileNumbers"`
	ApiKey        string `json:"ApiKey"`
	ClientId      string `json:"ClientId"`
}

type oiaOwnerSendResponse struct {
	ErrorCode        int                        `josn:"ErrorCode"`
	ErrorDescription string                     `josn:"ErrorDescription"`
	Data             []oiaOwnerSendResponseData `josn:"Data"`
}

type oiaOwnerSendResponseData struct {
	MobileNumber string `josn:"MobileNumber"`
	MessageId    string `josn:"MessageId"`
}

func (c *oiaOwnerClient) Send(ctx context.Context, to []string, message string) ([]string, error) {
	param := &oiaOwnerSendRequest{
		SenderId:      c.config.Sender,
		IsUnicode:     true,
		IsFlash:       false,
		SchedTime:     "",
		GroupId:       strings.Join(to, ","),
		Message:       message,
		MobileNumbers: "",
		ApiKey:        c.config.ApiKey,
		ClientId:      c.config.ClientId,
	}
	bb, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	apiUrl := fmt.Sprintf("%s/api/v2/SendSMS", c.config.Endpoint)
	req, err := http.NewRequest("POST", apiUrl, bytes.NewReader(bb))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	httpResp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	respBytes, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status Code %d, body: %s", httpResp.StatusCode, string(respBytes))
	}

	resp := &oiaOwnerSendResponse{}
	if err := json.Unmarshal(respBytes, resp); err != nil {
		return nil, err
	}

	ids := []string{}
	for _, item := range resp.Data {
		ids = append(ids, item.MessageId)
	}

	return ids, nil
}
