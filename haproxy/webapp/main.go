package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	room := values.Get("room")
	fmt.Println("room:", room)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("your room: " + room))
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	var address string
	flag.StringVar(&address, "a", ":54321", "listen address")
	flag.Parse()
	log.Printf("Starting on port %s", address)

	mux := http.NewServeMux()
	mux.HandleFunc("/api", handleHello)
	mux.HandleFunc("/health", handleHealth)
	log.Fatal(http.ListenAndServe(address, mux))
}
