package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	ginspnego "github.com/gunsluo/go-example/kerberos/ginspengo"
	"github.com/jcmturner/goidentity/v6"
	"github.com/jcmturner/gokrb5/v8/client"
	"github.com/jcmturner/gokrb5/v8/config"
	"github.com/jcmturner/gokrb5/v8/credentials"
	"github.com/jcmturner/gokrb5/v8/gssapi"
	"github.com/jcmturner/gokrb5/v8/keytab"
	"github.com/jcmturner/gokrb5/v8/service"
	"github.com/jcmturner/gokrb5/v8/spnego"
)

const (
	realm  = "TEST.KRB5.COM"
	domain = "test.krb5.com"
	spn    = "sso.test.krb5.com"
	port   = ":80"
)

var (
	krb5Conf  *config.Config
	srvKeytab *keytab.Keytab
)

func main() {
	krb5ConfPath := "./docker/keytabs/krb5.conf"
	b1, err := ioutil.ReadFile(krb5ConfPath)
	if err != nil {
		panic(err)
	}

	conf, err := config.NewFromString(string(b1))
	if err != nil {
		panic(err)
	}
	conf.LibDefaults.NoAddresses = true
	krb5Conf = conf

	// service keytab
	filename := fmt.Sprintf("./docker/keytabs/%s.svc.keytab", spn)
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	kt := keytab.New()
	if err := kt.Unmarshal(b); err != nil {
		panic(err)
	}
	srvKeytab = kt

	r := gin.Default()
	pr := r.Group("/")
	pr.Use(ginspnego.SPNEGOKRB5AuthenticateHandler(kt,
		service.SessionManager(NewSessionMgr("ssokrb5")),
	))

	pr.GET("/", index)
	r.GET("/login", renderLogin)
	r.POST("/login", login)

	r.Run(port)
}

func index(c *gin.Context) {
	creds := goidentity.FromHTTPRequestContext(c.Request)
	c.String(http.StatusOK, fmt.Sprintf(`<html>
<h1>GOKRB5 Handler</h1>
<ul>
<li>Authenticed user: %s</li>
<li>User's realm: %s</li>
<li>Authn time: %v</li>
<li>Session ID: %s</li>
<ul>
</html>`,
		creds.UserName(),
		creds.Domain(),
		creds.AuthTime(),
		creds.SessionID(),
	))
}

func renderLogin(c *gin.Context) {
	c.Header("content-type", "text/html")
	c.String(http.StatusOK, loginHtml)
}

type loginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func login(c *gin.Context) {
	var param loginRequest
	if err := c.ShouldBind(&param); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	cl := client.NewWithPassword(param.Username, realm, param.Password, krb5Conf)
	if err := cl.Login(); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	st, err := getSpnegoToken(cl, "HTTP/"+spn)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	creds, err := getCredentialsFromContextTokena(srvKeytab, st)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.String(http.StatusOK, fmt.Sprintf(`<html>
<h1>GOKRB5 Handler</h1>
<ul>
<li>Authenticed user: %s</li>
<li>User's realm: %s</li>
<li>Authn time: %v</li>
<li>Session ID: %s</li>
<ul>
</html>`,
		creds.UserName(),
		creds.Domain(),
		creds.AuthTime(),
		creds.SessionID(),
	))
}

func getSpnegoToken(cl *client.Client, spn string) (gssapi.ContextToken, error) {
	s := spnego.SPNEGOClient(cl, spn)

	if err := s.AcquireCred(); err != nil {
		return nil, err
	}

	st, err := s.InitSecContext()
	if err != nil {
		return nil, fmt.Errorf("could not initialize context: %v", err)
	}

	return st, nil
}

func getCredentialsFromContextTokena(kt *keytab.Keytab, token gssapi.ContextToken) (*credentials.Credentials, error) {
	s := spnego.SPNEGOService(kt)

	// Validate the context token
	authed, ctx, status := s.AcceptSecContext(token)
	if status.Code != gssapi.StatusComplete && status.Code != gssapi.StatusContinueNeeded {
		return nil, fmt.Errorf("SPNEGO validation error: %v", status)
	}
	if status.Code == gssapi.StatusContinueNeeded {
		return nil, fmt.Errorf("SPNEGO GSS-API continue needed")
	}

	if !authed {
		// If we get to here we have not authenticationed so just reject
		return nil, fmt.Errorf("SPNEGO Kerberos authentication failed")
	}

	// Authentication successful; get user's credentials from the context
	id, ok := ctx.Value(ctxCredentials).(*credentials.Credentials)
	if !ok {
		return nil, fmt.Errorf("SPNEGO Kerberos credentials not found")
	}
	return id, nil
}

const (
	ctxCredentials = "github.com/jcmturner/gokrb5/v8/ctxCredentials"
)

type SessionMgr struct {
	skey       []byte
	store      sessions.Store
	cookieName string
}

func NewSessionMgr(cookieName string) SessionMgr {
	skey := []byte("thisistestsecret") // Best practice is to load this key from a secure location.
	return SessionMgr{
		skey:       skey,
		store:      sessions.NewCookieStore(skey),
		cookieName: cookieName,
	}
}

func (smgr SessionMgr) Get(r *http.Request, k string) ([]byte, error) {
	s, err := smgr.store.Get(r, smgr.cookieName)
	if err != nil {
		return nil, err
	}
	if s == nil {
		return nil, errors.New("nil session")
	}
	b, ok := s.Values[k].([]byte)
	if !ok {
		return nil, fmt.Errorf("could not get bytes held in session at %s", k)
	}
	return b, nil
}

func (smgr SessionMgr) New(w http.ResponseWriter, r *http.Request, k string, v []byte) error {
	s, err := smgr.store.New(r, smgr.cookieName)
	if err != nil {
		return fmt.Errorf("could not get new session from session manager: %v", err)
	}
	s.Values[k] = v
	return s.Save(r, w)
}

var loginHtml = `<html>

  <body>
    <form action="/login" method="post">
      <p>
        Username:<input type="text" name="username" value="luoji" placeholder="">
      </p>
      <p>
      Password: <input type="text" name="password" value="password" placeholder="password">
      </p>
      <input type="submit" value="Login">
    </form>
  </body>

  </html>`
