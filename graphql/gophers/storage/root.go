package storage

import (
	"context"
	"errors"
	"fmt"
)

// RootResolver is a graphql root resolver
type RootResolver struct {
	G options
}

func NewRootResolver(opts ...Option) *RootResolver {
	r := &RootResolver{G: defaultOptions()}
	for _, option := range opts {
		option.apply(&r.G)
	}

	return r
}

// User resolves the query `thing`.
func (r *RootResolver) User(ctx context.Context, args *UserArgs) (*UserResolver, error) {
	// Here we are extracting the loader that we've placed on the context at the
	// beginning of the request, and asserting the type of the value is `*dataloader.Loader`.

	c, found := ctx.Value(GraphqlContextKey).(*Context)
	if !found {
		return nil, errors.New("unable to find the custom context")
	}
	c.WithValue("customkey", "custom value")
	fmt.Println("root resolver:", r.G.Logger, r.G.DB, "customkey:", c.Value("customkey"))

	id := string(args.ID)
	loader, err := c.Loaders.Key("getUserById")
	if err != nil {
		return nil, err
	}

	thunk := loader.Load(ctx, GetUserByIdKey{Id: id})
	data, err := thunk()
	if err != nil {
		return nil, fmt.Errorf("getUserById: %v", err)
	}
	user, ok := data.(*User)
	if !ok {
		return nil, fmt.Errorf("User: loaded the wrong type of data: %#v", data)
	}
	//user := getUserById(id)

	return &UserResolver{node: user, root: r}, nil
}

// User resolves the query `thing`.
func (r *RootResolver) Users(ctx context.Context, args struct {
	Limit  int32
	Offset int32
}) ([]*UserResolver, error) {
	c, found := ctx.Value(GraphqlContextKey).(*Context)
	if !found {
		return nil, errors.New("unable to find the custom context")
	}

	if args.Limit == 0 {
		args.Limit = 50
	}

	loader, err := c.Loaders.Key("getUsers")
	if err != nil {
		return nil, err
	}
	thunk := loader.Load(ctx, GetUsersKey{
		Limit:  int(args.Limit),
		Offset: int(args.Offset)})
	data, err := thunk()
	if err != nil {
		return nil, fmt.Errorf("getUsers: %v", err)
	}
	users, ok := data.([]*User)
	if !ok {
		return nil, fmt.Errorf("Users: loaded the wrong type of data: %#v", data)
	}

	//users := getUsers()
	var resolvers []*UserResolver
	for _, user := range users {
		resolvers = append(resolvers, &UserResolver{node: user, root: r})
	}
	return resolvers, nil
}

var CallTimes int
var OriginTimes int

func getUserById(id string) *User {
	CallTimes++
	for _, u := range mockUsers {
		if u.id == id {
			return u
		}
	}
	return nil
}

func getUserByIds(ids []string) []*User {
	CallTimes++
	var users []*User
	for _, id := range ids {
		for _, u := range mockUsers {
			if u.id == id {
				users = append(users, u)
				break
			}
		}
	}

	return users
}

func getUsers(limit, offset int) []*User {
	CallTimes++
	return mockUsers
}

var mockUsers = []*User{
	&User{id: "1", name: "luoji", fullname: "luoji"},
	&User{id: "2", name: "jerry", fullname: "jerry"},
	&User{id: "3", name: "mary", fullname: "mary"},
	&User{id: "4", name: "lili", fullname: "lili"},
	&User{id: "5", name: "mark", fullname: "mark"},
}
