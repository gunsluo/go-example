package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/gunsluo/go-example/graphql/gophers/storage"
)

func main() {
	rs := storage.NewRootResolver(storage.WithLogger(200), storage.WithDB(10))
	schema, err := graphql.ParseSchema(storage.SchemaString, rs)
	if err != nil {
		panic(err)
	}

	withContext := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("set dataloader for each request")
			times := storage.CallTimes
			loader := storage.DefaultLoaderFactory.NewLoader()
			c := &storage.Context{Values: make(map[string]interface{}), Loader: loader}

			ctx := r.Context()
			ctx = context.WithValue(ctx, storage.GraphqlContextKey, c)
			h.ServeHTTP(w, r.WithContext(ctx))

			fmt.Println("call times:", storage.CallTimes-times)
		})
	}

	http.Handle("/graphql", withContext(&relay.Handler{Schema: schema}))

	debugPage := bytes.Replace(storage.GraphiQLPage, []byte("fetch('/'"), []byte("fetch('/graphql'"), -1)
	http.HandleFunc("/debug.html", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(debugPage)
	})
	log.Println("run graphql server, :12345")
	log.Fatal(http.ListenAndServe(":12345", nil))
}
