package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/graph-gophers/dataloader"
)

type LoaderFactory struct {
	loaders map[string]Loader
}

func (f *LoaderFactory) AddLoader(loader Loader) *LoaderFactory {
	f.loaders[loader.Key()] = loader
	return f
}

func (f *LoaderFactory) Loader(key string) (Loader, error) {
	l, ok := f.loaders[key]
	if !ok {
		return nil, fmt.Errorf("not found key %s", key)
	}
	return l, nil
}

func (f *LoaderFactory) NewDataLoaders() DataLoaders {
	dls := DataLoaders{}
	cache := NewCache()
	for _, l := range f.loaders {
		dl := dataloader.NewBatchedLoader(l.Batch, dataloader.WithCache(cache))
		dls[l.Key()] = dl
	}

	return dls
}

type Loader interface {
	Batch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result
	Key() string
}

type DataLoaders map[string]*dataloader.Loader

func (dls DataLoaders) Key(key string) (*dataloader.Loader, error) {
	dl, ok := dls[key]
	if !ok {
		return nil, fmt.Errorf("not found key %s", key)
	}
	return dl, nil
}

var defaultLoaders = map[string]Loader{
	"getUserById":        &getUserByIdLoader{},
	"getUsers":           &getUsersLoader{},
	"getFirendsByUserId": &getFirendsByUserIdLoader{},
}

var DefaultLoaderFactory = LoaderFactory{loaders: defaultLoaders}

// GetUserByIdKey implements the Key interface for a int
type GetUserByIdKey struct {
	Id string
}

func (k GetUserByIdKey) String() string { return k.Id }

func (k GetUserByIdKey) Raw() interface{} { return k.Id }

// getUserById loader
type getUserByIdLoader struct {
}

func (l *getUserByIdLoader) Batch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	OriginTimes += len(keys)

	var ids []string
	for _, key := range keys {
		//id, ok := key.Raw().(string)
		k, ok := key.(GetUserByIdKey)
		if ok {
			ids = append(ids, k.Id)
		}
	}

	users := getUserByIds(ids)

	var results []*dataloader.Result
	// should be sort by keys
	for _, u := range users {
		results = append(results, &dataloader.Result{Data: u})
	}

	return results
}

func (l *getUserByIdLoader) Key() string {
	return "getUserById"
}

// GetUserByIdKey implements the Key interface for a int
type GetUsersKey struct {
	Limit  int
	Offset int
}

func (k GetUsersKey) String() string { return "" }

func (k GetUsersKey) Raw() interface{} { return []int{k.Limit, k.Offset} }

// getUsers loader
type getUsersLoader struct {
}

func (l *getUsersLoader) Batch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	OriginTimes += len(keys)

	var results []*dataloader.Result
	// not merge request
	for _, key := range keys {
		k, ok := key.(GetUsersKey)
		if !ok {
			results = append(results, &dataloader.Result{Error: errors.New("invalid key")})
		} else {
			users := getUsers(k.Limit, k.Offset)
			results = append(results, &dataloader.Result{Data: users})
		}
	}

	return results
}
func (l *getUsersLoader) Key() string {
	return "getUsers"
}

// GetUserByIdKey implements the Key interface for a int
type GetFirendsByUserIdKey struct {
	UserId string
}

func (k GetFirendsByUserIdKey) String() string { return "" }

func (k GetFirendsByUserIdKey) Raw() interface{} { return k.UserId }

// getFirendsByUserId loader
type getFirendsByUserIdLoader struct {
}

func (l *getFirendsByUserIdLoader) Batch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	OriginTimes += len(keys)

	var ids []string
	for _, key := range keys {
		//id, ok := key.Raw().(string)
		k, ok := key.(GetFirendsByUserIdKey)
		if ok {
			ids = append(ids, k.UserId)
		}
	}

	firends := getFirendsByUserIds(ids)

	// sort
	sortFirends := map[string][]*Firend{}
	for _, f := range firends {
		if v, ok := sortFirends[f.userId]; ok {
			sortFirends[f.userId] = append(v, f)
		} else {
			sortFirends[f.userId] = []*Firend{f}
		}
	}

	var results []*dataloader.Result
	for _, id := range ids {
		if v, ok := sortFirends[id]; ok {
			results = append(results, &dataloader.Result{Data: v})
		} else {
			results = append(results, &dataloader.Result{Data: []*Firend{}})
		}
	}

	return results
}

func (l *getFirendsByUserIdLoader) Key() string {
	return "getFirendsByUserId"
}
