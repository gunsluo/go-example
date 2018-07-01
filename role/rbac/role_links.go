package rbac

import (
	"errors"
	"sync"
)

type roleLinks struct {
	allRoles          *sync.Map
	maxHierarchyLevel int
}

// newRolLinks is the constructor for creating an instance of the
// default roleLinks implementation.
func newRoleLinks(maxHierarchyLevel int) *roleLinks {
	rl := roleLinks{}
	rl.allRoles = &sync.Map{}
	rl.maxHierarchyLevel = maxHierarchyLevel
	return &rl
}

func (rl *roleLinks) hasRole(name string) bool {
	_, ok := rl.allRoles.Load(name)
	return ok
}

func (rl *roleLinks) createRole(name string) *roleLayer {
	role, _ := rl.allRoles.LoadOrStore(name, newRole(name))
	return role.(*roleLayer)
}

// Clear clears all stored data and resets the role manager to the initial state.
func (rl *roleLinks) Clear() error {
	rl.allRoles = &sync.Map{}
	return nil
}

// AddLink adds the inheritance link between role: name1 and role: name2.
// aka role: name1 inherits role: name2.
// domain or ns is a prefix to the roles.
func (rl *roleLinks) AddLink(name1 string, name2 string, domain ...string) error {
	if len(domain) == 1 {
		name1 = domain[0] + "::" + name1
		name2 = domain[0] + "::" + name2
	} else if len(domain) > 1 {
		return errors.New("error: domain should be 1 parameter")
	}

	role1 := rl.createRole(name1)
	role2 := rl.createRole(name2)
	role1.addRole(role2)
	return nil
}

// DeleteLink deletes the inheritance link between role: name1 and role: name2.
// aka role: name1 does not inherit role: name2 any more.
// domain is a prefix to the roles.
func (rl *roleLinks) DeleteLink(name1 string, name2 string, domain ...string) error {
	if len(domain) == 1 {
		name1 = domain[0] + "::" + name1
		name2 = domain[0] + "::" + name2
	} else if len(domain) > 1 {
		return errors.New("error: domain should be 1 parameter")
	}

	if !rl.hasRole(name1) || !rl.hasRole(name2) {
		return errors.New("error: name1 or name2 does not exist")
	}

	role1 := rl.createRole(name1)
	role2 := rl.createRole(name2)
	role1.deleteRole(role2)
	return nil
}

// HasLink determines whether role: name1 inherits role: name2.
// domain is a prefix to the roles.
func (rl *roleLinks) HasLink(name1 string, name2 string, domain ...string) (bool, error) {
	if len(domain) == 1 {
		name1 = domain[0] + "::" + name1
		name2 = domain[0] + "::" + name2
	} else if len(domain) > 1 {
		return false, errors.New("error: domain should be 1 parameter")
	}

	if name1 == name2 {
		return true, nil
	}

	if !rl.hasRole(name1) || !rl.hasRole(name2) {
		return false, nil
	}

	role1 := rl.createRole(name1)
	return role1.hasRole(name2, rl.maxHierarchyLevel), nil
}

// GetRoles gets the roles that a subject inherits.
// domain is a prefix to the roles.
func (rl *roleLinks) GetRoles(name string, domain ...string) ([]string, error) {
	if len(domain) == 1 {
		name = domain[0] + "::" + name
	} else if len(domain) > 1 {
		return nil, errors.New("error: domain should be 1 parameter")
	}

	if !rl.hasRole(name) {
		return nil, errors.New("error: name does not exist")
	}

	roles := rl.createRole(name).getRoles()
	if len(domain) == 1 {
		for i := range roles {
			roles[i] = roles[i][len(domain[0])+2:]
		}
	}
	return roles, nil
}

// GetDeepRoles gets the roles that a subject inherits.
// domain is a prefix to the roles.
func (rl *roleLinks) GetDeepRoles(name string, domain ...string) ([]string, error) {
	if len(domain) == 1 {
		name = domain[0] + "::" + name
	} else if len(domain) > 1 {
		return nil, errors.New("error: domain should be 1 parameter")
	}

	if !rl.hasRole(name) {
		return nil, errors.New("error: name does not exist")
	}

	roles := rl.createRole(name).getDeepRoles(rl.maxHierarchyLevel)
	ArrayRemoveDuplicates(&roles)
	if len(domain) == 1 {
		for i := range roles {
			roles[i] = roles[i][len(domain[0])+2:]
		}
	}

	return roles, nil
}

// GetUsers gets the users that inherits a subject.
// domain is an unreferenced parameter here, may be used in other implementations.
func (rl *roleLinks) GetUsers(name string, domain ...string) ([]string, error) {
	if !rl.hasRole(name) {
		return nil, errors.New("error: name does not exist")
	}

	names := []string{}
	rl.allRoles.Range(func(_, value interface{}) bool {
		role := value.(*roleLayer)
		if role.hasDirectRole(name) {
			names = append(names, role.name)
		}

		return true
	})

	ArrayRemoveDuplicates(&names)
	return names, nil
}

// GetDeepUsers gets the users that inherits a subject.
// domain is an unreferenced parameter here, may be used in other implementations.
func (rl *roleLinks) GetDeepUsers(name string, domain ...string) ([]string, error) {
	if !rl.hasRole(name) {
		return nil, errors.New("error: name does not exist")
	}

	names := []string{}
	rl.allRoles.Range(func(_, value interface{}) bool {
		role := value.(*roleLayer)
		if role.hasRole(name, rl.maxHierarchyLevel) {
			if role.name != name {
				names = append(names, role.name)
			}
		}

		return true
	})

	ArrayRemoveDuplicates(&names)
	return names, nil
}

// PrintRoles prints all the roles to log.
func (rl *roleLinks) PrintRoles() error {
	rl.allRoles.Range(func(_, value interface{}) bool {
		LogPrint(value.(*roleLayer).toString())
		return true
	})
	return nil
}

// roleLayer represents the data structure for a role in RBAC.
type roleLayer struct {
	name  string
	roles []*roleLayer
}

func newRole(name string) *roleLayer {
	r := roleLayer{}
	r.name = name
	return &r
}

func (r *roleLayer) addRole(role *roleLayer) {
	for _, rr := range r.roles {
		if rr.name == role.name {
			return
		}
	}

	r.roles = append(r.roles, role)
}

func (r *roleLayer) deleteRole(role *roleLayer) {
	for i, rr := range r.roles {
		if rr.name == role.name {
			r.roles = append(r.roles[:i], r.roles[i+1:]...)
			return
		}
	}
}

func (r *roleLayer) hasRole(name string, hierarchyLevel int) bool {
	if r.name == name {
		return true
	}

	if hierarchyLevel <= 0 {
		return false
	}

	for _, role := range r.roles {
		if role.hasRole(name, hierarchyLevel-1) {
			return true
		}
	}
	return false
}

func (r *roleLayer) hasDirectRole(name string) bool {
	for _, role := range r.roles {
		if role.name == name {
			return true
		}
	}

	return false
}

func (r *roleLayer) toString() string {
	names := ""
	for i, role := range r.roles {
		if i == 0 {
			names += role.name
		} else {
			names += ", " + role.name
		}
	}
	return r.name + " < " + names
}

func (r *roleLayer) getRoles() []string {
	names := []string{}
	for _, role := range r.roles {
		names = append(names, role.name)
	}
	return names
}

func (r *roleLayer) getDeepRoles(hierarchyLevel int) []string {
	names := []string{}
	if hierarchyLevel <= 0 {
		return names
	}

	names = append(names, r.getRoles()...)
	for _, role := range r.roles {
		ret := role.getDeepRoles(hierarchyLevel - 1)
		names = append(names, ret...)
	}

	return names
}
