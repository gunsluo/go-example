package srv

import (
	"net/http"

	"github.com/justinas/nosurf"
	client "github.com/ory/hydra-client-go"
)

func (s *Server) consent(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		// Parse the form
		if err := req.ParseForm(); err != nil {
			http.Error(w, "Could not parse form", http.StatusBadRequest)
			return
		}

		challenge := req.Form.Get("challenge")
		grantScopes := req.Form["grant_scope"]
		rememberFrom := req.Form.Get("remember")
		action := req.Form.Get("submit")

		if action != "Allow access" {
			// reject
			rejectRequest := client.NewRejectRequest()
			rejectRequest.SetError("access_denied")
			rejectRequest.SetErrorDescription("The resource owner denied the request")
			rejectRequest.SetStatusCode(http.StatusForbidden)
			completedReqResp, _, err := s.apiClient.AdminApi.RejectConsentRequest(req.Context()).
				ConsentChallenge(challenge).
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

		consentResp, _, err := s.apiClient.AdminApi.GetConsentRequest(req.Context()).
			ConsentChallenge(challenge).
			Execute()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if consentResp == nil {
			http.Error(w, "invalid response from get consent", http.StatusInternalServerError)
			return
		}

		var remember bool
		if rememberFrom == "1" {
			remember = true
		}

		acceptConsentRequest := client.NewAcceptConsentRequest()
		acceptConsentRequest.SetGrantScope(grantScopes)
		acceptConsentRequest.SetRemember(remember)
		acceptConsentRequest.SetRememberFor(3600)
		acceptConsentRequest.SetSession(oidcConformityMaybeFakeSession(consentResp, grantScopes))
		acceptConsentRequest.SetGrantAccessTokenAudience(consentResp.RequestedAccessTokenAudience)
		completedReqResp, _, err := s.apiClient.AdminApi.AcceptConsentRequest(req.Context()).
			ConsentChallenge(challenge).
			AcceptConsentRequest(*acceptConsentRequest).
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
	challenge := req.URL.Query().Get("consent_challenge")

	consentResp, _, err := s.apiClient.AdminApi.GetConsentRequest(req.Context()).
		ConsentChallenge(challenge).
		Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if consentResp == nil {
		http.Error(w, "invalid response from get consent", http.StatusInternalServerError)
		return
	}

	if consentResp.Skip != nil && *consentResp.Skip {
		grantScopes := consentResp.RequestedScope

		acceptConsentRequest := client.NewAcceptConsentRequest()
		acceptConsentRequest.SetGrantScope(grantScopes)
		acceptConsentRequest.SetRemember(true)
		acceptConsentRequest.SetRememberFor(3600)
		acceptConsentRequest.SetSession(oidcConformityMaybeFakeSession(consentResp, grantScopes))
		acceptConsentRequest.SetGrantAccessTokenAudience(consentResp.RequestedAccessTokenAudience)
		completedReqResp, _, err := s.apiClient.AdminApi.AcceptConsentRequest(req.Context()).
			ConsentChallenge(challenge).
			AcceptConsentRequest(*acceptConsentRequest).
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

	var data = struct {
		CsrfToken string
		Challenge string
	}{
		csrfToken,
		challenge,
	}

	if err := s.consentTemplate.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
