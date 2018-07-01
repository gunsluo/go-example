package rbac

import (
	"testing"
)

func testRole(t *testing.T, rl *roleLinks, name1 string, name2 string, res bool) {
	t.Helper()
	myRes, _ := rl.HasLink(name1, name2)
	//log.Printf("%s, %s: %t", name1, name2, myRes)

	if myRes != res {
		t.Errorf("%s < %s: %t, supposed to be %t", name1, name2, !res, res)
	}
}

func testDomainRole(t *testing.T, rl *roleLinks, name1 string, name2 string, domain string, res bool) {
	t.Helper()
	myRes, _ := rl.HasLink(name1, name2, domain)
	//log.Printf("%s :: %s, %s: %t", domain, name1, name2, myRes)

	if myRes != res {
		t.Errorf("%s :: %s < %s: %t, supposed to be %t", domain, name1, name2, !res, res)
	}
}

func testPrintRoles(t *testing.T, rl *roleLinks, name string, res []string) {
	t.Helper()
	myRes, _ := rl.GetRoles(name)
	//log.Printf("%s: %s", name, myRes)

	if !ArrayEquals(myRes, res) {
		t.Errorf("%s: %s, supposed to be %s", name, myRes, res)
	}
}

func testGetDeepRoles(t *testing.T, rl *roleLinks, name string, res []string) {
	t.Helper()
	myRes, _ := rl.GetDeepRoles(name)
	//log.Printf("%s: %s", name, myRes)

	if !ArrayEquals(myRes, res) {
		t.Errorf("%s: %s, supposed to be %s", name, myRes, res)
	}
}

func testUsers(t *testing.T, rl *roleLinks, name string, res []string) {
	t.Helper()
	myRes, _ := rl.GetUsers(name)
	//log.Printf("%s: %s", name, myRes)

	if !ArrayEquals(myRes, res) {
		t.Errorf("%s: %s, supposed to be %s", name, myRes, res)
	}
}

func testGetDeepUsers(t *testing.T, rl *roleLinks, name string, res []string) {
	t.Helper()
	myRes, _ := rl.GetDeepUsers(name)
	//log.Printf("%s: %s", name, myRes)

	if !ArrayEquals(myRes, res) {
		t.Errorf("%s: %s, supposed to be %s", name, myRes, res)
	}
}

func TestRole(t *testing.T) {
	rl := newRoleLinks(3)
	rl.AddLink("u1", "g1")
	rl.AddLink("u2", "g1")
	rl.AddLink("u3", "g2")
	rl.AddLink("u4", "g2")
	rl.AddLink("u4", "g3")
	rl.AddLink("g1", "g3")

	// Current role inheritance tree:
	//             g3    g2
	//            /  \  /  \
	//          g1    u4    u3
	//         /  \
	//       u1    u2

	testRole(t, rl, "u1", "g1", true)
	testRole(t, rl, "u1", "g2", false)
	testRole(t, rl, "u1", "g3", true)
	testRole(t, rl, "u2", "g1", true)
	testRole(t, rl, "u2", "g2", false)
	testRole(t, rl, "u2", "g3", true)
	testRole(t, rl, "u3", "g1", false)
	testRole(t, rl, "u3", "g2", true)
	testRole(t, rl, "u3", "g3", false)
	testRole(t, rl, "u4", "g1", false)
	testRole(t, rl, "u4", "g2", true)
	testRole(t, rl, "u4", "g3", true)

	testPrintRoles(t, rl, "u1", []string{"g1"})
	testPrintRoles(t, rl, "u2", []string{"g1"})
	testPrintRoles(t, rl, "u3", []string{"g2"})
	testPrintRoles(t, rl, "u4", []string{"g2", "g3"})
	testPrintRoles(t, rl, "g1", []string{"g3"})
	testPrintRoles(t, rl, "g2", []string{})
	testPrintRoles(t, rl, "g3", []string{})

	testGetDeepRoles(t, rl, "u1", []string{"g1", "g3"})
	testGetDeepRoles(t, rl, "u2", []string{"g1", "g3"})
	testGetDeepRoles(t, rl, "u3", []string{"g2"})
	testGetDeepRoles(t, rl, "u4", []string{"g2", "g3"})
	testGetDeepRoles(t, rl, "g1", []string{"g3"})
	testGetDeepRoles(t, rl, "g2", []string{})
	testGetDeepRoles(t, rl, "g3", []string{})

	testUsers(t, rl, "u1", []string{})
	testUsers(t, rl, "u2", []string{})
	testUsers(t, rl, "u3", []string{})
	testUsers(t, rl, "u4", []string{})
	testUsers(t, rl, "g1", []string{"u1", "u2"})
	testUsers(t, rl, "g2", []string{"u3", "u4"})
	testUsers(t, rl, "g3", []string{"g1", "u4"})

	testGetDeepUsers(t, rl, "u1", []string{})
	testGetDeepUsers(t, rl, "u2", []string{})
	testGetDeepUsers(t, rl, "u3", []string{})
	testGetDeepUsers(t, rl, "u4", []string{})
	testGetDeepUsers(t, rl, "g1", []string{"u1", "u2"})
	testGetDeepUsers(t, rl, "g2", []string{"u3", "u4"})
	testGetDeepUsers(t, rl, "g3", []string{"u1", "u2", "u4", "g1"})

	rl.DeleteLink("g1", "g3")
	rl.DeleteLink("u4", "g2")

	// Current role inheritance tree after deleting the links:
	//             g3    g2
	//               \     \
	//          g1    u4    u3
	//         /  \
	//       u1    u2

	testRole(t, rl, "u1", "g1", true)
	testRole(t, rl, "u1", "g2", false)
	testRole(t, rl, "u1", "g3", false)
	testRole(t, rl, "u2", "g1", true)
	testRole(t, rl, "u2", "g2", false)
	testRole(t, rl, "u2", "g3", false)
	testRole(t, rl, "u3", "g1", false)
	testRole(t, rl, "u3", "g2", true)
	testRole(t, rl, "u3", "g3", false)
	testRole(t, rl, "u4", "g1", false)
	testRole(t, rl, "u4", "g2", false)
	testRole(t, rl, "u4", "g3", true)

	testPrintRoles(t, rl, "u1", []string{"g1"})
	testPrintRoles(t, rl, "u2", []string{"g1"})
	testPrintRoles(t, rl, "u3", []string{"g2"})
	testPrintRoles(t, rl, "u4", []string{"g3"})
	testPrintRoles(t, rl, "g1", []string{})
	testPrintRoles(t, rl, "g2", []string{})
	testPrintRoles(t, rl, "g3", []string{})

	testGetDeepRoles(t, rl, "u1", []string{"g1"})
	testGetDeepRoles(t, rl, "u2", []string{"g1"})
	testGetDeepRoles(t, rl, "u3", []string{"g2"})
	testGetDeepRoles(t, rl, "u4", []string{"g3"})
	testGetDeepRoles(t, rl, "g1", []string{})
	testGetDeepRoles(t, rl, "g2", []string{})
	testGetDeepRoles(t, rl, "g3", []string{})
}

func TestDomainRole(t *testing.T) {
	rl := newRoleLinks(3)
	rl.AddLink("u1", "g1", "domain1")
	rl.AddLink("u2", "g1", "domain1")
	rl.AddLink("u3", "admin", "domain2")
	rl.AddLink("u4", "admin", "domain2")
	rl.AddLink("u4", "admin", "domain1")
	rl.AddLink("g1", "admin", "domain1")

	// Current role inheritance tree:
	//       domain1:admin    domain2:admin
	//            /       \  /       \
	//      domain1:g1     u4         u3
	//         /  \
	//       u1    u2

	testDomainRole(t, rl, "u1", "g1", "domain1", true)
	testDomainRole(t, rl, "u1", "g1", "domain2", false)
	testDomainRole(t, rl, "u1", "admin", "domain1", true)
	testDomainRole(t, rl, "u1", "admin", "domain2", false)

	testDomainRole(t, rl, "u2", "g1", "domain1", true)
	testDomainRole(t, rl, "u2", "g1", "domain2", false)
	testDomainRole(t, rl, "u2", "admin", "domain1", true)
	testDomainRole(t, rl, "u2", "admin", "domain2", false)

	testDomainRole(t, rl, "u3", "g1", "domain1", false)
	testDomainRole(t, rl, "u3", "g1", "domain2", false)
	testDomainRole(t, rl, "u3", "admin", "domain1", false)
	testDomainRole(t, rl, "u3", "admin", "domain2", true)

	testDomainRole(t, rl, "u4", "g1", "domain1", false)
	testDomainRole(t, rl, "u4", "g1", "domain2", false)
	testDomainRole(t, rl, "u4", "admin", "domain1", true)
	testDomainRole(t, rl, "u4", "admin", "domain2", true)

	rl.DeleteLink("g1", "admin", "domain1")
	rl.DeleteLink("u4", "admin", "domain2")

	// Current role inheritance tree after deleting the links:
	//       domain1:admin    domain2:admin
	//                    \          \
	//      domain1:g1     u4         u3
	//         /  \
	//       u1    u2

	testDomainRole(t, rl, "u1", "g1", "domain1", true)
	testDomainRole(t, rl, "u1", "g1", "domain2", false)
	testDomainRole(t, rl, "u1", "admin", "domain1", false)
	testDomainRole(t, rl, "u1", "admin", "domain2", false)

	testDomainRole(t, rl, "u2", "g1", "domain1", true)
	testDomainRole(t, rl, "u2", "g1", "domain2", false)
	testDomainRole(t, rl, "u2", "admin", "domain1", false)
	testDomainRole(t, rl, "u2", "admin", "domain2", false)

	testDomainRole(t, rl, "u3", "g1", "domain1", false)
	testDomainRole(t, rl, "u3", "g1", "domain2", false)
	testDomainRole(t, rl, "u3", "admin", "domain1", false)
	testDomainRole(t, rl, "u3", "admin", "domain2", true)

	testDomainRole(t, rl, "u4", "g1", "domain1", false)
	testDomainRole(t, rl, "u4", "g1", "domain2", false)
	testDomainRole(t, rl, "u4", "admin", "domain1", true)
	testDomainRole(t, rl, "u4", "admin", "domain2", false)
}

func TestClear(t *testing.T) {
	rl := newRoleLinks(3)
	rl.AddLink("u1", "g1")
	rl.AddLink("u2", "g1")
	rl.AddLink("u3", "g2")
	rl.AddLink("u4", "g2")
	rl.AddLink("u4", "g3")
	rl.AddLink("g1", "g3")

	// Current role inheritance tree:
	//             g3    g2
	//            /  \  /  \
	//          g1    u4    u3
	//         /  \
	//       u1    u2

	rl.Clear()

	// All data is cleared.
	// No role inheritance now.

	testRole(t, rl, "u1", "g1", false)
	testRole(t, rl, "u1", "g2", false)
	testRole(t, rl, "u1", "g3", false)
	testRole(t, rl, "u2", "g1", false)
	testRole(t, rl, "u2", "g2", false)
	testRole(t, rl, "u2", "g3", false)
	testRole(t, rl, "u3", "g1", false)
	testRole(t, rl, "u3", "g2", false)
	testRole(t, rl, "u3", "g3", false)
	testRole(t, rl, "u4", "g1", false)
	testRole(t, rl, "u4", "g2", false)
	testRole(t, rl, "u4", "g3", false)
}
