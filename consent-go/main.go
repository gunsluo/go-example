package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"text/template"

	"github.com/justinas/nosurf"
	"github.com/ory/hydra-client-go/client"
	"github.com/ory/hydra-client-go/client/admin"
	"github.com/ory/hydra-client-go/models"
)

const (
	envHydraAdminUrl = "HYDRA_ADMIN_URL"
)

var (
	adminClient *client.OryHydra
)

func main() {
	t, err := template.New("").Parse(loginHtml)
	if err != nil {
		panic(err)
	}
	loginTemplate = t

	t1, err := template.New("").Parse(consentHtml)
	if err != nil {
		panic(err)
	}
	consentTemplate = t1

	t2, err := template.New("").Parse(logoutHtml)
	if err != nil {
		panic(err)
	}
	logoutTemplate = t2

	adminUrl := os.Getenv(envHydraAdminUrl)
	if adminUrl == "" {
		adminUrl = "http://127.0.0.1:4445"
	}

	adminURL, err := url.Parse(adminUrl)
	if err != nil {
		panic(err)
	}
	hydraAdminClient := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{Schemes: []string{adminURL.Scheme}, Host: adminURL.Host, BasePath: adminURL.Path})
	adminClient = hydraAdminClient

	mux := http.NewServeMux()
	mux.Handle("/login", nosurf.New(http.HandlerFunc(loginHandler)))
	mux.Handle("/consent", nosurf.New(http.HandlerFunc(consentHandler)))
	mux.Handle("/logout", nosurf.New(http.HandlerFunc(logoutHandler)))

	fmt.Println("Now server is running on port 3000.")
	http.ListenAndServe(":3000", mux)
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
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
			rejectRequest := admin.NewRejectLoginRequestParams().
				WithLoginChallenge(challenge).
				WithContext(req.Context()).
				WithBody(&models.RejectRequest{
					Error:            "access_denied",
					ErrorDescription: "The resource owner denied the request",
					StatusCode:       http.StatusForbidden,
				})
			rejectReply, err := adminClient.Admin.RejectLoginRequest(rejectRequest)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if rejectReply == nil || rejectReply.Payload == nil ||
				rejectReply.Payload.RedirectTo == nil {
				http.Error(w, "invalid response from reject", http.StatusInternalServerError)
				return
			}

			redirectUrl := *rejectReply.Payload.RedirectTo
			http.Redirect(w, req, redirectUrl, http.StatusFound)
			return
		}

		if email != "foo@bar.com" && password != "foobar" {
			http.Error(w, "The username / password combination is not correct", http.StatusUnauthorized)
			return
		}

		loginRequest := admin.NewGetLoginRequestParams().
			WithLoginChallenge(challenge).
			WithContext(req.Context())
		loginReply, err := adminClient.Admin.GetLoginRequest(loginRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if loginReply == nil || loginReply.Payload == nil {
			http.Error(w, "invalid response from get login", http.StatusInternalServerError)
			return
		}

		if loginReply.Payload.RequestURL != nil {
			fmt.Printf("-------->login request url, %v\n", *loginReply.Payload.RequestURL)
		}

		var remember bool
		if rememberFrom == "1" {
			remember = true
		}
		acceptRequest := admin.NewAcceptLoginRequestParams().
			WithLoginChallenge(challenge).
			WithContext(req.Context()).
			WithBody(&models.AcceptLoginRequest{
				Subject:     &email,
				Remember:    remember,
				RememberFor: 3600,
				Acr:         oidcConformityMaybeFakeAcr(loginReply, "0"),
			})
		acceptReply, err := adminClient.Admin.AcceptLoginRequest(acceptRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if acceptReply == nil || acceptReply.Payload == nil || acceptReply.Payload.RedirectTo == nil {
			http.Error(w, "invalid response from accept", http.StatusInternalServerError)
			return
		}

		// redirect
		redirectUrl := *acceptReply.Payload.RedirectTo
		http.Redirect(w, req, redirectUrl, http.StatusFound)

		return
	}

	csrfToken := nosurf.Token(req)
	challenge := req.URL.Query().Get("login_challenge")

	loginRequest := admin.NewGetLoginRequestParams().
		WithLoginChallenge(challenge).
		WithContext(req.Context())
	loginReply, err := adminClient.Admin.GetLoginRequest(loginRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if loginReply == nil || loginReply.Payload == nil {
		http.Error(w, "invalid response from get login", http.StatusInternalServerError)
		return
	}

	if loginReply.Payload.Skip != nil && *loginReply.Payload.Skip {
		// accept login
		acceptRequest := admin.NewAcceptLoginRequestParams().
			WithLoginChallenge(challenge).
			WithContext(req.Context()).
			WithBody(&models.AcceptLoginRequest{
				Subject: loginReply.Payload.Subject,
			})
		acceptReply, err := adminClient.Admin.AcceptLoginRequest(acceptRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if acceptReply == nil || acceptReply.Payload == nil || acceptReply.Payload.RedirectTo == nil {
			http.Error(w, "invalid response from accept", http.StatusInternalServerError)
			return
		}

		// redirect
		redirectUrl := *acceptReply.Payload.RedirectTo
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

	if err := loginTemplate.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func consentHandler(w http.ResponseWriter, req *http.Request) {
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
			rejectRequest := admin.NewRejectConsentRequestParams().
				WithConsentChallenge(challenge).
				WithContext(req.Context()).
				WithBody(&models.RejectRequest{
					Error:            "access_denied",
					ErrorDescription: "The resource owner denied the request",
					StatusCode:       http.StatusForbidden,
				})
			rejectReply, err := adminClient.Admin.RejectConsentRequest(rejectRequest)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if rejectReply == nil || rejectReply.Payload == nil ||
				rejectReply.Payload.RedirectTo == nil {
				http.Error(w, "invalid response from reject", http.StatusInternalServerError)
				return
			}

			redirectUrl := *rejectReply.Payload.RedirectTo
			http.Redirect(w, req, redirectUrl, http.StatusFound)
			return
		}

		consentRequest := admin.NewGetConsentRequestParams().
			WithConsentChallenge(challenge).
			WithContext(req.Context())
		consentReply, err := adminClient.Admin.GetConsentRequest(consentRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if consentReply == nil || consentReply.Payload == nil {
			http.Error(w, "invalid response from get consent", http.StatusInternalServerError)
			return
		}

		var remember bool
		if rememberFrom == "1" {
			remember = true
		}
		acceptRequest := admin.NewAcceptConsentRequestParams().
			WithConsentChallenge(challenge).
			WithContext(req.Context()).
			WithBody(&models.AcceptConsentRequest{
				GrantScope:               grantScopes,
				Remember:                 remember,
				RememberFor:              3600,
				Session:                  oidcConformityMaybeFakeSession(consentReply, grantScopes),
				GrantAccessTokenAudience: consentReply.Payload.RequestedAccessTokenAudience,
			})
		acceptReply, err := adminClient.Admin.AcceptConsentRequest(acceptRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if acceptReply == nil || acceptReply.Payload == nil || acceptReply.Payload.RedirectTo == nil {
			http.Error(w, "invalid response from accept", http.StatusInternalServerError)
			return
		}

		// redirect
		redirectUrl := *acceptReply.Payload.RedirectTo
		http.Redirect(w, req, redirectUrl, http.StatusFound)

		return
	}

	csrfToken := nosurf.Token(req)
	challenge := req.URL.Query().Get("consent_challenge")

	consentRequest := admin.NewGetConsentRequestParams().
		WithConsentChallenge(challenge).
		WithContext(req.Context())
	consentReply, err := adminClient.Admin.GetConsentRequest(consentRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if consentReply == nil || consentReply.Payload == nil {
		http.Error(w, "invalid response from get consent", http.StatusInternalServerError)
		return
	}

	if consentReply.Payload.Skip {
		grantScopes := consentReply.Payload.RequestedScope
		acceptRequest := admin.NewAcceptConsentRequestParams().
			WithConsentChallenge(challenge).
			WithContext(req.Context()).
			WithBody(&models.AcceptConsentRequest{
				GrantScope:               grantScopes,
				Remember:                 false,
				RememberFor:              3600,
				Session:                  oidcConformityMaybeFakeSession(consentReply, grantScopes),
				GrantAccessTokenAudience: consentReply.Payload.RequestedAccessTokenAudience,
			})
		acceptReply, err := adminClient.Admin.AcceptConsentRequest(acceptRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if acceptReply == nil || acceptReply.Payload == nil || acceptReply.Payload.RedirectTo == nil {
			http.Error(w, "invalid response from accept", http.StatusInternalServerError)
			return
		}

		// redirect
		redirectUrl := *acceptReply.Payload.RedirectTo
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

	if err := consentTemplate.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func logoutHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		// Parse the form
		if err := req.ParseForm(); err != nil {
			http.Error(w, "Could not parse form", http.StatusBadRequest)
			return
		}

		challenge := req.Form.Get("challenge")
		action := req.Form.Get("submit")

		if action == "No" {
			rejectRequest := admin.NewRejectLogoutRequestParams().
				WithLogoutChallenge(challenge).
				WithContext(req.Context())
			_, err := adminClient.Admin.RejectLogoutRequest(rejectRequest)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			return
		}

		acceptRequest := admin.NewAcceptLogoutRequestParams().
			WithLogoutChallenge(challenge).
			WithContext(req.Context())
		acceptReply, err := adminClient.Admin.AcceptLogoutRequest(acceptRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if acceptReply == nil || acceptReply.Payload == nil || acceptReply.Payload.RedirectTo == nil {
			http.Error(w, "invalid response from accept", http.StatusInternalServerError)
			return
		}

		// redirect
		redirectUrl := *acceptReply.Payload.RedirectTo
		http.Redirect(w, req, redirectUrl, http.StatusFound)
		return
	}

	csrfToken := nosurf.Token(req)
	challenge := req.URL.Query().Get("logout_challenge")

	logoutRequest := admin.NewGetLogoutRequestParams().
		WithLogoutChallenge(challenge).
		WithContext(req.Context())
	logoutReply, err := adminClient.Admin.GetLogoutRequest(logoutRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if logoutReply == nil || logoutReply.Payload == nil {
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

	if err := logoutTemplate.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func oidcConformityMaybeFakeAcr(reply *admin.GetLoginRequestOK, defaultValue string) string {
	if reply == nil || reply.Payload == nil || reply.Payload.OidcContext == nil {
		return defaultValue
	}

	if len(reply.Payload.OidcContext.AcrValues) == 0 {
		return defaultValue
	}

	return reply.Payload.OidcContext.AcrValues[len(reply.Payload.OidcContext.AcrValues)-1]
}

func oidcConformityMaybeFakeSession(reply *admin.GetConsentRequestOK, grantScopes []string) *models.ConsentRequestSession {
	session := &models.ConsentRequestSession{}

	idToken := make(map[string]interface{})

	// fake email
	idToken["email"] = "foo@bar.com"
	idToken["email_verified"] = true

	// fake phone
	idToken["phone_number"] = "133713371337"
	idToken["phone_number_verified"] = false

	// fake profile
	idToken["name"] = "gunsluo"
	idToken["given_name"] = "gun"
	idToken["family_name"] = "luo"
	idToken["gender"] = "male"
	idToken["picture"] = "xxxx"
	idToken["updated_at"] = 1604416603

	//session.AccessToken = "mock access token"
	session.IDToken = idToken
	return session
}

var (
	loginTemplate   *template.Template
	consentTemplate *template.Template
	logoutTemplate  *template.Template
)

var loginHtml = `
<!DOCTYPE html>
<html>

<head>
  <title></title>
</head>

<body>
  <h1 id="login-title">Please log in</h1>
  <form action="login" method="POST">
  		<input type="hidden" name="csrf_token" value="{{ .CsrfToken }}">
		<input type="hidden" name="challenge" value="{{ .Challenge }}">
    <table>
      <tr>
        <td><input type="email" id="email" name="email" value="" placeholder="email@foobar.com"></td>
        <td>(it's "foo@bar.com")</td>
      </tr>
      <tr>
        <td><input type="password" id="password" name="password"></td>
        <td>(it's "foobar")</td>
      </tr>
    </table><input type="checkbox" id="remember" name="remember" value="1"><label for="remember">Remember
      me</label><br>
	  <input type="submit" id="accept" name="submit" value="Log in">
	  <input type="submit" id="reject" name="submit" value="Deny access">
  </form>
</body>

</html>
`

var consentHtml = `
<!DOCTYPE html>
<html>

<head>
  <title></title>
</head>

<body>
  <h1>An application requests access to your data!</h1>
  <form action="consent" method="POST">
    <input type="hidden" name="challenge" value="{{ .Challenge }}">
    <input type="hidden" name="csrf_token" value="{{ .CsrfToken }}">
    <p>Hi foo@bar.com, application <strong>auth-code-client</strong> wants access resources on your behalf and to:
    </p>
    <input class="grant_scope" type="checkbox" id="openid" value="openid" name="grant_scope">
    <label for="openid">openid</label><br>
    <input class="grant_scope" type="checkbox" id="offline" value="offline" name="grant_scope">
    <label for="offline">offline</label><br>
    <p>Do you want to be asked next time when this application wants to access your data? The application will
      not be able to ask for more permissions without your consent.</p>
    <ul></ul>
    <p><input type="checkbox" id="remember" name="remember" value="1"><label for="remember">Do not ask me again</label>
    </p>
    <p><input type="submit" id="accept" name="submit" value="Allow access"><input type="submit" id="reject"
        name="submit" value="Deny access"></p>
  </form>
</body>

</html>
`

var logoutHtml = `
<!DOCTYPE html>
<html>

<head>
  <title></title>
</head>

<body>
  <h1>Do you wish to log out?</h1>
  <form action="logout" method="POST">
    <input type="hidden" name="csrf_token" value="{{ .CsrfToken }}">
    <input type="hidden" name="challenge" value="{{ .Challenge }}">
    <input type="submit" id="accept" value="Yes"><input type="submit" id="reject" value="No">
  </form>
</body>

</html>`
