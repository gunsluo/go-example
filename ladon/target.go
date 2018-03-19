package main

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/ory/ladon"
	"github.com/ory/ladon/manager/memory"
)

// subject
// user: peter luoji zac ken
// role: admin reach cadre none

// actions
// all create delete update get list ...

// resource
// domain;namespace;orgin-resource
// namespace
// domain

// A bunch of exemplary policies
var pols = []ladon.Policy{
	&ladon.DefaultPolicy{
		ID: "r:reach",
		Description: `This policy allows reach user to create, delete and get the listed resources,
			but only if the client ip matches and the request states that they are the owner of those resources as well.`,
		Subjects:  []string{"reach", "luoji", "peter", "<zac|ken>"},
		Resources: []string{"target.com;reach;reach.domain.com:resource:123", "target.com;cadre;cadre.domain.com:resource:345", "other.com;other;something:foo:<.+>"},
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
		ID:          "u:luoji",
		Description: `This policy allows luoji to update any resources`,
		Subjects:    []string{"luoji"},
		Actions:     []string{"update"},
		Resources:   []string{"target.com;reach;<.*>"},
		Effect:      ladon.AllowAccess,
	},
}

func main() {
	/*
		// The database manager expects a sqlx.DB object
		db, err := sqlx.Open("postgres", "user=root password=root host=127.0.0.1 port=5432 dbname=ladon sslmode=disable") // Your driver and data source.
		//db, err := sqlx.Open("postgres", "user=root host=127.0.0.1 port=26257 dbname=ladon sslmode=disable") // Your driver and data source.
		if err != nil {
			log.Fatalf("Could not connect to database: %s", err)
		}

		// You must call SQLManager.CreateSchemas(schema, table) before use
		manager := sql.NewSQLManager(db, nil)
		n, err := manager.CreateSchemas("", "")
		if err != nil {
			log.Fatalf("Failed to create schemas: %s", err)
		}
		log.Printf("applied %d migrations", n)
	*/
	var err error
	manager := memory.NewMemoryManager()

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

	/*
		all, err := manager.GetAll(10, 0)
		fmt.Println("-->", len(all), err)
		for _, one := range all {
			fmt.Println("===>", one)
		}
	*/

	accessRequest := &ladon.Request{
		Subject:  "reach",
		Action:   "create",
		Resource: "target.com;reach;reach.domain.com:resource:123",
		Context: ladon.Context{
			"owner":    "reach",
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
