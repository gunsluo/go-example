package srv

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func (s *Server) recovery(c *gin.Context) {
	flowId := c.Query("flow")
	if flowId == "" {
		vs := url.Values{}
		s.gotoRecovery(c, vs)
		return
	}

	ctx := c.Request.Context()
	cookie := c.Request.Header.Get("cookie")
	flow, _, err := s.apiClient.RecoveryApi.GetRecoveryFlowRequest(ctx).
		Id(flowId).
		Cookie(cookie).
		Execute()
	if err != nil {
		s.gotoExecptionWithError(c, err)
		return
	}

	// debug
	s.debugPrint("recovery ui", flow.Ui)
	froms := groupRecovery(flow.Ui)
	c.HTML(http.StatusOK, "recovery.html", gin.H{
		"ui":       froms,
		"loginUrl": "/login",
	})
}

func (s *Server) gotoRecovery(c *gin.Context, vs url.Values) {
	redirectUrl := fmt.Sprintf("%s/self-service/recovery/browser?%s", s.identityEndpoint, vs.Encode())
	c.Redirect(http.StatusSeeOther, redirectUrl)
}
