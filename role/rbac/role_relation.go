package rbac

// inheritance relation tree of role
type roleRelation struct {
	rules model
	rl    *roleLinks
}
