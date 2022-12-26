package srv

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gunsluo/nosurf"
	client "github.com/ory/hydra-client-go/v2"
)

const (
	envHydraAdminUrl = "HYDRA_ADMIN_URL"
)

type Server struct {
	loginTemplate   *template.Template
	consentTemplate *template.Template
	logoutTemplate  *template.Template

	apiClient *client.APIClient
	mux       *http.ServeMux
}

func New() (*Server, error) {
	s := &Server{}

	if t, err := template.New("").Parse(loginHtml); err != nil {
		return nil, err
	} else {
		s.loginTemplate = t
	}

	if t, err := template.New("").Parse(consentHtml); err != nil {
		return nil, err
	} else {
		s.consentTemplate = t
	}

	if t, err := template.New("").Parse(logoutHtml); err != nil {
		return nil, err
	} else {
		s.logoutTemplate = t
	}

	adminUrl := os.Getenv(envHydraAdminUrl)
	if adminUrl == "" {
		adminUrl = "http://127.0.0.1:4445"
	}

	adminURL, err := url.Parse(adminUrl)
	if err != nil {
		return nil, err
	}

	cfg := client.NewConfiguration()
	cfg.Host = adminURL.Host
	cfg.Scheme = adminURL.Scheme

	s.apiClient = client.NewAPIClient(cfg)

	mux := http.NewServeMux()
	mux.Handle("/oauth2/login", nosurf.New(http.HandlerFunc(s.login)))
	mux.Handle("/oauth2/consent", nosurf.New(http.HandlerFunc(s.consent)))
	mux.Handle("/oauth2/logout", nosurf.New(http.HandlerFunc(s.logout)))
	s.mux = mux

	return s, nil
}

func (s *Server) Run() {
	fmt.Println("Now server is running on port 3000.")
	http.ListenAndServe(":3000", s.mux)
}

func oidcConformityMaybeFakeAcr(loginRequest *client.OAuth2LoginRequest, defaultValue string) string {
	if loginRequest == nil || loginRequest.OidcContext == nil {
		return defaultValue
	}

	if len(loginRequest.OidcContext.AcrValues) == 0 {
		return defaultValue
	}

	return loginRequest.OidcContext.AcrValues[len(loginRequest.OidcContext.AcrValues)-1]
}

func oidcConformityMaybeFakeSession(consentRequest *client.OAuth2ConsentRequest, grantScopes []string) client.AcceptOAuth2ConsentRequestSession {
	session := client.AcceptOAuth2ConsentRequestSession{}

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
	session.IdToken = map[string]map[string]interface{}{
		"user": idToken,
	}
	return session
}

func getClientIdFromLogoutURL(logoutURL string) (string, error) {
	u, err := url.Parse(logoutURL)
	if err != nil {
		return "", err
	}

	idTokenHint := u.Query().Get("id_token_hint")
	if idTokenHint == "" {
		return "", nil
	}

	claims := jwt.MapClaims{}
	_, _, err = new(jwt.Parser).ParseUnverified(idTokenHint, claims)
	if err != nil {
		return "", err
	}

	var clientId string
	if items, ok := claims["aud"].([]interface{}); !ok {
		return "", nil
	} else if len(items) == 0 {
		return "", nil
	} else {
		for _, item := range items {
			if s, ok := item.(string); ok {
				clientId = s
				break
			}
		}
	}

	return clientId, nil
}
