package main

import (
	"context"
	"fmt"
	"log"

	"github.com/graph-gophers/dataloader"
)

func main() {
	// setup batch function
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result
		// do some async work to get data for specified keys
		// append to this list resolved values
		for _, key := range keys {
			fmt.Printf("->%s %v\n", key.String(), key.Raw())
			if "key1" == key.String() {
				results = append(results, &dataloader.Result{Data: "100"})
			}
		}
		return results
	}

	// create Loader with an in-memory cache
	loader := dataloader.NewBatchedLoader(batchFn)

	/**
	 * Use loader
	 *
	 * A thunk is a function returned from a function that is a
	 * closure over a value (in this case an interface value and error).
	 * When called, it will block until the value is resolved.
	 */
	thunk := loader.Load(context.TODO(), dataloader.StringKey("key1")) // StringKey is a convenience method that make wraps string to implement `Key` interface
	result, err := thunk()
	if err != nil {
		// handle data error
	}
	log.Printf("value: %#v", result)

	thunk = loader.Load(context.TODO(), dataloader.StringKey("key1")) // StringKey is a convenience method that make wraps string to implement `Key` interface
	result, err = thunk()
	if err != nil {
		// handle data error
	}

	log.Printf("value: %#v", result)
}
