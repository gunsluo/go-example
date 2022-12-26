package srv

import (
	"fmt"
	"net/http"

	"github.com/gunsluo/nosurf"
	client "github.com/ory/hydra-client-go/v2"
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

		ctx := req.Context()
		if action != "Log in" {
			// reject
			rejectRequest := client.NewRejectOAuth2Request()
			rejectRequest.SetError("access_denied")
			rejectRequest.SetErrorDescription("The resource owner denied the request")
			rejectRequest.SetStatusCode(http.StatusForbidden)
			oauth2RedirectTo, _, err := s.apiClient.OAuth2Api.RejectOAuth2LoginRequest(ctx).
				LoginChallenge(challenge).
				RejectOAuth2Request(*rejectRequest).
				Execute()

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if oauth2RedirectTo == nil || oauth2RedirectTo.RedirectTo == "" {
				http.Error(w, "invalid response from reject", http.StatusInternalServerError)
				return
			}

			redirectUrl := oauth2RedirectTo.RedirectTo
			http.Redirect(w, req, redirectUrl, http.StatusFound)
			return
		}

		if email != "foo@bar.com" && password != "foobar" {
			http.Error(w, "The username / password combination is not correct", http.StatusUnauthorized)
			return
		}

		loginResp, _, err := s.apiClient.OAuth2Api.GetOAuth2LoginRequest(ctx).
			LoginChallenge(challenge).
			Execute()
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

		acceptRequest := client.NewAcceptOAuth2LoginRequest(email)
		acceptRequest.SetRemember(remember)
		acceptRequest.SetRememberFor(3600)
		acceptRequest.SetAcr(oidcConformityMaybeFakeAcr(loginResp, "0"))

		oauth2RedirectTo, _, err := s.apiClient.OAuth2Api.AcceptOAuth2LoginRequest(ctx).
			LoginChallenge(challenge).
			AcceptOAuth2LoginRequest(*acceptRequest).
			Execute()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if oauth2RedirectTo == nil || oauth2RedirectTo.RedirectTo == "" {
			http.Error(w, "invalid response from accept", http.StatusInternalServerError)
			return
		}

		// redirect
		redirectUrl := oauth2RedirectTo.RedirectTo
		http.Redirect(w, req, redirectUrl, http.StatusFound)

		return
	}

	ctx := req.Context()
	csrfToken := nosurf.Token(req)
	challenge := req.URL.Query().Get("login_challenge")
	loginResp, _, err := s.apiClient.OAuth2Api.
		GetOAuth2LoginRequest(ctx).
		LoginChallenge(challenge).
		Execute()

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
		oauth2RedirectTo, _, err := s.apiClient.OAuth2Api.AcceptOAuth2LoginRequest(ctx).
			LoginChallenge(challenge).
			AcceptOAuth2LoginRequest(*client.NewAcceptOAuth2LoginRequest(loginResp.Subject)).
			Execute()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if oauth2RedirectTo == nil || oauth2RedirectTo.RedirectTo == "" {
			http.Error(w, "invalid response from accept", http.StatusInternalServerError)
			return
		}

		// redirect
		redirectUrl := oauth2RedirectTo.RedirectTo
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
