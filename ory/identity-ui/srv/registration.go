package srv

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func (s *Server) registration(c *gin.Context) {
	flowId := c.Query("flow")
	if flowId == "" {
		vs := url.Values{}
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

	// debug
	s.debugPrint("registration ui", flow.Ui)
	froms := groupRegistrationUi(flow.Ui)
	c.HTML(http.StatusOK, "registration.html", gin.H{
		"ui":       froms,
		"loginUrl": "/login",
	})
}

func (s *Server) gotoRegistration(c *gin.Context, vs url.Values) {
	redirectUrl := fmt.Sprintf("%s/self-service/registration/browser?%s", s.identityEndpoint, vs.Encode())
	c.Redirect(http.StatusSeeOther, redirectUrl)
}
