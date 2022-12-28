package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"flag"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gunsluo/go-example/ory/oauth-client/pkce"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/oauth2"
)

var (
	clientID     string
	clientSecret string
	endpoint     string
	port         int
	scopes       []string
	enablePKCE   bool
)

func main() {
	var scope string
	flag.StringVar(&clientID, "client-id", "", "Please provide a Client ID using -client-id flag.")
	flag.StringVar(&clientSecret, "client-secret", "", "Please provide a Client Secret using -client-secret flag.")
	flag.IntVar(&port, "port", 5555, "Please provide a Port using -port flag.")
	flag.StringVar(&scope, "scope", "", "Please provide a Scope using -scopes flag.")
	flag.StringVar(&endpoint, "endpoint", "", "Please provide a endpoint using -endpoint flag.")
	flag.BoolVar(&enablePKCE, "pkce", false, "enable pkce.")

	flag.Parse()

	var (
		redirectUrl string
		backend     string
		frontend    string
		audience    []string
		prompt      []string
		maxAge      int
	)

	proto := "http"
	scopes = strings.Split(scope, ",")
	serverLocation := fmt.Sprintf("%s://127.0.0.1:%d/", proto, port)
	if redirectUrl == "" {
		redirectUrl = serverLocation + "callback"
	}

	if backend == "" {
		backend = joinUrl(endpoint, "/oauth2/token")
	}
	if frontend == "" {
		frontend = joinUrl(endpoint, "/oauth2/auth")
	}
	logoutEndpoint := joinUrl(endpoint, "/oauth2/sessions/logout")
	logoutRedirectUri := serverLocation + "logout"
	refreshUrl := serverLocation + "refresh"
	checkUrl := serverLocation + "check"

	conf := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: backend,
			AuthURL:  frontend,
		},
		RedirectURL: redirectUrl,
		Scopes:      scopes,
	}

	var generateAuthCodeURL = func() (string, string) {
		nonce, err := RuneSequence(24, AlphaLower)
		if err != nil {
			panic(err)
		}

		state, err := genState("", 32)
		if err != nil {
			panic(err)
		}

		var authCodeURL string
		if enablePKCE {
			// generate code verifier and code chanllenge
			suite, err := pkce.Generate()
			if err != nil {
				panic(err)
			}

			fmt.Println("---->", state, suite.CodeVerifier, suite.CodeChallenge, suite.CodeChallengeMethod)

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

		return authCodeURL, state
	}
	authCodeURL, state := generateAuthCodeURL()

	fmt.Println("Setting up home route on " + serverLocation)
	fmt.Println("Setting up callback listener on " + serverLocation + "callback")
	fmt.Println("Press ctrl + c on Linux / Windows or cmd + c on OSX to end the process.")
	fmt.Printf("If your browser does not open automatically, navigate to:\n\n\t%s\n\n", serverLocation)

	r := httprouter.New()

	server := &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: r, TLSConfig: nil}
	var onDone = func() {
		// regenerate because we don't want to shutdown and we don't want to reuse nonce & state
		authCodeURL, state = generateAuthCodeURL()
	}

	r.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		_ = tokenUserWelcome.Execute(w, &struct{ URL string }{URL: authCodeURL})
	})

	type ed struct {
		Name        string
		Description string
		Hint        string
		Debug       string
	}

	r.GET("/callback", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		if len(r.URL.Query().Get("error")) > 0 {
			fmt.Printf("Got error: %s\n", r.URL.Query().Get("error_description"))

			w.WriteHeader(http.StatusInternalServerError)
			_ = tokenUserError.Execute(w, &ed{
				Name:        r.URL.Query().Get("error"),
				Description: r.URL.Query().Get("error_description"),
				Hint:        r.URL.Query().Get("error_hint"),
				Debug:       r.URL.Query().Get("error_debug"),
			})

			onDone()
			return
		}

		if r.URL.Query().Get("state") != string(state) {
			fmt.Printf("States do not match. Expected %s, got %s\n", string(state), r.URL.Query().Get("state"))

			w.WriteHeader(http.StatusInternalServerError)
			_ = tokenUserError.Execute(w, &ed{
				Name:        "States do not match",
				Description: "Expected state " + string(state) + " but got " + r.URL.Query().Get("state"),
			})
			onDone()
			return
		}

		code := r.URL.Query().Get("code")
		state := r.URL.Query().Get("state")
		var options []oauth2.AuthCodeOption
		if enablePKCE {
			codeVerifier, exist := authReqRelations.get(state)
			if exist {
				options = append(options, oauth2.SetAuthURLParam("code_verifier", codeVerifier))
			}
		}
		token, err := conf.Exchange(r.Context(), code, options...)
		if err != nil {
			fmt.Printf("Unable to exchange code for token: %s\n", err)

			w.WriteHeader(http.StatusInternalServerError)
			_ = tokenUserError.Execute(w, &ed{
				Name: err.Error(),
			})
			onDone()
			return
		}

		idt := token.Extra("id_token")
		var idToken string
		if idt != nil {
			idToken = fmt.Sprintf("%v", idt)
		}

		logoutUrl := logoutEndpoint + "?id_token_hint=" + idToken +
			"&post_logout_redirect_uri=" + logoutRedirectUri + "&state="
		checkUrlWithToken := checkUrl + "?access_token=" + token.AccessToken +
			"&id_token=" + idToken
		fmt.Printf("Access Token:\n\t%s\n", token.AccessToken)
		fmt.Printf("Refresh Token:\n\t%s\n", token.RefreshToken)
		fmt.Printf("Expires in:\n\t%s\n", token.Expiry.Format(time.RFC1123))
		fmt.Printf("ID Token:\n\t%v\n\n", idt)
		fmt.Printf("Logout Url:\n\t%s\n\n", logoutUrl)

		_ = tokenUserResult.Execute(w, struct {
			AccessToken       string
			RefreshToken      string
			Expiry            string
			IDToken           string
			BackURL           string
			DisplayBackButton bool
			LogoutURL         string
			RefreshUrl        string
			CheckUrl          string
		}{
			AccessToken:       token.AccessToken,
			RefreshToken:      token.RefreshToken,
			Expiry:            token.Expiry.Format(time.RFC1123),
			IDToken:           idToken,
			BackURL:           serverLocation,
			DisplayBackButton: true,
			LogoutURL:         logoutUrl,
			RefreshUrl:        refreshUrl,
			CheckUrl:          checkUrlWithToken,
		})
		onDone()
	})

	r.GET("/logout", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("logout successfully"))
	})

	r.POST("/refresh", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		refresh := r.FormValue("refresh_token")
		if refresh == "" {
			http.Error(w, fmt.Sprintf("no refresh_token in request: %q", r.Form), http.StatusBadRequest)
			return
		}
		t := &oauth2.Token{
			RefreshToken: refresh,
			Expiry:       time.Now().Add(time.Hour),
		}

		token, err := conf.TokenSource(r.Context(), t).Token()
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get token: %v", err), http.StatusInternalServerError)
			return
		}

		idt := token.Extra("id_token")
		idToken := fmt.Sprintf("%v", idt)
		logoutUrl := logoutEndpoint + "?id_token_hint=" + idToken +
			"&post_logout_redirect_uri=" + logoutRedirectUri + "&state="
		checkUrlWithToken := checkUrl + "?access_token=" + token.AccessToken +
			"&id_token=" + idToken
		fmt.Printf("Access Token:\n\t%s\n", token.AccessToken)
		fmt.Printf("Refresh Token:\n\t%s\n", token.RefreshToken)
		fmt.Printf("Expires in:\n\t%s\n", token.Expiry.Format(time.RFC1123))
		fmt.Printf("ID Token:\n\t%v\n\n", idt)
		fmt.Printf("Logout Url:\n\t%s\n\n", logoutUrl)

		_ = tokenUserResult.Execute(w, struct {
			AccessToken       string
			RefreshToken      string
			Expiry            string
			IDToken           string
			BackURL           string
			DisplayBackButton bool
			LogoutURL         string
			RefreshUrl        string
			CheckUrl          string
		}{
			AccessToken:       token.AccessToken,
			RefreshToken:      token.RefreshToken,
			Expiry:            token.Expiry.Format(time.RFC1123),
			IDToken:           idToken,
			BackURL:           serverLocation,
			DisplayBackButton: true,
			LogoutURL:         logoutUrl,
			RefreshUrl:        refreshUrl,
			CheckUrl:          checkUrlWithToken,
		})
	})

	oidcProvider, err := oidc.NewProvider(context.Background(), endpoint)
	if err != nil {
		panic(err)
	}

	//verifier = oidcProvider.Verifier(&oidc.Config{SkipClientIDCheck: true})
	verifier := oidcProvider.Verifier(&oidc.Config{ClientID: clientID})

	r.GET("/check", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// token from query or header
		idToken := r.URL.Query().Get("id_token")

		if idToken != "" {
			token, err := verifier.Verify(r.Context(), idToken)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			var raw json.RawMessage
			err = token.Claims(&raw)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			content := &bytes.Buffer{}
			content.WriteString("Authorized")
			content.WriteString("<br/>")
			content.WriteString("Subject: " + token.Subject)
			content.WriteString("<br/>")
			content.WriteString("Issuer: " + token.Issuer)
			content.WriteString("<br/>")
			content.WriteString("Audience: " + strings.Join(token.Audience, " "))
			content.WriteString("<br/>")
			content.WriteString("Expiry: " + token.Expiry.String())
			content.WriteString("<br/>")
			content.WriteString("IssuedAt: " + token.IssuedAt.String())
			content.WriteString("<br/>")
			content.WriteString("Nonce: " + token.Nonce)
			content.WriteString("<br/>")
			content.WriteString("Claims: " + string(raw))
			content.WriteString("<br/>")

			w.Header().Add("Content-Type", "text/html; charset=utf-8")
			w.Write(content.Bytes())
			return
		}

		accessToken := r.URL.Query().Get("access_token")

		ts := oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: accessToken,
			TokenType:   "Bearer",
		})

		userinfo, err := oidcProvider.UserInfo(r.Context(), ts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var raw json.RawMessage
		err = userinfo.Claims(&raw)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		content := &bytes.Buffer{}
		content.WriteString("Authorized")
		content.WriteString("<br/>")
		content.WriteString("Subject: " + userinfo.Subject)
		content.WriteString("<br/>")
		content.WriteString("Profile: " + userinfo.Profile)
		content.WriteString("<br/>")
		content.WriteString("Email: " + userinfo.Email)
		content.WriteString("<br/>")
		content.WriteString("EmailVerified: " + fmt.Sprintf("%v", userinfo.EmailVerified))
		content.WriteString("<br/>")
		content.WriteString("Claims: " + string(raw))
		content.WriteString("<br/>")

		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.Write(content.Bytes())
	})

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

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

var tokenUserWelcome = template.Must(template.New("").Parse(`<html>
<body>
<h1>Welcome to the exemplary OAuth 2.0 Consumer!</h1>
<p>This is an example app which emulates an OAuth 2.0 consumer application. Usually, this would be your web or mobile
    application and would use an <a href="https://oauth.net/code/">OAuth 2.0</a> or <a href="https://oauth.net/code/">OpenID
        Connect</a> library.</p>
<p>This example requests an OAuth 2.0 Access, Refresh, and OpenID Connect ID Token from the OAuth 2.0 Server (ORY
    Hydra).
    To initiate the flow, click the "Authorize Application" button.</p>
<p><a href="{{ .URL }}">Authorize application</a></p>
</body>
</html>`))

var tokenUserError = template.Must(template.New("").Parse(`<html>
<body>
<h1>An error occurred</h1>
<h2>{{ .Name }}</h2>
<p>{{ .Description }}</p>
<p>{{ .Hint }}</p>
<p>{{ .Debug }}</p>
</body>
</html>`))

var tokenUserResult = template.Must(template.New("").Parse(`<html>
<head></head>
<body>
<ul>
    <li>Access Token: <code>{{ .AccessToken }}</code></li>
    <li>Refresh Token: <code>{{ .RefreshToken }}</code></li>
    <li>Expires in: <code>{{ .Expiry }}</code></li>
    <li>ID Token: <code>{{ .IDToken }}</code></li>
</ul>
{{ if .DisplayBackButton }}
<a href="{{ .BackURL }}">Back to Welcome Page</a>
{{ end }}

<br/>
	<form action="{{ .RefreshUrl }}" method="post">
	  <input type="hidden" name="refresh_token" value="{{ .RefreshToken }}">
	  <input type="submit" value="Redeem refresh token">
    </form>

<br/>
<a href="{{ .CheckUrl }}">Check Token</a>

<br/>
<a href="{{ .LogoutURL }}">Logout</a>
</body>
</html>`))
