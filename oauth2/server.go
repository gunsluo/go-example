package main

import (
	"fmt"
	"net/http"

	"github.com/RangelReale/osin"
	"github.com/gunsluo/go-example/oauth2/store"
)

func main() {
	// ex.NewTestStorage implements the "osin.Storage" interface
	server := osin.NewServer(osin.NewServerConfig(), store.NewTestStorage())

	// Authorization code endpoint
	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		resp := server.NewResponse()
		defer resp.Close()

		if ar := server.HandleAuthorizeRequest(resp, r); ar != nil {
			// HANDLE LOGIN PAGE HERE

			ar.Authorized = true
			server.FinishAuthorizeRequest(resp, r, ar)
		}
		osin.OutputJSON(resp, w, r)
	})

	// Access token endpoint
	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		resp := server.NewResponse()
		defer resp.Close()

		if ar := server.HandleAccessRequest(resp, r); ar != nil {
			ar.Authorized = true
			server.FinishAccessRequest(resp, r, ar)
		}
		osin.OutputJSON(resp, w, r)
	})

	fmt.Println("oauth2 server run http://localhost:14000")
	http.ListenAndServe(":14000", nil)
}
