package storage

import (
	"context"
	"errors"
	"fmt"
)

type Friend struct {
	userId   string
	friendId string
}

// FriendResolver holds information about the thing we are resolving.
type FriendResolver struct {
	node *Friend
	root *RootResolver
	//loader *dataloader.Loader // we'll get this from the request context
}

func (r *FriendResolver) FriendId() *string {
	return &r.node.friendId
}

func (r *FriendResolver) User(ctx context.Context) (*UserResolver, error) {
	c, found := ctx.Value(GraphqlContextKey).(*Context)
	if !found {
		return nil, errors.New("unable to find the custom context")
	}

	loader, err := c.Loaders.Key("getUserById")
	if err != nil {
		return nil, err
	}

	//getUserById(id)
	id := r.node.friendId
	thunk := loader.Load(ctx, GetUserByIdKey{Id: id})
	data, err := thunk()
	if err != nil {
		return nil, fmt.Errorf("getUserById: %v", err)
	}
	user, ok := data.(*User)
	if !ok {
		return nil, fmt.Errorf("User: loaded the wrong type of data: %#v", data)
	}

	//user := getUserById(r.node.friendId)
	return &UserResolver{node: user, root: r.root}, nil
}

func getFriendsByUserId(id string) []*Friend {
	CallTimes++

	var friends []*Friend
	for _, f := range mockUserFriends {
		if f.userId == id {
			friends = append(friends, f)
		}
	}

	return friends
}

func getFriendsByUserIds(ids []string) []*Friend {
	CallTimes++
	var friends []*Friend
	for _, f := range mockUserFriends {
		for _, id := range ids {
			if f.userId == id {
				friends = append(friends, f)
				break
			}
		}
	}

	return friends
}

var mockUserFriends = []*Friend{
	&Friend{userId: "1", friendId: "2"},
	&Friend{userId: "1", friendId: "3"},
	&Friend{userId: "2", friendId: "1"},
	&Friend{userId: "2", friendId: "4"},
	&Friend{userId: "3", friendId: "1"},
	&Friend{userId: "4", friendId: "1"},
	&Friend{userId: "4", friendId: "5"},
	&Friend{userId: "5", friendId: "4"},
}
