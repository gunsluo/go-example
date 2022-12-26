package srv

import (
	"net/http"

	"github.com/gunsluo/nosurf"
	client "github.com/ory/hydra-client-go/v2"
)

func (s *Server) consent(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
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
			rejectRequest := client.NewRejectOAuth2Request()
			rejectRequest.SetError("access_denied")
			rejectRequest.SetErrorDescription("The resource owner denied the request")
			rejectRequest.SetStatusCode(http.StatusForbidden)

			oauth2RedirectTo, _, err := s.apiClient.OAuth2Api.RejectOAuth2ConsentRequest(ctx).
				ConsentChallenge(challenge).
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

		consentResp, _, err := s.apiClient.OAuth2Api.GetOAuth2ConsentRequest(ctx).
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

		acceptConsentRequest := client.NewAcceptOAuth2ConsentRequest()
		acceptConsentRequest.SetGrantScope(grantScopes)
		acceptConsentRequest.SetRemember(remember)
		acceptConsentRequest.SetRememberFor(3600)
		acceptConsentRequest.SetSession(oidcConformityMaybeFakeSession(consentResp, grantScopes))
		acceptConsentRequest.SetGrantAccessTokenAudience(consentResp.RequestedAccessTokenAudience)

		oauth2RedirectTo, _, err := s.apiClient.OAuth2Api.AcceptOAuth2ConsentRequest(ctx).
			ConsentChallenge(challenge).
			AcceptOAuth2ConsentRequest(*acceptConsentRequest).
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

	csrfToken := nosurf.Token(req)
	challenge := req.URL.Query().Get("consent_challenge")

	consentResp, _, err := s.apiClient.OAuth2Api.GetOAuth2ConsentRequest(ctx).
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

		acceptConsentRequest := client.NewAcceptOAuth2ConsentRequest()
		acceptConsentRequest.SetGrantScope(grantScopes)
		acceptConsentRequest.SetRemember(true)
		acceptConsentRequest.SetRememberFor(3600)
		acceptConsentRequest.SetSession(oidcConformityMaybeFakeSession(consentResp, grantScopes))
		acceptConsentRequest.SetGrantAccessTokenAudience(consentResp.RequestedAccessTokenAudience)
		completedReqResp, _, err := s.apiClient.OAuth2Api.AcceptOAuth2ConsentRequest(ctx).
			ConsentChallenge(challenge).
			AcceptOAuth2ConsentRequest(*acceptConsentRequest).
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
