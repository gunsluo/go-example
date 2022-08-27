package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	c, err := NewClient(
		Config{
			Endpoint: "https://api.localsms.com",
			Sender:   "OIA",
			ClientId: "a3a26d70-2b7d-4ff4-b840-c228e643593b",
			ApiKey:   "Shc9/xtBjwUpTyhA96vFCBnIuENMQeBS/kiF91P3JT0=",
			SkipTLS:  true,
		},
	)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	data, err := c.Send(ctx, []string{"96821213333"}, "test")
	if err != nil {
		fmt.Printf("failed to send: %v\n", err)
		return
	}

	for _, item := range data {
		if item.MessageErrorCode != 0 {
			fmt.Printf("Error Code %d, to %v, description: %s\n", item.MessageErrorCode, item.MobileNumber, item.MessageErrorDescription)
			//return nil,
		} else {
			fmt.Printf("success, to %v, mid: %s\n", item.MobileNumber, item.MessageId)
		}
	}
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

	MessageErrorCode        int    `josn:"MessageErrorCode"`
	MessageErrorDescription string `josn:"MessageErrorDescription"`
	Custom                  string `josn:"Custom"`
}

func (c *oiaOwnerClient) Send(ctx context.Context, to []string, message string) ([]oiaOwnerSendResponseData, error) {
	v := url.Values{}
	v.Add("SenderId", c.config.Sender)
	v.Add("Is_Unicode", "false")
	v.Add("Is_Flash", "false")
	v.Add("Message", message)
	v.Add("MobileNumbers", strings.Join(to, ","))
	v.Add("ApiKey", c.config.ApiKey)
	v.Add("ClientId", c.config.ClientId)
	fmt.Println("-->", v.Encode())

	apiUrl := fmt.Sprintf("%s/api/v2/SendSMS?%s", c.config.Endpoint, v.Encode())
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}

	// param := &oiaOwnerSendRequest{
	// 	SenderId:      c.config.Sender,
	// 	IsUnicode:     false,
	// 	IsFlash:       false,
	// 	SchedTime:     "",
	// 	GroupId:       "",
	// 	Message:       message,
	// 	MobileNumbers: strings.Join(to, ","),
	// 	ApiKey:        c.config.ApiKey,
	// 	ClientId:      c.config.ClientId,
	// }
	// bb, err := json.Marshal(param)
	// if err != nil {
	// 	return nil, err
	// }

	// apiUrl := fmt.Sprintf("%s/api/v2/SendSMS", c.config.Endpoint)
	// req, err := http.NewRequest("POST", apiUrl, bytes.NewReader(bb))
	// if err != nil {
	// 	return nil, err
	// }
	// req.Header.Set("Content-Type", "application/json")

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
		var errResp = struct {
			ErrorCode        int    `josn:"ErrorCode"`
			ErrorDescription string `josn:"ErrorDescription"`
			Data             string `josn:"Data"`
		}{}

		if err := json.Unmarshal(respBytes, &errResp); err != nil {
			return nil, fmt.Errorf("Code %d, body: %s, err: %w", httpResp.StatusCode, string(respBytes), err)
		}

		return nil, fmt.Errorf("Error code %d, description: %s", errResp.ErrorCode, errResp.ErrorDescription)
	}

	fmt.Println("---->", string(respBytes))
	resp := &oiaOwnerSendResponse{}
	if err := json.Unmarshal(respBytes, resp); err != nil {
		return nil, fmt.Errorf("Code %d, body: %s, err: %w", httpResp.StatusCode, string(respBytes), err)
	}

	if resp.ErrorCode != 0 {
		return nil, fmt.Errorf("Error code %d, description: %s", resp.ErrorCode, resp.ErrorDescription)
	}

	return resp.Data, nil
}
