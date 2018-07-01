package rbac

import (
	"testing"
)

func TestRuleMode(t *testing.T) {
	model, err := newModel()
	if err != nil {
		t.Fatalf("rule mode: %s", err)
	}

	if !model.AddRule("g", "g", []string{"user 1", "member", "S"}) {
		t.Fatal("failed: add rule")
	}

	if !model.AddRule("g", "g", []string{"user 2", "member", "S"}) {
		t.Fatal("failed: add rule")
	}

	if !model.AddRule("g", "g", []string{"user 3", "admin", "S"}) {
		t.Fatal("failed: add rule")
	}

	if !model.AddRule("g", "g", []string{"admin", "member", "R"}) {
		t.Fatal("failed: add rule")
	}

	if !model.HasRule("g", "g", []string{"user 3", "admin", "S"}) {
		t.Fatal("failed: rule not exist")
	}

	if !model.HasRule("g", "g", []string{"admin", "member", "R"}) {
		t.Fatal("failed: rule not exist")
	}

	roles := model.GetRule("g", "g")
	if len(roles) != 4 {
		t.Fatal("failed: get rule")
	}

	roles = model.GetFilteredRule("g", "g", 2, "S")
	if len(roles) != 3 {
		t.Fatal("failed: get filtered rule")
	}

	roles = model.GetFilteredRule("g", "g", 1, "member", "S")
	if len(roles) != 2 {
		t.Fatal("failed: get filtered rule")
	}

	roles = model.GetFilteredRule("g", "g", 0, "admin", "member", "R")
	if len(roles) != 1 {
		t.Fatal("failed: get filtered rule")
	}

	values := model.GetValuesForFieldInRule("g", "g", 0)
	if len(values) != 4 {
		t.Fatal("failed: get filtered rule")
	}

	values = model.GetValuesForFieldInRule("g", "g", 1)
	if len(values) != 2 {
		t.Fatal("failed: get filtered rule")
	}

	values = model.GetValuesForFieldInRule("g", "g", 2)
	if len(values) != 2 {
		t.Fatal("failed: get filtered rule")
	}

	if !model.Remove("g", "g", []string{"admin", "member", "R"}) {
		t.Fatal("failed: remove rule")
	}

	if !model.RemoveFilteredRule("g", "g", 1, "member", "S") {
		t.Fatal("failed: remove rule")
	}

	if !model.RemoveFilteredRule("g", "g", 2, "S") {
		t.Fatal("failed: remove rule")
	}

	count := model.Count("g", "g")
	if count != 0 {
		t.Fatal("failed: count rule")
	}

	model.Clear()
}
