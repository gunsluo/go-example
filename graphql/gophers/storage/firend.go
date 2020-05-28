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

	loader, err := c.Loaders.Key("getUserById")
	if err != nil {
		return nil, err
	}

	//getUserById(id)
	id := r.node.firendId
	thunk := loader.Load(ctx, GetUserByIdKey{Id: id})
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

	var firends []*Firend
	for _, f := range mockUserFirends {
		if f.userId == id {
			firends = append(firends, f)
		}
	}

	return firends
}

func getFirendsByUserIds(ids []string) []*Firend {
	CallTimes++
	var firends []*Firend
	for _, f := range mockUserFirends {
		for _, id := range ids {
			if f.userId == id {
				firends = append(firends, f)
				break
			}
		}
	}

	return firends
}

var mockUserFirends = []*Firend{
	&Firend{userId: "1", firendId: "2"},
	&Firend{userId: "1", firendId: "3"},
	&Firend{userId: "2", firendId: "1"},
	&Firend{userId: "2", firendId: "4"},
	&Firend{userId: "3", firendId: "1"},
	&Firend{userId: "4", firendId: "1"},
	&Firend{userId: "4", firendId: "5"},
	&Firend{userId: "5", firendId: "4"},
}
