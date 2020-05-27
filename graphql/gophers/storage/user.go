package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/graph-gophers/graphql-go"
)

type User struct {
	id       string
	name     string
	fullname string
}

// UserResolver holds information about the thing we are resolving.
type UserResolver struct {
	node *User
	root *RootResolver
	//loader *dataloader.Loader // we'll get this from the request context
}

// UserArgs holds the arguments you pass into the `thing` query.
type UserArgs struct {
	ID graphql.ID
}

// ID is the User's identifier.
// We already have it, so let's not load any data.
func (r *UserResolver) ID() *graphql.ID {
	id := graphql.ID(r.node.id)
	return &id
}

// Name is what we call this User.
// We don't know it yet, so we have to load it.
func (r *UserResolver) Name(ctx context.Context) (*string, error) {
	return &r.node.name, nil
}

func (r *UserResolver) Fullname(ctx context.Context) (*string, error) {
	return &r.node.fullname, nil
}

func (r *UserResolver) Firends(ctx context.Context) ([]*FirendResolver, error) {
	c, found := ctx.Value(GraphqlContextKey).(*Context)
	if !found {
		return nil, errors.New("unable to find the custom context")
	}

	thunk := c.Loader.Load(ctx, ValueKey{K: "getFirendsByUserId", V: r.node.id})
	data, err := thunk()
	if err != nil {
		return nil, fmt.Errorf("getFirendsByUserId: %v", err)
	}
	firends, ok := data.([]*Firend)
	if !ok {
		return nil, fmt.Errorf("Firends: loaded the wrong type of data: %#v", data)
	}

	//firends := getFirendsByUserId(r.node.id)
	var resolvers []*FirendResolver
	for _, firend := range firends {
		resolvers = append(resolvers, &FirendResolver{node: firend, root: r.root})
	}
	return resolvers, nil
}
