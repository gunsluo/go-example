package srv

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func (s *Server) verification(c *gin.Context) {
	flowId := c.Query("flow")
	if flowId == "" {
		vs := url.Values{}
		loginChallenge := c.Query("login_challenge")
		if loginChallenge != "" {
			vs.Add("login_challenge", loginChallenge)
		}
		s.gotoVerification(c, vs)
		return
	}

	ctx := c.Request.Context()
	cookie := c.Request.Header.Get("cookie")
	flow, _, err := s.apiClient.VerificationApi.GetVerificationFlowRequest(ctx).
		Id(flowId).
		Cookie(cookie).
		Execute()
	if err != nil {
		s.gotoExecptionWithError(c, err)
		return
	}

	// debug
	s.debugPrint("verification ui", flow.Ui)
	froms := groupVerificationUi(flow.Ui)
	c.HTML(http.StatusOK, "verification.html", gin.H{
		"ui":              froms,
		"state":           flow.State,
		"loginUrl":        "/login?login_challenge=" + flow.GetOauth2LoginChallenge(),
		"registrationUrl": "/registration?login_challenge=" + flow.GetOauth2LoginChallenge(),
	})
}

func (s *Server) gotoVerification(c *gin.Context, vs url.Values) {
	redirectUrl := fmt.Sprintf("%s/self-service/verification/browser?%s", s.identityEndpoint, vs.Encode())
	c.Redirect(http.StatusSeeOther, redirectUrl)
}
