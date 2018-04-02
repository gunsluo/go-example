package main

import (
	"fmt"

	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
)

var policys = [][]interface{}{
	[]interface{}{"anonymous", "/register", "read"},
	[]interface{}{"member", "/member/a", "write"},
	[]interface{}{"admin", "/a", "write"},
}

var roleUsers = map[string]string{
	"admin":  "member",
	"member": "v0",
	"luoji":  "admin",
	"tom":    "member",
	/*
		"member": "admin",
		"luoji":  "member",
		"tom":    "admin",
		"v0":     "member",
	*/
}

func main() {
	a := gormadapter.NewAdapter("postgres", "user=root password=root host=127.0.0.1 port=5432 sslmode=disable") // Your driver and data source.
	e := casbin.NewEnforcer("model.conf", a)
	for _, p := range policys {
		if ok := e.AddPolicy(p...); !ok {
			fmt.Printf("policy %s exist already.\n", p)
		}
	}

	for user, role := range roleUsers {
		if ok := e.AddRoleForUser(user, role); !ok {
			fmt.Printf("add role[%s] for user[%s] already.\n", role, user)
		}
	}
	e.SavePolicy()

	test(e, "luoji", "/a", "write")
	test(e, "luoji", "/member/a", "write")
	test(e, "tom", "/a", "write")
	test(e, "tom", "/member/a", "write")

	roles := e.GetRolesForUser("luoji")
	fmt.Println("-->", roles)
	roles = e.GetRolesForUser("tom")
	fmt.Println("-->", roles)

	users := e.GetUsersForRole("member")
	fmt.Println("-->", users)
	users = e.GetUsersForRole("admin")
	fmt.Println("-->", users)

	//test(e, "anonymous", "/register", "read")
}

func test(e *casbin.Enforcer, sub, obj, act string) {
	//sub := "luoji" // the user that wants to access a resource.
	//obj := "/abc"  // the resource that is going to be accessed.
	//act := "read"  // the operation that the user performs on the resource.

	ret := e.Enforce(sub, obj, act)
	fmt.Println("-->", ret)
	/*
		if DeepEnforce(e, sub, obj, act) == true {
			// permit alice to read data1
			fmt.Printf("%s %s %s ok.\n", sub, act, obj)
		} else {
			fmt.Printf("%s %s %s no.\n", sub, act, obj)
			// deny the request, show an error
		}
	*/
}

func DeepEnforce(e *casbin.Enforcer, sub, obj, act string) bool {
	if e.Enforce(sub, obj, act) == true {
		return true
	}

	roles := e.GetRolesForUser(sub)
	for _, role := range roles {
		if e.Enforce(role, obj, act) == true {
			return true
		}
	}

	users := e.GetUsersForRole(sub)
	for _, user := range users {
		if e.Enforce(user, obj, act) == true {
			return true
		}
	}

	return false
}
