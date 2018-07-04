package rbac

import "github.com/pkg/errors"

type roleCache struct {
	// using model as cache model, and then can be replaced it
	roles model
}

func newRoleCache() (*roleCache, error) {
	model, err := newModel()
	if err != nil {
		return nil, err
	}

	return &roleCache{roles: model}, nil
}

// Create adds a role to the cache.
func (cache *roleCache) Create(role, desc string) error {
	// The role name comes at the end of the arguments and is easy to find role by name
	return cache.add("g", "g", desc, role)
}

func (cache *roleCache) BatchCreate(roles []*Role) error {
	for _, role := range roles {
		if err := cache.add("g", "g", role.Description, role.Name); err != nil {
			return err
		}
	}

	return nil
}

func (cache *roleCache) add(sec string, ptype string, params ...string) error {
	if cache.roles.Add(sec, ptype, params) {
		return nil
	}

	return errors.Errorf("failed to add %s %s %v", sec, ptype, params)
}

// Delete deletes a role in the cache.
func (cache *roleCache) Delete(role string) error {
	return cache.removeFiltered("g", "g", 1, role)
}

func (cache *roleCache) removeFiltered(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	cache.roles.RemoveFilteredRule("g", "g", fieldIndex, fieldValues...)

	// role is not exist and return nil
	return nil
}

// Exist get a role from the cache if it exist.
func (cache *roleCache) Exist(role string) (bool, error) {
	roles := cache.getFiltered("g", "g", 1, role)
	return len(roles) > 0, nil
}

// GetByName gets a role by role' name from the cache.
func (cache *roleCache) GetByName(name string) (*Role, error) {
	roles := cache.getFiltered("g", "g", 1, name)
	if len(roles) == 0 {
		return nil, errors.New("role not exist")
	}

	return roles[0], nil
}

func (cache *roleCache) getFiltered(sec string, ptype string, fieldIndex int, fieldValues ...string) []*Role {
	var roles []*Role
	rules := cache.roles.GetFilteredRule(sec, ptype, fieldIndex, fieldValues...)
	for _, rule := range rules {
		if len(rule) >= 2 {
			role := &Role{Name: rule[1], Description: rule[0]}
			roles = append(roles, role)
		}
	}

	return roles
}

// GetAll gets roles from the cache.
func (cache *roleCache) GetAll() ([]*Role, error) {
	var roles []*Role
	rules := cache.roles.GetRule("g", "g")
	for _, rule := range rules {
		if len(rule) >= 2 {
			role := &Role{Name: rule[1], Description: rule[0]}
			roles = append(roles, role)
		}
	}

	return roles, nil
}

// Count gets total member of role from the cache.
func (cache *roleCache) Count() (int64, error) {
	count := cache.roles.Count("g", "g")
	return int64(count), nil
}
