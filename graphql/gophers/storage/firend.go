package storage

import (
	"context"
	"errors"
	"fmt"
)

type Firend struct {
	userId   string
	firendId string
}

// FirendResolver holds information about the thing we are resolving.
type FirendResolver struct {
	node *Firend
	root *RootResolver
	//loader *dataloader.Loader // we'll get this from the request context
}

func (r *FirendResolver) FirendId() *string {
	return &r.node.firendId
}

func (r *FirendResolver) User(ctx context.Context) (*UserResolver, error) {
	c, found := ctx.Value(GraphqlContextKey).(*Context)
	if !found {
		return nil, errors.New("unable to find the custom context")
	}

	id := r.node.firendId
	thunk := c.Loader.Load(ctx, ValueKey{K: "getUserById", V: id})
	data, err := thunk()
	if err != nil {
		return nil, fmt.Errorf("getUserById: %v", err)
	}
	user, ok := data.(*User)
	if !ok {
		return nil, fmt.Errorf("User: loaded the wrong type of data: %#v", data)
	}

	//user := getUserById(r.node.firendId)
	return &UserResolver{node: user, root: r.root}, nil
}

func getFirendsByUserId(id string) []*Firend {
	CallTimes++
	return []*Firend{
		&Firend{userId: id, firendId: "1"},
		&Firend{userId: id, firendId: "2"},
	}
}
