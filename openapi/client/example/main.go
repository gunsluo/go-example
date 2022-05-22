package main

import (
	"context"
	"fmt"

	"github.com/gunsluo/go-example/openapi/client/api"
)

func main() {
	cfg := api.NewConfiguration()
	cfg.Host = "127.0.0.1:9090"
	cfg.Scheme = "http"
	cfg.Servers = api.ServerConfigurations{
		{
			URL:         "/v2",
			Description: "version 2",
		},
	}

	client := api.NewAPIClient(cfg)

	ctx := context.WithValue(context.Background(), api.ContextAccessToken, "token")
	resp, _, err := client.OrganizationApi.AddMembersRequest(ctx).
		OrganizationId(100).OrganizatoinMember(
		[]api.OrganizatoinMember{
			{
				Name: "luoji",
				Mail: "luoji@ex.com",
			},
		}).Execute()
	if err != nil {
		panic(err)
	}

	if resp != nil {
		fmt.Println("->", resp.Data)
	}
}
