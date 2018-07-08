package rbac

import (
	"testing"
)

func TestRuleModel(t *testing.T) {
	model, err := newModel()
	if err != nil {
		t.Fatalf("rule mode: %s", err)
	}

	if !model.Add("g", "g", []string{"user 1", "member", "U"}) {
		t.Fatal("failed: add rule")
	}

	if !model.Add("g", "g", []string{"user 2", "member", "U"}) {
		t.Fatal("failed: add rule")
	}

	if !model.Add("g", "g", []string{"user 3", "admin", "U"}) {
		t.Fatal("failed: add rule")
	}

	if !model.Add("g", "g", []string{"admin", "member", "R"}) {
		t.Fatal("failed: add rule")
	}

	if !model.HasRule("g", "g", []string{"user 3", "admin", "U"}) {
		t.Fatal("failed: rule not exist")
	}

	if !model.HasRule("g", "g", []string{"admin", "member", "R"}) {
		t.Fatal("failed: rule not exist")
	}

	roles := model.GetRule("g", "g")
	if len(roles) != 4 {
		t.Fatal("failed: get rule")
	}

	roles = model.GetFilteredRule("g", "g", 2, "U")
	if len(roles) != 3 {
		t.Fatal("failed: get filtered rule")
	}

	roles = model.GetFilteredRule("g", "g", 1, "member", "U")
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

	if !model.RemoveFilteredRule("g", "g", 1, "member", "U") {
		t.Fatal("failed: remove rule")
	}

	if !model.RemoveFilteredRule("g", "g", 2, "U") {
		t.Fatal("failed: remove rule")
	}

	count := model.Count("g", "g")
	if count != 0 {
		t.Fatal("failed: count rule")
	}

	model.Clear()
}
