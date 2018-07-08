package rbac

import (
	"testing"
)

func TestRoleCache(t *testing.T) {
	cache, err := newRoleCache()
	if err != nil {
		t.Fatalf("Failed to create cache: %s", err)
	}

	err = cache.Insert(&Role{Name: "qs:admin", Description: "test"})
	if err != nil {
		t.Fatalf("Failed to add role: %s", err)
	}

	err = cache.Insert(&Role{Name: "qs:member", Description: "test"})
	if err != nil {
		t.Fatalf("Failed to add role: %s", err)
	}

	err = cache.Insert(&Role{Name: "qs:anonymous", Description: "test"})
	if err != nil {
		t.Fatalf("Failed to add role: %s", err)
	}

	exist, err := cache.Exist("qs:admin")
	if err != nil || !exist {
		t.Fatalf("Failed to get role: %s", err)
	}

	role, err := cache.Get("qs:admin")
	if err != nil {
		t.Fatalf("Failed to get all role: %s", err)
	}
	if role == nil {
		t.Fatalf("Failed to get role")
	}

	roles, err := cache.GetAll()
	if err != nil {
		t.Fatalf("Failed to get all role: %s", err)
	}
	if len(roles) != 3 {
		t.Fatalf("Failed to get all role")
	}

	count, err := cache.Count()
	if err != nil {
		t.Fatalf("Failed to count role: %s", err)
	}
	if count != 3 {
		t.Fatalf("Failed to count role")
	}

	err = cache.Delete("qs:admin")
	if err != nil {
		t.Fatalf("Failed to remove role: %s", err)
	}

	err = cache.Delete("qs:member")
	if err != nil {
		t.Fatalf("Failed to remove role: %s", err)
	}

	err = cache.Delete("qs:anonymous")
	if err != nil {
		t.Fatalf("Failed to remove role: %s", err)
	}
}
