package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.

	idTokenCookie := &http.Cookie{
		Name:    "id_token",
		Value:   "luoji",
		Domain:  "target.com",
		Expires: time.Now().Add(time.Hour),
	}
	http.SetCookie(w, idTokenCookie)

	//w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", HealthCheckHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
