package main

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"sync"

	"github.com/gunsluo/go-example/consent-go/ropc/pkce"
	"github.com/ory/hydra-client-go/client"
	"github.com/ory/hydra-client-go/client/admin"
	"github.com/ory/hydra-client-go/models"
	"golang.org/x/oauth2"
)

var (
	endpoint    string
	port        int
	scopes      []string
	redirectUrl string
	backend     string
	frontend    string

	serverLocation string

	adminUrl    string
	adminClient *client.OryHydra
)

func main() {
	var scope string
	flag.IntVar(&port, "port", 8888, "Please provide a Port using -port flag.")
	flag.StringVar(&scope, "scope", "", "Please provide a Scope using -scopes flag.")
	flag.StringVar(&endpoint, "endpoint", "", "Please provide a endpoint using -endpoint flag.")
	flag.StringVar(&adminUrl, "admin", "", "Please provide a endpoint using -admin flag.")

	flag.Parse()

	proto := "http"
	scopes = strings.Split(scope, ",")
	serverLocation = fmt.Sprintf("%s://127.0.0.1:%d/", proto, port)

	redirectUrl = serverLocation + "callback"
	backend = joinUrl(endpoint, "/oauth2/token")
	frontend = joinUrl(endpoint, "/oauth2/auth")

	adminURL, err := url.Parse(adminUrl)
	if err != nil {
		panic(err)
	}
	hydraAdminClient := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{Schemes: []string{adminURL.Scheme}, Host: adminURL.Host, BasePath: adminURL.Path})
	adminClient = hydraAdminClient

	mux := http.NewServeMux()
	mux.HandleFunc("/", loginHandler)
	mux.HandleFunc("/token", tokenHandler)

	fmt.Println("Setting up home route on " + serverLocation)
	fmt.Println("Press ctrl + c on Linux / Windows or cmd + c on OSX to end the process.")
	fmt.Printf("If your browser does not open automatically, navigate to:\n\n\t%s\n\n", serverLocation)

	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		_ = mockLoginPage.Execute(w, &struct{ URL string }{URL: "/"})
		return
	}

	if req.Method == http.MethodPost {
		clientId := req.FormValue("clientId")
		clientSecret := req.FormValue("clientSecret")
		username := req.FormValue("username")
		password := req.FormValue("password")

		// call token api
		tokenUrl := fmt.Sprintf("%stoken?grant_type=%s&username=%s&password=%s", serverLocation, "password", username, password)
		r, err := http.NewRequest(http.MethodPost, tokenUrl, nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
			return
		}
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.SetBasicAuth(clientId, clientSecret)

		client := http.DefaultClient
		resp, err := client.Do(r.WithContext(req.Context()))
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
			return
		}

		if resp.StatusCode != http.StatusOK {
			http.Error(w, fmt.Sprintf("code: %d, err: %s", resp.StatusCode, string(body)), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json;charset=UTF-8")
		w.Write(body)
		return
	}

	http.Error(w, "Method not allowed", http.StatusBadRequest)
}

func tokenHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	clientId, clientSecret, ok := req.BasicAuth()
	if clientId == "" || !ok {
		http.Error(w, fmt.Sprintf("missing client Id"), http.StatusBadRequest)
		return
	}
	// TODO: check clientId

	grantType := req.URL.Query().Get("grant_type")
	if grantType != "password" {
		http.Error(w, fmt.Sprintf("invalid grant_type"), http.StatusBadRequest)
		return
	}
	username := req.URL.Query().Get("username")
	password := req.URL.Query().Get("password")

	if username != "foo@bar.com" && password != "foobar" {
		http.Error(w, "The username / password combination is not correct", http.StatusUnauthorized)
		return
	}

	// token proxy
	pwdToken, err := tokenProxy(clientId, clientSecret, username)
	if err != nil {
		http.Error(w, fmt.Sprintf("error:%v", err), http.StatusBadRequest)
		return
	}

	buffer, err := json.Marshal(pwdToken)
	if err != nil {
		http.Error(w, fmt.Sprintf("error: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(buffer)
}

type passwordToken struct {
	AccessToken  string `json:"access_token"`
	IdToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
}

func tokenProxy(clientId, clientSecret, subject string) (*passwordToken, error) {
	conf := oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: backend,
			AuthURL:  frontend,
		},
		RedirectURL: redirectUrl,
		Scopes:      scopes,
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	httpClient := &http.Client{
		Jar: jar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, HTTPClient, httpClient)

	// 1. redirect to login page
	authCodeURL, state, err := generateAuthCodeURL(conf)
	if err != nil {
		return nil, err
	}

	redirectReqUrl, err := getLocationRequest(ctx, authCodeURL, http.MethodGet)
	if err != nil {
		return nil, err
	}

	// accept login
	loginChallenge := redirectReqUrl.Query().Get("login_challenge")
	if loginChallenge == "" {
		return nil, fmt.Errorf("missing login_challenge")
	}

	acceptRequest := admin.NewAcceptLoginRequestParams().
		WithLoginChallenge(loginChallenge).
		WithContext(ctx).
		WithBody(&models.AcceptLoginRequest{
			Subject:     &subject,
			Remember:    false,
			RememberFor: 3600,
			Acr:         "0",
		})
	acceptReply, err := adminClient.Admin.AcceptLoginRequest(acceptRequest)
	if err != nil {
		return nil, err
	}
	if acceptReply == nil || acceptReply.Payload == nil || acceptReply.Payload.RedirectTo == nil {
		return nil, fmt.Errorf("invalid response from accept login")
	}

	acceptLoginRedirectUrl := *acceptReply.Payload.RedirectTo

	// 2. redirect to consent page
	consentRedirectReqUrl, err := getLocationRequest(ctx, acceptLoginRedirectUrl, http.MethodGet)
	if err != nil {
		return nil, err
	}

	consentChallenge := consentRedirectReqUrl.Query().Get("consent_challenge")
	if consentChallenge == "" {
		return nil, fmt.Errorf("missing consent_challenge")
	}

	// accept consent
	consentRequest := admin.NewGetConsentRequestParams().
		WithConsentChallenge(consentChallenge).
		WithContext(ctx)
	consentReply, err := adminClient.Admin.GetConsentRequest(consentRequest)
	if err != nil {
		return nil, err
	}

	if consentReply == nil || consentReply.Payload == nil {
		return nil, fmt.Errorf("invalid response from get consent")
	}

	acceptConsentRequest := admin.NewAcceptConsentRequestParams().
		WithConsentChallenge(consentChallenge).
		WithContext(ctx).
		WithBody(&models.AcceptConsentRequest{
			GrantScope:               scopes,
			Remember:                 false,
			RememberFor:              3600,
			Session:                  oidcConformityMaybeFakeSession(consentReply, scopes),
			GrantAccessTokenAudience: consentReply.Payload.RequestedAccessTokenAudience,
		})
	acceptConsentReply, err := adminClient.Admin.AcceptConsentRequest(acceptConsentRequest)
	if err != nil {
		return nil, err
	}

	if acceptConsentReply == nil || acceptConsentReply.Payload == nil || acceptConsentReply.Payload.RedirectTo == nil {
		return nil, fmt.Errorf("invalid response from accept consent")
	}

	acceptConsentRedirectUrl := *acceptConsentReply.Payload.RedirectTo

	// 3. get to callback url
	callbackRedirectReqUrl, err := getLocationRequest(ctx, acceptConsentRedirectUrl, http.MethodGet)
	if err != nil {
		return nil, err
	}

	code := callbackRedirectReqUrl.Query().Get("code")
	stateInReq := callbackRedirectReqUrl.Query().Get("state")
	if stateInReq != state {
		return nil, fmt.Errorf("state mismatch")
	}

	// 4. get token using code

	var options []oauth2.AuthCodeOption
	if clientSecret == "" {
		codeVerifier, exist := authReqRelations.get(state)
		if exist {
			options = append(options, oauth2.SetAuthURLParam("code_verifier", codeVerifier))
		}
	}
	token, err := conf.Exchange(ctx, code, options...)
	if err != nil {
		fmt.Printf("Unable to exchange code for token: %s\n", err)
		return nil, err
	}

	idt := token.Extra("id_token")
	var idToken string
	if idt != nil {
		idToken = fmt.Sprintf("%v", idt)
	}

	pwdToken := &passwordToken{
		AccessToken:  token.AccessToken,
		IdToken:      idToken,
		RefreshToken: token.RefreshToken,
		TokenType:    token.TokenType,
		ExpiresIn:    token.Expiry.Unix(),
	}
	return pwdToken, nil
}

func generateAuthCodeURL(conf oauth2.Config) (string, string, error) {
	nonce, err := RuneSequence(24, AlphaLower)
	if err != nil {
		return "", "", err
	}

	state, err := genState("", 32)
	if err != nil {
		return "", "", err
	}

	var (
		audience []string
		prompt   []string
		maxAge   int
	)

	var authCodeURL string
	if conf.ClientSecret == "" {
		// generate code verifier and code chanllenge
		suite, err := pkce.Generate()
		if err != nil {
			return "", "", err
		}

		authReqRelations.save(state, suite.CodeVerifier)
		authCodeURL = conf.AuthCodeURL(
			state,
			oauth2.SetAuthURLParam("audience", strings.Join(audience, "+")),
			oauth2.SetAuthURLParam("nonce", string(nonce)),
			oauth2.SetAuthURLParam("prompt", strings.Join(prompt, "+")),
			oauth2.SetAuthURLParam("max_age", strconv.Itoa(maxAge)),
			oauth2.SetAuthURLParam("code_challenge", suite.CodeChallenge),
			oauth2.SetAuthURLParam("code_challenge_method", suite.CodeChallengeMethod),
		)
	} else {

		authCodeURL = conf.AuthCodeURL(
			state,
			oauth2.SetAuthURLParam("audience", strings.Join(audience, "+")),
			oauth2.SetAuthURLParam("nonce", string(nonce)),
			oauth2.SetAuthURLParam("prompt", strings.Join(prompt, "+")),
			oauth2.SetAuthURLParam("max_age", strconv.Itoa(maxAge)),
		)
	}

	return authCodeURL, state, nil
}

func getLocationRequest(ctx context.Context, url, method string) (*url.URL, error) {
	resp, err := proxyRequest(ctx, url, method)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusMovedPermanently &&
		resp.StatusCode != http.StatusFound &&
		resp.StatusCode != http.StatusSeeOther &&
		resp.StatusCode != http.StatusTemporaryRedirect {
		return nil, fmt.Errorf("fetch url, error: %s %s", url, method)
	}

	redirectUrl, err := resp.Location()
	if err != nil {
		return nil, err
	}

	return redirectUrl, nil

	//if resp.StatusCode != http.StatusOK {
	//	return nil, fmt.Errorf("fetch url, error: %s %s", url, method)
	//}

	//return resp.Request.URL, nil
}

var HTTPClient struct{}

func proxyRequest(ctx context.Context, url, method string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create GET request: %v", err)
	}

	client := http.DefaultClient
	if c, ok := ctx.Value(HTTPClient).(*http.Client); ok {
		client = c
	}

	return client.Do(req.WithContext(ctx))
}

func doRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	client := http.DefaultClient
	if c, ok := ctx.Value(oauth2.HTTPClient).(*http.Client); ok {
		client = c
	}
	return client.Do(req.WithContext(ctx))
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
	// AlphaNum contains runes [abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789].
	AlphaNum = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	// Alpha contains runes [abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ].
	Alpha = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	// AlphaLowerNum contains runes [abcdefghijklmnopqrstuvwxyz0123456789].
	AlphaLowerNum = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	// AlphaUpperNum contains runes [ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789].
	AlphaUpperNum = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	// AlphaLower contains runes [abcdefghijklmnopqrstuvwxyz].
	AlphaLower = []rune("abcdefghijklmnopqrstuvwxyz")
	// AlphaUpper contains runes [ABCDEFGHIJKLMNOPQRSTUVWXYZ].
	AlphaUpper = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	// Numeric contains runes [0123456789].
	Numeric = []rune("0123456789")
)

var rander = rand.Reader // random function

// RuneSequence returns a random sequence using the defined allowed runes.
func RuneSequence(l int, allowedRunes []rune) (seq []rune, err error) {
	c := big.NewInt(int64(len(allowedRunes)))
	seq = make([]rune, l)

	for i := 0; i < l; i++ {
		r, err := rand.Int(rander, c)
		if err != nil {
			return seq, err
		}
		rn := allowedRunes[r.Uint64()]
		seq[i] = rn
	}

	return seq, nil
}

func genState(prefix string, l int) (string, error) {
	seq, err := RuneSequence(l, AlphaLowerNum)
	if err != nil {
		return "", err
	}

	if prefix == "" {
		return string(seq), nil
	}

	return prefix + string(seq), nil
}

type requestRelation struct {
	lock      sync.RWMutex
	relations map[string]string
}

var authReqRelations = requestRelation{relations: make(map[string]string, 10)}

func (m *requestRelation) save(state string, codeVerifier string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.relations[state] = codeVerifier
}

func (m *requestRelation) get(state string) (string, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	codeVerifier, ok := m.relations[state]
	if !ok {
		return "", false
	}

	return codeVerifier, true
}

var mockLoginPage = template.Must(template.New("").Parse(`<!DOCTYPE html>
<html>

<head>
  <title></title>
</head>

<body>
  <h1 id="login-title">Please log in</h1>
  <form action="{{ .URL }}" method="POST">
    <table>
      <tr>
        <td><input type="text" id="clientId" name="clientId" value="auth-code-client" placeholder="public-client"></td>
        <td>(auth-code-client)</td>
      </tr>
      <tr>
        <td><input type="text" id="clientSecret" name="clientSecret" value="secret" placeholder="secret"></td>
        <td>(secret)</td>
      </tr>
      <tr>
        <td><input type="text" id="username" name="username" value="foo@bar.com" placeholder="email@foobar.com"></td>
        <td>(it's "foo@bar.com")</td>
      </tr>
      <tr>
        <td><input type="password" id="password" name="password"></td>
        <td>(it's "foobar")</td>
      </tr>
    </table>
	  <input type="submit" id="accept" name="submit" value="Log in">
  </form>
</body>

</html>`))

func joinUrl(base, path string) string {
	b := strings.HasSuffix(base, "/")
	p := strings.HasPrefix(path, "/")
	switch {
	case b && p:
		return base + path[1:]
	case b || p:
		return base + path
	default:
		return base + "/" + path
	}
}
