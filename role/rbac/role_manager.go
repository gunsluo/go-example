package rbac

import (
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/ory/pagination"
	"github.com/pkg/errors"
)

const (
	// user for role relation
	UserForRoleRelation = "U"
	// role for role relation
	RoleForRoleRelation = "R"
)

// RoleManager manage the user's role and role hierarchy.
type RoleManager struct {
	cache *roleCache
	store *roleStore

	rules   model
	rlinks  *roleLinks
	adapter *ruleAdapter

	db                 *sqlx.DB
	autoSave           bool
	autoBuildRoleLinks bool
}

// NewRoleManager returns a new role manager
func NewRoleManager(params ...interface{}) (*RoleManager, error) {
	rules, err := newModel()
	if err != nil {
		return nil, err
	}

	cache, err := newRoleCache()
	if err != nil {
		return nil, err
	}

	m := &RoleManager{
		cache:              cache,
		store:              newRoleStore(),
		rules:              rules,
		rlinks:             newRoleLinks(10),
		adapter:            newRuleAdapter(),
		autoSave:           true,
		autoBuildRoleLinks: true,
	}

	if len(params) == 1 {
		if db, ok := params[0].(*sqlx.DB); ok {
			m.db = db
		}
	}

	return m, nil
}

// Load reloads the role from database.
func (m *RoleManager) Load() error {
	if m.db == nil {
		m.buildRoleLinks()
		return nil
	}

	if err := m.store.load(m.db, m.cache.roles); err != nil {
		return err
	}

	if err := m.adapter.load(m.db, m.rules); err != nil {
		return err
	}

	if m.autoBuildRoleLinks {
		m.buildRoleLinks()
	}

	return nil
}

// buildRoleLinks manually rebuild the role inheritance ms.
func (m *RoleManager) buildRoleLinks() {
	m.rlinks.Clear()
	m.rules.buildRoleLinks(m.rlinks)
}

// AddRole adds a role inside a rule.
// Returns false if the user already has the role (aka not affected).
func (m *RoleManager) AddRole(role *Role) (bool, error) {
	exist, err := m.cache.Exist(role.Name)
	if err != nil || exist {
		return false, err
	}

	if m.db == nil {
		if err = m.cache.Insert(role); err != nil {
			return false, err
		}

		return true, nil
	}

	tx, err := m.db.Beginx()
	if err != nil {
		return false, err
	}

	if err := m.store.Insert(tx, role); err != nil {
		if rollErr := tx.Rollback(); rollErr != nil {
			return false, errors.Wrap(err, "rollback")
		}

		return false, err
	}

	if err = m.cache.Insert(role); err != nil {
		if rollErr := tx.Rollback(); rollErr != nil {
			return false, errors.Wrap(err, "rollback")
		}

		return false, err
	}

	if err = tx.Commit(); err != nil {
		if rollErr := tx.Rollback(); rollErr != nil {
			return false, errors.Wrap(err, "rollback")
		}

		return false, err
	}

	return true, nil
}

// AddRoles adds multiple roles inside the rule.
// Returns false if the user already has the role (aka not affected).
func (m *RoleManager) AddRoles(roles []*Role) (bool, error) {
	if len(roles) == 0 {
		return true, nil
	}

	if m.db == nil {
		for _, role := range roles {
			if err := m.cache.Insert(role); err != nil {
				return false, err
			}
		}

		return true, nil
	}

	tx, err := m.db.Beginx()
	if err != nil {
		return false, err
	}

	for _, role := range roles {
		exist, err := m.cache.Exist(role.Name)
		if err != nil || exist {
			if rollErr := tx.Rollback(); rollErr != nil {
				return false, errors.Wrap(err, "rollback")
			}

			return false, err
		}

		if err := m.store.Insert(tx, role); err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				return false, errors.Wrap(err, "rollback")
			}

			return false, err
		}

		if err = m.cache.Insert(role); err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				return false, errors.Wrap(err, "rollback")
			}

			return false, err
		}
	}

	if err = tx.Commit(); err != nil {
		if rollErr := tx.Rollback(); rollErr != nil {
			return false, errors.Wrap(err, "rollback")
		}

		return false, err
	}

	return true, nil
}

// HasRole determines whether a role inside the rule.
func (m *RoleManager) HasRole(roleName string) (bool, error) {
	exist, err := m.cache.Exist(roleName)
	if m.db == nil || err != nil || (exist && err == nil) {
		return exist, err
	}

	// from database
	return m.store.Exist(m.db, roleName)
}

// GetRole gets a role inside the rule.
func (m *RoleManager) GetRole(roleName string) (*Role, error) {
	role, err := m.cache.Get(roleName)
	if m.db == nil || err != nil || (role != nil && err == nil) {
		return role, err
	}

	// from database
	return m.store.Get(m.db, roleName)
}

// GetRoles gets the list of roles that show up in the current role.
func (m *RoleManager) GetRoles(limit, offset int64, likeRole ...string) ([]*Role, int64, error) {
	if m.db != nil {
		total, err := m.store.Count(m.db, likeRole...)
		if err != nil {
			return nil, 0, err
		}

		roles, err := m.store.GetAll(m.db, limit, offset, likeRole...)
		if err != nil {
			return nil, 0, err
		}

		return roles, total, nil
	}

	// getting from the cache
	roles, err := m.cache.GetAll()
	if err != nil {
		return nil, 0, err
	}

	total, err := m.cache.Count()
	if err != nil {
		return nil, 0, err
	}

	if total == 0 {
		return nil, 0, nil
	}

	s, e := pagination.Index(int(limit), int(offset), int(total))
	return roles[s:e], total, nil
}

// CountRoles gets the count of roles that show up in the current role.
func (m *RoleManager) CountRoles(likeRole ...string) (int64, error) {
	if m.db != nil {
		return m.store.Count(m.db, likeRole...)
	}

	return m.cache.Count()
}

// RemoveRole deletes the role and all its users.
func (m *RoleManager) RemoveRole(name string) error {
	if m.db == nil {
		rules := m.getRuleByName("g", "g", name)
		for _, rule := range rules {
			if err := m.removeRule(nil, "g", "g", rule.V0, rule.V1, rule.V2); err != nil {
				return err
			}
		}

		return m.cache.Delete(name)
	}

	tx, err := m.db.Beginx()
	if err != nil {
		return err
	}

	if err = m.store.Delete(tx, name); err != nil {
		if rollErr := tx.Rollback(); rollErr != nil {
			return errors.Wrap(err, "rollback")
		}

		return err
	}

	rules := m.getRuleByName("g", "g", name)
	for _, rule := range rules {
		if err := m.removeRule(m.db, "g", "g", rule.V0, rule.V1, rule.V2); err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				return errors.Wrap(err, "rollback")
			}
			return err
		}
	}

	if err = m.cache.Delete(name); err != nil {
		if rollErr := tx.Rollback(); rollErr != nil {
			return errors.Wrap(err, "rollback")
		}

		return err
	}

	if err = tx.Commit(); err != nil {
		if rollErr := tx.Rollback(); rollErr != nil {
			return errors.Wrap(err, "rollback")
		}

		return err
	}

	return nil
}

// AddRoleForUser adds a role to a user.
func (m *RoleManager) AddRoleForUser(role string, user string) error {
	if exist, err := m.HasRole(role); err != nil || !exist {
		return errors.Errorf("role:%s not exist", role)
	}

	if m.db == nil {
		return m.addRule(nil, "g", "g", user, role, UserForRoleRelation)
	}

	return m.addRule(m.db, "g", "g", user, role, UserForRoleRelation)
}

// AddRoleForUsers adds the same role to multiple users
func (m *RoleManager) AddRoleForUsers(role string, users []string) error {
	if len(users) == 0 {
		return nil
	}

	if exist, err := m.HasRole(role); err != nil || !exist {
		return errors.Errorf("role:%s not exist", role)
	}

	if m.db == nil {
		for _, user := range users {
			if err := m.addRule(nil, "g", "g", user, role, UserForRoleRelation); err != nil {
				return err
			}
		}

		return nil
	}

	tx, err := m.db.Beginx()
	if err != nil {
		return err
	}

	for _, user := range users {
		if err := m.addRule(tx, "g", "g", user, role, UserForRoleRelation); err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				return errors.Wrap(err, "rollback")
			}

			return err
		}
	}

	if err = tx.Commit(); err != nil {
		if rollErr := tx.Rollback(); rollErr != nil {
			return errors.Wrap(err, "rollback")
		}

		return err
	}

	return nil
}

// AddRolesForUser adds the multiple roles to a user
func (m *RoleManager) AddRolesForUser(user string, roles []string) error {
	if len(roles) == 0 {
		return nil
	}

	if m.db == nil {
		for _, role := range roles {
			if exist, err := m.HasRole(role); err != nil || !exist {
				return errors.Errorf("role:%s not exist", role)
			}

			if err := m.addRule(nil, "g", "g", user, role, UserForRoleRelation); err != nil {
				return err
			}
		}

		return nil
	}

	tx, err := m.db.Beginx()
	if err != nil {
		return err
	}

	for _, role := range roles {
		if exist, err := m.HasRole(role); err != nil || !exist {
			if rollErr := tx.Rollback(); rollErr != nil {
				return errors.Wrap(err, "rollback")
			}

			return errors.Errorf("role:%s not exist", role)
		}

		if err := m.addRule(tx, "g", "g", user, role, UserForRoleRelation); err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				return errors.Wrap(err, "rollback")
			}

			return err
		}
	}

	if err = tx.Commit(); err != nil {
		if rollErr := tx.Rollback(); rollErr != nil {
			return errors.Wrap(err, "rollback")
		}

		return err
	}

	return nil
}

// RemoveRoleForUser deletes a role for the user.
func (m *RoleManager) RemoveRoleForUser(role, user string) error {
	if exist, err := m.HasRole(role); err != nil || !exist {
		return errors.Errorf("role:%s not exist", role)
	}

	if m.db == nil {
		return m.removeRule(nil, "g", "g", user, role, UserForRoleRelation)
	}

	return m.removeRule(m.db, "g", "g", user, role, UserForRoleRelation)
}

// HasRoleForUser determines whether a user has a role.
func (m *RoleManager) HasRoleForUser(role, user string) (bool, error) {
	rules := m.getFiltered("g", "g", 0, user, role, UserForRoleRelation)
	if len(rules) > 0 {
		return true, nil
	}

	rules = m.getFiltered("g", "g", 0, user, role, RoleForRoleRelation)
	if len(rules) > 0 {
		return true, nil
	}

	return false, nil
}

// AssignRole assign a role to role group, the role includes all permissions for role group.
func (m *RoleManager) AssignRole(roleGroup string, role string) error {
	if exist, err := m.HasRole(roleGroup); err != nil || !exist {
		return errors.Errorf("role:%s not exist", roleGroup)
	}

	if exist, err := m.HasRole(role); err != nil || !exist {
		return errors.Errorf("role:%s not exist", role)
	}

	if m.db == nil {
		return m.addRule(nil, "g", "g", role, roleGroup, RoleForRoleRelation)
	}

	return m.addRule(m.db, "g", "g", role, roleGroup, RoleForRoleRelation)
}

// AssignRoles assign multiple roles to role group, the role includes all permissions for role group.
func (m *RoleManager) AssignRoles(roleGroup string, roles []string) error {
	if len(roles) == 0 {
		return nil
	}

	if exist, err := m.HasRole(roleGroup); err != nil || !exist {
		return errors.Errorf("role:%s not exist", roleGroup)
	}

	if m.db == nil {
		for _, role := range roles {
			if exist, err := m.HasRole(role); err != nil || !exist {
				return errors.Errorf("role:%s not exist", role)
			}

			if err := m.addRule(nil, "g", "g", role, roleGroup, RoleForRoleRelation); err != nil {
				return err
			}
		}

		return nil
	}

	tx, err := m.db.Beginx()
	if err != nil {
		return err
	}

	for _, role := range roles {
		if exist, err := m.HasRole(role); err != nil || !exist {
			if rollErr := tx.Rollback(); rollErr != nil {
				return errors.Wrap(err, "rollback")
			}

			return errors.Errorf("role:%s not exist", role)
		}

		if err := m.addRule(tx, "g", "g", role, roleGroup, RoleForRoleRelation); err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				return errors.Wrap(err, "rollback")
			}

			return err
		}
	}

	if err = tx.Commit(); err != nil {
		if rollErr := tx.Rollback(); rollErr != nil {
			return errors.Wrap(err, "rollback")
		}

		return err
	}

	return nil
}

// CancelAssignRole cancel assign a role to role group.
func (m *RoleManager) CancelAssignRole(roleGroup string, role string) error {
	if exist, err := m.HasRole(roleGroup); err != nil || !exist {
		return errors.Errorf("role:%s not exist", roleGroup)
	}

	if exist, err := m.HasRole(role); err != nil || !exist {
		return errors.Errorf("role:%s not exist", role)
	}

	if m.db == nil {
		return m.removeRule(nil, "g", "g", role, roleGroup, RoleForRoleRelation)
	}

	return m.removeRule(m.db, "g", "g", role, roleGroup, RoleForRoleRelation)
}

// CancelAssignRoles cancel assign multiple role to role group.
func (m *RoleManager) CancelAssignRoles(roleGroup string, roles []string) error {
	if len(roles) == 0 {
		return nil
	}

	if exist, err := m.HasRole(roleGroup); err != nil || !exist {
		return errors.Errorf("role:%s not exist", roleGroup)
	}

	if m.db == nil {
		for _, role := range roles {
			if exist, err := m.HasRole(role); err != nil || !exist {
				return errors.Errorf("role:%s not exist", role)
			}

			if err := m.removeRule(nil, "g", "g", role, roleGroup, RoleForRoleRelation); err != nil {
				return err
			}
		}

		return nil
	}

	tx, err := m.db.Beginx()
	if err != nil {
		return err
	}

	for _, role := range roles {
		if exist, err := m.HasRole(role); err != nil || !exist {
			if rollErr := tx.Rollback(); rollErr != nil {
				return errors.Wrap(err, "rollback")
			}

			return errors.Errorf("role:%s not exist", role)
		}

		if err := m.removeRule(tx, "g", "g", role, roleGroup, RoleForRoleRelation); err != nil {
			if rollErr := tx.Rollback(); rollErr != nil {
				return errors.Wrap(err, "rollback")
			}

			return err
		}
	}

	if err = tx.Commit(); err != nil {
		if rollErr := tx.Rollback(); rollErr != nil {
			return errors.Wrap(err, "rollback")
		}

		return err
	}

	return nil
}

// GetRolesForUser gets the roles that a user has.
func (m *RoleManager) GetRolesForUser(name string, likeRoles ...string) ([]*Role, error) {
	var likeRole string
	if len(likeRoles) > 0 {
		likeRole = likeRoles[0]
	}

	ast := m.rules["g"]["g"]
	if ast.RL == nil {
		return nil, nil
	}

	roleNames, err := ast.RL.GetRoles(name)
	if err != nil {
		return nil, err
	}

	var roles []*Role
	for _, roleName := range roleNames {
		if likeRole == "" || strings.HasPrefix(roleName, likeRole) {
			role, err := m.GetRole(roleName)
			if err != nil {
				return nil, err
			}

			roles = append(roles, role)
		}
	}

	return roles, nil
}

// GetRolesHierarchyForUser gets all roles of hierarchy that a user has.
func (m *RoleManager) GetRolesHierarchyForUser(name string, likeRoles ...string) ([]*Role, error) {
	var likeRole string
	if len(likeRoles) > 0 {
		likeRole = likeRoles[0]
	}

	ast := m.rules["g"]["g"]
	if ast.RL == nil {
		return nil, nil
	}

	roleNames, err := ast.RL.GetDeepRoles(name)
	if err != nil {
		return nil, err
	}

	var roles []*Role
	for _, roleName := range roleNames {
		if likeRole == "" || strings.HasPrefix(roleName, likeRole) {
			role, err := m.GetRole(roleName)
			if err != nil {
				return nil, err
			}

			roles = append(roles, role)
		}
	}

	return roles, nil
}

const (
	SubjectUserType UserType = "S"
	RoleUserType    UserType = "R"
)

type UserType string

// User the user have two user type includes subject or role
type User struct {
	Name string
	Type UserType
}

// GetUsersForRole gets the users that has a role inside the rule.
func (m *RoleManager) GetUsersForRole(role string) ([]string, error) {
	ast := m.rules["g"]["g"]
	if ast.RL == nil {
		return nil, nil
	}

	users, err := ast.RL.GetUsers(role)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetUsersForRoleHierarchy gets the users that has all role hierarchy inside the rule.
func (m *RoleManager) GetUsersForRoleHierarchy(role string) ([]*User, error) {
	subjects, err := m.GetUsersForRole(role)
	if err != nil {
		return nil, err
	}

	var users []*User
	for _, sub := range subjects {
		if exist, err := m.HasRole(sub); err == nil && !exist {
			users = append(users, &User{Name: sub, Type: SubjectUserType})
		}
	}

	roles, err := m.GetRolesForUser(role)
	if err != nil {
		return nil, err
	}

	for _, role := range roles {
		users = append(users, &User{Name: role.Name, Type: RoleUserType})
	}

	return users, nil
}

func (m *RoleManager) addRule(db XODB, sec string, ptype string, params ...string) error {
	if notExist := m.rules.Add(sec, ptype, params); !notExist {
		return errors.New("rule exist already")
	}

	if db != nil && m.adapter != nil && m.autoSave {
		if err := m.adapter.Insert(db, ptype, params); err != nil {
			return err
		}
	}

	if err := m.addLinks(sec, ptype, params...); err != nil {
		return err
	}

	return nil
}

func (m *RoleManager) addLinks(sec, ptype string, params ...string) error {
	if m.autoBuildRoleLinks {
		ast := m.rules["g"]["g"]
		count := strings.Count(ast.Value, "_")
		if count < 2 {
			return errors.New("the number of '_' in role definition should be at least 2")
		}
		if len(params) < count {
			return errors.New("grouping policy elements do not meet role definition")
		}

		if count == 2 {
			//ast.RL.AddLink(params[0], params[1])
			m.rlinks.AddLink(params[0], params[1])
		} else if count == 3 {
			m.rlinks.AddLink(params[0], params[1], params[2])
		} else if count == 4 {
			m.rlinks.AddLink(params[0], params[1], params[2], params[3])
		}
	}

	return nil
}

func (m *RoleManager) removeRule(db XODB, sec string, ptype string, params ...string) error {
	if exist := m.rules.Remove(sec, ptype, params); !exist {
		return nil
	}

	if db != nil && m.adapter != nil && m.autoSave {
		if err := m.adapter.Delete(db, ptype, params); err != nil {
			return err
		}
	}

	m.removeLinks(sec, ptype, params...)

	return nil
}

func (m *RoleManager) removeLinks(sec, ptype string, params ...string) error {
	if m.autoBuildRoleLinks {
		ast := m.rules["g"]["g"]
		count := strings.Count(ast.Value, "_")
		if count < 2 {
			return errors.New("the number of '_' in role definition should be at least 2")
		}
		if len(params) < count {
			return errors.New("grouping policy elements do not meet role definition")
		}

		if count == 2 {
			//ast.RL.DeleteLink(params[0], params[1])
			m.rlinks.DeleteLink(params[0], params[1])
		} else if count == 3 {
			m.rlinks.DeleteLink(params[0], params[1], params[2])
		} else if count == 4 {
			m.rlinks.DeleteLink(params[0], params[1], params[2], params[3])
		}
	}

	return nil
}

func (m *RoleManager) getFiltered(sec string, ptype string, fieldIndex int, fieldValues ...string) []*Rule {
	var rules []*Rule
	mRules := m.rules.GetFilteredRule(sec, ptype, fieldIndex, fieldValues...)
	for _, mRule := range mRules {
		if len(mRule) >= 3 {
			rule := &Rule{PType: ptype, V0: mRule[0], V1: mRule[1], V2: mRule[2]}
			rules = append(rules, rule)
		}
	}

	return rules
}

func (m *RoleManager) getRuleByName(sec string, ptype string, name string) []*Rule {
	var rules []*Rule

	for _, rule := range m.rules[sec][ptype].Rule {
		if len(rule) >= 3 {
			if rule[0] == name || rule[1] == name {
				rules = append(rules, &Rule{PType: ptype, V0: rule[0], V1: rule[1], V2: rule[2]})
			}
		}
	}

	return rules
}

// Close close database connection
func (m *RoleManager) Close() {
	if m.db != nil {
		m.db.Close()
		m.db = nil
	}
}
