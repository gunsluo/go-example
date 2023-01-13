package srv

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func (s *Server) settings(c *gin.Context) {
	flowId := c.Query("flow")
	if flowId == "" {
		vs := url.Values{}
		s.gotoSettings(c, vs)
		return
	}

	ctx := c.Request.Context()
	cookie := c.Request.Header.Get("cookie")
	flow, _, err := s.apiClient.SettingsApi.GetSettingsFlowRequest(ctx).
		Id(flowId).
		Cookie(cookie).
		Execute()
	if err != nil {
		s.gotoExecptionWithError(c, err)
		return
	}

	// debug
	s.debugPrint("settings ui", flow.Ui)
	froms, bind := groupSettingsUi(flow.Ui)
	c.HTML(http.StatusOK, "settings.html", gin.H{
		"ui":        froms,
		"bind":      bind,
		"logoutUrl": "/logout",
	})
}

func (s *Server) gotoSettings(c *gin.Context, vs url.Values) {
	redirectUrl := fmt.Sprintf("%s/self-service/settings/browser?%s", s.identityEndpoint, vs.Encode())
	c.Redirect(http.StatusSeeOther, redirectUrl)
}
