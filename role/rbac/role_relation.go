package rbac

import (
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const (
	UserForRoleRelation = "U"
	RoleForRoleRelation = "R"
)

// inheritance relation tree of role
type roleRelation struct {
	rules   model
	rlinks  *roleLinks
	adapter *ruleAdapter

	autoSave           bool
	autoBuildRoleLinks bool
}

func newRoleRelation(db *sqlx.DB) *roleRelation {
	rules, err := newModel()
	if err != nil {
		return nil
	}

	return &roleRelation{
		rules:              rules,
		rlinks:             newRoleLinks(10),
		adapter:            newRuleAdapter(db),
		autoSave:           true,
		autoBuildRoleLinks: true,
	}
}

// buildRoleLinks manually rebuild the role inheritance relations.
func (relation *roleRelation) buildRoleLinks() {
	relation.rlinks.Clear()
	relation.rules.buildRoleLinks(relation.rlinks)
}

// AddRoleForUser adds a role for a user.
// Returns false if the user already has the role (aka not affected).
func (relation *roleRelation) AddRoleForUser(user, role string) error {
	return relation.addRule("g", "g", user, role, UserForRoleRelation)
}

func (relation *roleRelation) addRule(sec string, ptype string, params ...string) error {
	notExist := relation.rules.Add(sec, ptype, params)
	if !notExist {
		return errors.New("rule exist already")
	}

	if err := relation.addLinks(sec, ptype, params...); err != nil {
		return err
	}

	if relation.adapter != nil && relation.autoSave {
		if err := relation.adapter.Add(ptype, params); err != nil {
			return err
		}
	}

	return nil
}

func (relation *roleRelation) addLinks(sec, ptype string, params ...string) error {
	if relation.autoBuildRoleLinks {
		ast := relation.rules["g"]["g"]
		count := strings.Count(ast.Value, "_")
		if count < 2 {
			return errors.New("the number of '_' in role definition should be at least 2")
		}
		if len(params) < count {
			return errors.New("grouping policy elements do not meet role definition")
		}

		if count == 2 {
			//ast.RL.AddLink(params[0], params[1])
			relation.rlinks.AddLink(params[0], params[1])
		} else if count == 3 {
			relation.rlinks.AddLink(params[0], params[1], params[2])
		} else if count == 4 {
			relation.rlinks.AddLink(params[0], params[1], params[2], params[3])
		}
	}

	return nil
}
