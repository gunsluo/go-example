package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
)

func main() {
	var address string
	flag.StringVar(&address, "a", ":8080", "listen address")
	flag.Parse()

	fmt.Printf("Echo server listening on %s.\n", address)

	err := http.ListenAndServe(
		address,
		http.HandlerFunc(handler),
		/*
			h2c.NewHandler(
				http.HandlerFunc(handler),
				&http2.Server{},
			),
		*/
	)
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(fmt.Sprintf("%s %s\n", r.Method, r.URL))
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	buffer := &bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf("%s %s %s\n", r.Proto, r.Method, r.URL))
	buffer.WriteString("\n")
	buffer.WriteString(fmt.Sprintf("Host: %s\n", r.Host))

	for key, values := range r.Header {
		for _, value := range values {
			buffer.WriteString(fmt.Sprintf("%s: %s\n", key, value))

		}
	}

	w.Write(buffer.Bytes())
}
