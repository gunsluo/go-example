package main

import (
	"encoding/json"
	"fmt"

	"github.com/ory/ladon"
	"gitlab.com/target-digital-transformation/access-control/pkg/ladon/condition"
)

func main() {
	context := ladon.Context{
		"owner":    "peter",
		"clientIP": "127.0.0.1",
	}

	buf, err := json.Marshal(context)
	fmt.Println("-->", string(buf), err)

	var ctx ladon.Context
	err = json.Unmarshal(buf, &ctx)
	fmt.Println("-->", ctx, err)

	conditions := ladon.Conditions{
		"owner": &condition.HasSubjectOfRoleCondition{},
		"clientIP": &ladon.CIDRCondition{
			CIDR: "127.0.0.1/32",
		},
	}
}
