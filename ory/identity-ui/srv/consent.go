package srv

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func (s *Server) consent(c *gin.Context) {
	flowId := c.Query("flow")
	if flowId == "" {
		vs := url.Values{}

		consentChallenge := c.Query("consent_challenge")
		vs.Add("consent_challenge", consentChallenge)

		s.gotoConsent(c, vs)
		return
	}

	ctx := c.Request.Context()
	cookie := c.Request.Header.Get("cookie")
	flow, _, err := s.apiClient.ConsentApi.GetConsentFlowRequest(ctx).
		Id(flowId).
		Cookie(cookie).
		Execute()
	if err != nil {
		s.gotoExecptionWithError(c, err)
		return
	}

	if flow.Oauth2ConsentRequest == nil {
		s.gotoExecptionWithError(c, errors.New("invalid consent"))
		return
	}
	oauthClient := flow.Oauth2ConsentRequest.Client

	// debug
	s.debugPrint("consent ui", flow.Ui)
	froms := groupConsentUi(flow.Ui)
	c.HTML(http.StatusOK, "consent.html", gin.H{
		"identity":    flow.Identity,
		"oauthClient": oauthClient,
		"ui":          froms,
		// "logoutUrl":       "/logout",
	})
}

func (s *Server) gotoConsent(c *gin.Context, vs url.Values) {
	redirectUrl := fmt.Sprintf("%s/self-service/consent/browser?%s", s.identityEndpoint, vs.Encode())

	c.Redirect(http.StatusSeeOther, redirectUrl)
}
