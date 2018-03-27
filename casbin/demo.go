package main

import (
	"fmt"

	"github.com/casbin/casbin"
	fileadapter "github.com/casbin/casbin/persist/file-adapter"
	_ "github.com/lib/pq"
)

var policys = [][]interface{}{
	[]interface{}{"anonymous", "domain", "/register", "read"},
	[]interface{}{"manager", "domain", "/manage", "read"},
	[]interface{}{"admin", "domain", "/*", "(read)|(write)"},
}

var roleForUserInDomainTable = [][]string{
	[]string{"luoji", "admin", "domain"},
}

/*
const (
	modelTxt = `
[request_definition]
r = sub, dom, obj, act

[policy_definition]
p = sub, dom, obj, act

[role_definition]
g = _, _, _

[policy_effect]
e = some(where (p.eft == allow)) && !some(where (p.eft == deny))

[matchers]
m = g(r.sub, p.sub, r.dom) && r.dom == p.dom && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)
`
)
*/

const (
	modelTxt = `
[request_definition]
r = sub, dom, obj, act

[policy_definition]
p = sub, dom, obj, act

[role_definition]
g = _, _, _, _

[policy_effect]
e = some(where (p.eft == allow)) && !some(where (p.eft == deny))

[matchers]
m = g(r.sub, p.sub, r.dom) && r.dom == p.dom && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)
`
)

func main() {
	m := casbin.NewModel(modelTxt)
	// Initialize a Xorm adapter and use it in a Casbin enforcer:
	// The adapter will use the Postgres database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	//password=postgres_password
	//a := xormadapter.NewAdapter("postgres", "user=root host=127.0.0.1 port=26257 sslmode=disable") // Your driver and data source.
	//a := gormadapter.NewAdapter("postgres", "user=root host=127.0.0.1 port=26257 sslmode=disable") // Your driver and data source.
	a := fileadapter.NewAdapter("policy.csv")

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := xormadapter.NewAdapter("postgres", "dbname=abc user=postgres_username password=postgres_password host=127.0.0.1 port=5432 sslmode=disable", true)
	e := casbin.NewEnforcer(m, a)

	e.AddGroupingPolicy("g", "Bruce Lee", "master", "china", "Kung Fu")
	e.AddGroupingPolicy("g", "YIP Man", "master", "china", "Kung Fu")
	e.AddGroupingPolicy("g", "Peter", "fan", "china", "Kung Fu")
	e.AddGroupingPolicy("g", "Tom", "fan", "japan", "Kung Fu")

	//e.LoadPolicy()

	roles := e.GetRolesForUserInDomain("Bruce Lee", "Kung Fu")
	roles := e.GetRolesForUserInNsAndDomain("Bruce Lee", "Kung Fu")
	fmt.Println("===>", roles)
	/*
		roles = e.GetAllRoles()
		fmt.Println("===>", roles)

		users := e.GetUsersForRole("master")
		fmt.Println("===>", users)
		e.GetAllSubjects()
	*/

	/*
		for _, p := range policys {
			if ok := e.AddPolicy(p...); !ok {
				fmt.Printf("policy %s exist already.\n", p)
			}
		}

		e.AddGroupingPolicy("luoji", "admin", "domain")
		for _, item := range roleForUserInDomainTable {
			ok := e.AddRoleForUserInDomain(item[0], item[1], item[2])
			fmt.Println("AddRoleForUserInDomain:", ok)
		}

		e.SavePolicy()
		// Load the policy from DB.
		//e.LoadPolicy()

		// Check the permission.
		//	e.Enforce("alice", "data1", "read")

		sub := "luoji" // the user that wants to access a resource.
		dom := "domain"
		obj := "/abc"  // the resource that is going to be accessed.
		act := "write" // the operation that the user performs on the resource.
		ret := e.Enforce(sub, dom, obj, act)
		fmt.Println("result->", ret)

		// Modify the policy.
		// e.AddPolicy(...)
		// e.RemovePolicy(...)

		//ok := e.DeleteUser("luoji")
		//fmt.Println("==>", ok)
		//e.DeleteRole("admin")
		// Save the policy back to DB.
		//e.SavePolicy()
	*/
}
