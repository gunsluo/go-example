package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

var (
	clientID     string
	clientSecret string
	port         int
	//scopes       []string
	oauth2Config *oauth2.Config
	verifier     *oidc.IDTokenVerifier
)

func main() {
	flag.StringVar(&clientID, "client-id", "", "Please provide a Client ID using -client-id flag.")
	flag.StringVar(&clientSecret, "client-secret", "", "Please provide a Client Secret using -client-secret flag.")
	//flag.StringVar(&domain, "domain", "http://127.0.0.1", "callback domain.")
	flag.IntVar(&port, "port", 20000, "Please provide a Port using -port flag.")

	flag.Parse()

	listen := fmt.Sprintf(":%d", port)
	endpoint := "https://accounts.google.com"
	scopes := []string{"profile", "email"}
	redirectURL := fmt.Sprintf("http://127.0.0.1:%d/callback", port)
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, endpoint)
	if err != nil {
		panic(err)
	}

	verifier = provider.Verifier(
		&oidc.Config{ClientID: clientID},
	)

	oauth2Config = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		Scopes:       scopes,
		RedirectURL:  redirectURL,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleIndex)
	mux.HandleFunc("/login", handleLogin)
	mux.HandleFunc("/callback", handleCallback)

	log.Printf("listening on %s", listen)

	http.ListenAndServe(listen, mux)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(loginHtml))
}

var loginHtml = `<html>
  <body>
  	<a href="/login">Login</a>
  </body>
</html>`

func handleLogin(w http.ResponseWriter, r *http.Request) {
	state := "mock"
	authCodeURL := oauth2Config.AuthCodeURL(state)

	http.Redirect(w, r, authCodeURL, http.StatusSeeOther)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	if errType := q.Get("error"); errType != "" {
		q.Get("error_description")
		http.Error(w, errType+": "+q.Get("error_description"), http.StatusBadRequest)
		return
	}

	code := q.Get("code")
	token, err := oauth2Config.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "failed to get token: "+err.Error(), http.StatusBadRequest)
		return
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "no id_token in response", http.StatusBadRequest)
		return
	}

	idToken, err := verifier.Verify(r.Context(), rawIDToken)
	if err != nil {
		http.Error(w, "invalid id_token: "+err.Error(), http.StatusBadRequest)
		return
	}

	/*
		var claims struct {
			Username      string `json:"name"`
			Email         string `json:"email"`
			EmailVerified bool   `json:"email_verified"`
			HostedDomain  string `json:"hd"`
		}
	*/
	var claims json.RawMessage
	if err := idToken.Claims(&claims); err != nil {
		http.Error(w, "failed to decode claims: "+err.Error(), http.StatusBadRequest)
		return
	}

	buff := new(bytes.Buffer)
	json.Indent(buff, []byte(claims), "", "  ")

	//w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(buff.Bytes())
}
