package srv

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
	client "github.com/ory/hydra-client-go"
)

func (s *Server) login(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		// Parse the form
		if err := req.ParseForm(); err != nil {
			http.Error(w, "Could not parse form", http.StatusBadRequest)
			return
		}

		//csrfToken := req.Form.Get("csrf_token")
		challenge := req.Form.Get("challenge")
		rememberFrom := req.Form.Get("remember")
		email := req.Form.Get("email")
		password := req.Form.Get("password")
		action := req.Form.Get("submit")

		if action != "Log in" {
			// reject
			rejectRequest := client.NewRejectRequest()
			rejectRequest.SetError("access_denied")
			rejectRequest.SetErrorDescription("The resource owner denied the request")
			rejectRequest.SetStatusCode(http.StatusForbidden)
			completedReqResp, _, err := s.apiClient.AdminApi.RejectLoginRequest(req.Context()).
				LoginChallenge(challenge).
				RejectRequest(*rejectRequest).
				Execute()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if completedReqResp == nil || completedReqResp.RedirectTo == "" {
				http.Error(w, "invalid response from reject", http.StatusInternalServerError)
				return
			}

			redirectUrl := completedReqResp.RedirectTo
			http.Redirect(w, req, redirectUrl, http.StatusFound)
			return
		}

		if email != "foo@bar.com" && password != "foobar" {
			http.Error(w, "The username / password combination is not correct", http.StatusUnauthorized)
			return
		}

		loginResp, _, err := s.apiClient.AdminApi.
			GetLoginRequest(req.Context()).
			LoginChallenge(challenge).Execute()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if loginResp == nil {
			http.Error(w, "invalid response from get login", http.StatusInternalServerError)
			return
		}

		if loginResp.RequestUrl != "" {
			fmt.Printf("-------->login request url, %v\n", loginResp.RequestUrl)
		}

		var remember bool
		if rememberFrom == "1" {
			remember = true
		}
		acceptRequest := client.NewAcceptLoginRequest(email)
		acceptRequest.SetRemember(remember)
		acceptRequest.SetRememberFor(3600)
		acceptRequest.SetAcr(oidcConformityMaybeFakeAcr(loginResp, "0"))
		completedReqResp, _, err := s.apiClient.AdminApi.AcceptLoginRequest(req.Context()).
			LoginChallenge(challenge).
			AcceptLoginRequest(*acceptRequest).
			Execute()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if completedReqResp == nil || completedReqResp.RedirectTo == "" {
			http.Error(w, "invalid response from accept", http.StatusInternalServerError)
			return
		}

		// redirect
		redirectUrl := completedReqResp.RedirectTo
		http.Redirect(w, req, redirectUrl, http.StatusFound)

		return
	}

	csrfToken := nosurf.Token(req)
	challenge := req.URL.Query().Get("login_challenge")

	loginResp, _, err := s.apiClient.AdminApi.
		GetLoginRequest(req.Context()).
		LoginChallenge(challenge).Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if loginResp == nil {
		http.Error(w, "invalid response from get login", http.StatusInternalServerError)
		return
	}

	if loginResp.Skip {
		// accept login
		completedReqResp, _, err := s.apiClient.AdminApi.AcceptLoginRequest(req.Context()).
			LoginChallenge(challenge).
			AcceptLoginRequest(*client.NewAcceptLoginRequest(loginResp.Subject)).
			Execute()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if completedReqResp == nil || completedReqResp.RedirectTo == "" {
			http.Error(w, "invalid response from accept", http.StatusInternalServerError)
			return
		}

		// redirect
		redirectUrl := completedReqResp.RedirectTo
		fmt.Println("2---->", redirectUrl)
		http.Redirect(w, req, redirectUrl, http.StatusFound)

		return
	}

	var data = struct {
		CsrfToken string
		Challenge string
	}{
		csrfToken,
		challenge,
	}

	if err := s.loginTemplate.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
