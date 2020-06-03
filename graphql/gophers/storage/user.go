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

func (r *UserResolver) Friends(ctx context.Context) ([]*FriendResolver, error) {
	c, found := ctx.Value(GraphqlContextKey).(*Context)
	if !found {
		return nil, errors.New("unable to find the custom context")
	}

	loader, err := c.Loaders.Key("getFriendsByUserId")
	if err != nil {
		return nil, err
	}

	thunk := loader.Load(ctx, GetFriendsByUserIdKey{UserId: r.node.id})
	data, err := thunk()
	if err != nil {
		return nil, fmt.Errorf("getFriendsByUserId: %v", err)
	}
	friends, ok := data.([]*Friend)
	if !ok {
		return nil, fmt.Errorf("Friends: loaded the wrong type of data: %#v", data)
	}

	//friends := getFriendsByUserId(r.node.id)
	var resolvers []*FriendResolver
	for _, friend := range friends {
		resolvers = append(resolvers, &FriendResolver{node: friend, root: r.root})
	}
	return resolvers, nil
}
