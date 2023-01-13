package srv

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gunsluo/go-example/ory/identity-ui/swagger/identityclient"
)

type Server struct {
	dev bool

	dir              string
	identityEndpoint string
	port             int
	logoutReturnTo   string

	apiClient *identityclient.APIClient
	engine    *gin.Engine
}

func New() (*Server, error) {
	s := &Server{dev: true, dir: "./tmpl"}

	identityEndpoint := os.Getenv(envIdentityEndpoint)
	_, err := url.Parse(identityEndpoint)
	if err != nil {
		log.Fatalf("unable to parse identity endpoint[%s]", identityEndpoint)
	}
	s.identityEndpoint = identityEndpoint

	srvPort := os.Getenv(envSrvPort)
	if port, err := strconv.Atoi(srvPort); err != nil {
		log.Fatalf("unable to parse srv port[%s]", srvPort)
	} else {
		s.port = port
	}

	configuration := identityclient.NewConfiguration()
	configuration.Servers = identityclient.ServerConfigurations{{URL: identityEndpoint}}
	s.apiClient = identityclient.NewAPIClient(configuration)

	s.setRouter()
	return s, nil
}

func (s *Server) setRouter() {
	s.engine = gin.New()
	s.engine.Use(gin.Logger(), gin.Recovery())

	s.engine.LoadHTMLGlob("tmpl/*.html")

	// static file(css js...)
	assetsPath := fmt.Sprintf("%s/assets", s.dir)
	s.engine.StaticFS("/assets", http.Dir(assetsPath))

	authmd := &authMiddleware{
		getApiClient: func() *identityclient.APIClient {
			return s.apiClient
		},
		unauthHandle: func(c *gin.Context, uai UnAuthInfo) {
			// redirect to login page
			c.Redirect(http.StatusSeeOther, "/login")
		},
	}
	s.engine.GET("/", authmd.Handler(true), s.index)

	// login
	s.engine.GET("/login", s.login)

	// logout
	s.engine.GET("/logout", s.logout)

	// sessions
	s.engine.GET("/sessions", authmd.Handler(), s.session)

	// registration
	s.engine.GET("/registration", s.registration)

	// verification
	s.engine.GET("/verification", s.verification)

	// recovery
	s.engine.GET("/recovery", s.recovery)

	// settings
	s.engine.GET("/settings", authmd.Handler(), s.settings)

	// execption
	s.engine.GET("/execption", s.execption)
}

func (s *Server) Run() {
	// listen and serve on 0.0.0.0:
	s.engine.Run(fmt.Sprintf(":%d", s.port))
}

const (
	envIdentityEndpoint = "IDENTITY_ENDPOINT"
	envSrvPort          = "PORT"
)
