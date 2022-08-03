package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api", api)
	http.HandleFunc("/login", login)

	fmt.Println("Listening on http://127.0.0.1:8000/")
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func api(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorized\n")
		return
	}

	if cookie.Value != "xxx" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorized\n")
		return
	}

	fmt.Fprintf(w, "api\n")
}

func login(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:    "access_token",
		Value:   "xxx",
		Expires: time.Now().Add(10 * time.Minute),
		//HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "success\n")
}
