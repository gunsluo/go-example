package srv

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type navTmplValues struct {
	SessionUrl  string `json:"sessionUrl"`
	SignInUrl   string `json:"signInUrl"`
	SignUpUrl   string `json:"signUpUrl"`
	RecoveryUrl string `json:"recoveryUrl"`
	VerifyUrl   string `json:"verifyUrl"`
	SettingsUrl string `json:"settingsUrl"`
	LogoutUrl   string `json:"logoutUrl"`
}

func (s *Server) index(c *gin.Context) {
	ctx := c.Request.Context()

	session := AuthSession(ctx)
	if session == nil {
		uai := UnAuth(ctx)
		if uai.AskAal2 {
			vs := url.Values{}
			vs.Add("aal", "aal2")

			s.gotoLogin(c, vs)
			return
		}

		// display ineex.html
		c.HTML(http.StatusOK, "index.html", gin.H{
			"authed": false,
			"nav": navTmplValues{
				SessionUrl:  "/sessions",
				SignInUrl:   "/login",
				SignUpUrl:   "/registration",
				RecoveryUrl: "/recovery",
				VerifyUrl:   "/verification",
			},
		})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"authed": true,
		"nav": navTmplValues{
			SessionUrl:  "/sessions",
			VerifyUrl:   "/verification",
			SettingsUrl: "/settings",
			LogoutUrl:   "/logout",
		},
	})
}
