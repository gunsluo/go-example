package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ory/ladon"
	"github.com/ory/ladon/manager/sql"
)

// A bunch of exemplary policies
var pols = []ladon.Policy{
	&ladon.DefaultPolicy{
		ID: "1",
		Description: `This policy allows max, peter, zac and ken to create, delete and get the listed resources,
			but only if the client ip matches and the request states that they are the owner of those resources as well.`,
		Subjects:  []string{"max", "peter", "<zac|ken>"},
		Resources: []string{"myrn:some.domain.com:resource:123", "myrn:some.domain.com:resource:345", "myrn:something:foo:<.+>"},
		Actions:   []string{"<create|delete>", "get"},
		Effect:    ladon.AllowAccess,
		Conditions: ladon.Conditions{
			"owner": &ladon.EqualsSubjectCondition{},
			"clientIP": &ladon.CIDRCondition{
				CIDR: "127.0.0.1/32",
			},
		},
	},
	&ladon.DefaultPolicy{
		ID:          "2",
		Description: "This policy allows max to update any resource",
		Subjects:    []string{"max"},
		Actions:     []string{"update"},
		Resources:   []string{"<.*>"},
		Effect:      ladon.AllowAccess,
	},
	&ladon.DefaultPolicy{
		ID:          "3",
		Description: "This policy denies max to broadcast any of the resources",
		Subjects:    []string{"max"},
		Actions:     []string{"broadcast"},
		Resources:   []string{"<.*>"},
		Effect:      ladon.DenyAccess,
	},
	&ladon.DefaultPolicy{
		ID: "4",
		Description: `This policy allows max, peter, zac and ken to create, delete and get the listed resources,
			but only if the client ip matches and the request states that they are the owner of those resources as well.`,
		Subjects:  []string{"luoji"},
		Resources: []string{"myrn:some.domain.com:resource:123", "myrn:some.domain.com:resource:345", "myrn:something:foo:<.+>"},
		Actions:   []string{"<create|delete>", "get"},
		Effect:    ladon.AllowAccess,
		Conditions: ladon.Conditions{
			"owner": &ladon.EqualsSubjectCondition{},
			"clientIP": &ladon.CIDRCondition{
				CIDR: "127.0.0.1/32",
			},
		},
	},
}

// Some test cases
var cases = []struct {
	description   string
	accessRequest *ladon.Request
	expectErr     bool
}{
	{
		description: "should fail because no policy is matching as field clientIP does not satisfy the CIDR condition of policy 1.",
		accessRequest: &ladon.Request{
			Subject:  "peter",
			Action:   "delete",
			Resource: "myrn:some.domain.com:resource:123",
			Context: ladon.Context{
				"owner":    "peter",
				"clientIP": "0.0.0.0",
			},
		},
		expectErr: true,
	},
	{
		description: "should fail because no policy is matching as the owner of the resource 123 is zac, not peter!",
		accessRequest: &ladon.Request{
			Subject:  "peter",
			Action:   "delete",
			Resource: "myrn:some.domain.com:resource:123",
			Context: ladon.Context{
				"owner":    "zac",
				"clientIP": "127.0.0.1",
			},
		},
		expectErr: true,
	},
	{
		description: "should pass because policy 1 is matching and has effect allow.",
		accessRequest: &ladon.Request{
			Subject:  "peter",
			Action:   "delete",
			Resource: "myrn:some.domain.com:resource:123",
			Context: ladon.Context{
				"owner":    "peter",
				"clientIP": "127.0.0.1",
			},
		},
		expectErr: false,
	},
	{
		description: "should pass because max is allowed to update all resources.",
		accessRequest: &ladon.Request{
			Subject:  "max",
			Action:   "update",
			Resource: "myrn:some.domain.com:resource:123",
		},
		expectErr: false,
	},
	{
		description: "should pass because max is allowed to update all resource, even if none is given.",
		accessRequest: &ladon.Request{
			Subject:  "max",
			Action:   "update",
			Resource: "",
		},
		expectErr: false,
	},
	{
		description: "should fail because max is not allowed to broadcast any resource.",
		accessRequest: &ladon.Request{
			Subject:  "max",
			Action:   "broadcast",
			Resource: "myrn:some.domain.com:resource:123",
		},
		expectErr: true,
	},
	{
		description: "should fail because max is not allowed to broadcast any resource, even empty ones!",
		accessRequest: &ladon.Request{
			Subject: "max",
			Action:  "broadcast",
		},
		expectErr: true,
	},
}

func main() {
	// The database manager expects a sqlx.DB object
	//
	// For MySQL, be sure to include parseTime=true in the connection string
	// You can find all of the supported MySQL connection string options for the
	// driver at: https://github.com/go-sql-driver/mysql
	//
	db, err := sqlx.Open("postgres", "user=root password=root host=127.0.0.1 port=5432 dbname=ladon sslmode=disable") // Your driver and data source.
	//db, err := sqlx.Open("postgres", "user=root host=127.0.0.1 port=26257 dbname=ladon sslmode=disable") // Your driver and data source.
	// Or, if using postgres:
	//  import _ "github.com/lib/pq"
	//
	//  db, err = sqlx.Open("postgres", "postgres://foo:bar@localhost/ladon")
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	// You must call SQLManager.CreateSchemas(schema, table) before use
	// to apply the necessary SQL migrations
	//
	// You can provide your own schema and table name or pass
	// empty strings to use the default
	manager := sql.NewSQLManager(db, nil)
	n, err := manager.CreateSchemas("", "")
	if err != nil {
		log.Fatalf("Failed to create schemas: %s", err)
	}
	log.Printf("applied %d migrations", n)

	// Instantiate ladon with the default in-memory store.
	warden := &ladon.Ladon{
		Manager: manager,
	}
	// Add the policies defined above to the memory manager.
	for _, pol := range pols {
		err := warden.Manager.Create(pol)
		if err != nil {
			panic(err)
		}
	}

	//one, err := manager.Get("1")
	//fmt.Println("-->", one, err)

	all, err := manager.GetAll(10, 0)
	fmt.Println("-->", len(all), err)
	for _, one := range all {
		fmt.Println("===>", one)
	}

	//	for k, c := range cases {
	//	err := warden.IsAllowed(c.accessRequest)
	//	if err != nil {
	//		fmt.Printf("case=%d-%s   :%s\n", k, c.description, err)
	//	} else {
	//		fmt.Printf("case=%d-%s   :success\n", k, c.description)
	//	}
	//}
	accessRequest := &ladon.Request{
		Subject:  "max",
		Action:   "update",
		Resource: "myrn:something:foo:100",
		Context: ladon.Context{
			"owner":    "max",
			"clientIP": "127.0.0.1",
		},
	}
	err = warden.IsAllowed(accessRequest)
	if err != nil {
		fmt.Printf("case=   :%s\n", err)
	} else {
		fmt.Printf("case=   :success\n")
	}
}
