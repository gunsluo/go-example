package srv

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/gunsluo/go-example/ory/identity-ui/swagger/identityclient"
)

func (s *Server) registration(c *gin.Context) {
	flowId := c.Query("flow")
	if flowId == "" {
		vs := url.Values{}
		loginChallenge := c.Query("login_challenge")
		if loginChallenge != "" {
			vs.Add("login_challenge", loginChallenge)
		}

		s.gotoRegistration(c, vs)
		return
	}

	ctx := c.Request.Context()
	cookie := c.Request.Header.Get("cookie")

	flow, _, err := s.apiClient.RegistrationApi.GetRegistrationFlowRequest(ctx).
		Id(flowId).
		Cookie(cookie).
		Execute()
	if err != nil {
		s.gotoExecptionWithError(c, err)
		return
	}

	var oauthClient *identityclient.OAuth2Client
	if flow.Oauth2LoginRequest != nil {
		oauthClient = flow.Oauth2LoginRequest.Client
	}

	// debug
	s.debugPrint("registration ui", flow.Ui)
	froms := groupRegistrationUi(flow.Ui)
	c.HTML(http.StatusOK, "registration.html", gin.H{
		"oauthClient": oauthClient,
		"ui":          froms,
		"loginUrl":    "/login?login_challenge=" + flow.GetOauth2LoginChallenge(),
	})
}

func (s *Server) gotoRegistration(c *gin.Context, vs url.Values) {
	redirectUrl := fmt.Sprintf("%s/self-service/registration/browser?%s", s.identityEndpoint, vs.Encode())
	c.Redirect(http.StatusSeeOther, redirectUrl)
}
