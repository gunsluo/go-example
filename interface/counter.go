package main

import (
	"fmt"
	"log"
	"net/http"
)

type Counter int

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	*ctr++
	fmt.Fprintf(w, "counter = %d", *ctr)
}

func main() {
	ctr := new(Counter)
	http.Handle("/counter", ctr)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
