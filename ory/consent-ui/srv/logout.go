package srv

import (
	"fmt"
	"net/http"

	"github.com/gunsluo/nosurf"
)

func (s *Server) logout(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	if req.Method == http.MethodPost {
		// Parse the form
		if err := req.ParseForm(); err != nil {
			http.Error(w, "Could not parse form", http.StatusBadRequest)
			return
		}

		challenge := req.Form.Get("challenge")
		action := req.Form.Get("submit")

		if action == "No" {
			_, err := s.apiClient.OAuth2Api.RejectOAuth2LogoutRequest(ctx).
				LogoutChallenge(challenge).
				Execute()
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			return
		}

		logoutResp, _, err := s.apiClient.OAuth2Api.GetOAuth2LogoutRequest(ctx).
			LogoutChallenge(challenge).
			Execute()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if logoutResp == nil {
			http.Error(w, "invalid response from get logout", http.StatusInternalServerError)
			return
		}

		oauth2RedirectTo, _, err := s.apiClient.OAuth2Api.AcceptOAuth2LogoutRequest(ctx).
			LogoutChallenge(challenge).
			Execute()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if oauth2RedirectTo == nil || oauth2RedirectTo.RedirectTo == "" {
			http.Error(w, "invalid response from accept", http.StatusInternalServerError)
			return
		}

		clientId, err := getClientIdFromLogoutURL(*logoutResp.RequestUrl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("logout--->", clientId)
		if clientId != "" {
			_, err = s.apiClient.OAuth2Api.RevokeOAuth2ConsentSessions(ctx).
				Subject(*logoutResp.Subject).
				Client(clientId).
				Execute()
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		// redirect
		redirectUrl := oauth2RedirectTo.RedirectTo
		http.Redirect(w, req, redirectUrl, http.StatusFound)
		return
	}

	csrfToken := nosurf.Token(req)
	challenge := req.URL.Query().Get("logout_challenge")

	logoutResp, _, err := s.apiClient.OAuth2Api.GetOAuth2LogoutRequest(ctx).
		LogoutChallenge(challenge).
		Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if logoutResp == nil {
		http.Error(w, "invalid response from get logout", http.StatusInternalServerError)
		return
	}

	var data = struct {
		CsrfToken string
		Challenge string
	}{
		csrfToken,
		challenge,
	}

	if err := s.logoutTemplate.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
