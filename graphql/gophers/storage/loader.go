package storage

import (
	"context"
	"errors"

	"github.com/graph-gophers/dataloader"
)

type LoaderFunc func(key dataloader.Key) *dataloader.Result

type LoaderFactory struct {
	loadFuncs map[string]LoaderFunc
}

var DefaultLoaderFactory = LoaderFactory{loadFuncs: defaultLoadFuncs}

func (f *LoaderFactory) AddLoaderFunc(key string, fn LoaderFunc) *LoaderFactory {
	f.loadFuncs[key] = fn
	return f
}

func (f *LoaderFactory) NewLoader() *dataloader.Loader {
	batchFunc := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		results := []*dataloader.Result{}
		for _, key := range keys {
			fn, ok := f.loadFuncs[key.String()]
			if ok {
				result := fn(key)
				results = append(results, result)
			}
		}

		return results
	}

	return dataloader.NewBatchedLoader(batchFunc, dataloader.WithCache(NewCache()))
}

// ValueKey implements the Key interface for a string
type ValueKey struct {
	K string
	V interface{}
}

// String is an identity method. Used to implement String interface
func (v ValueKey) String() string { return v.K }

// Raw is an identity method. Used to implement Key Raw
func (v ValueKey) Raw() interface{} { return v.V }

var defaultLoadFuncs = map[string]LoaderFunc{
	"getUserById": func(key dataloader.Key) *dataloader.Result {
		id, ok := key.Raw().(string)
		if !ok {
			return &dataloader.Result{Error: errors.New("invalid parameter")}
		}

		d := getUserById(id)
		return &dataloader.Result{Data: d}
	},
	"getUsers": func(key dataloader.Key) *dataloader.Result {
		d := getUsers()
		return &dataloader.Result{Data: d}
	},
	"getFirendsByUserId": func(key dataloader.Key) *dataloader.Result {
		id, ok := key.Raw().(string)
		if !ok {
			return &dataloader.Result{Error: errors.New("invalid parameter")}
		}

		d := getFirendsByUserId(id)
		return &dataloader.Result{Data: d}
	},
}
