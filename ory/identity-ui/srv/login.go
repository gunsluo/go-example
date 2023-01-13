package srv

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/gunsluo/go-example/ory/identity-ui/swagger/identityclient"
)

func (s *Server) login(c *gin.Context) {
	flowId := c.Query("flow")
	if flowId == "" {
		vs := url.Values{}
		s.gotoLogin(c, vs)
		return
	}

	ctx := c.Request.Context()
	cookie := c.Request.Header.Get("cookie")
	flow, _, err := s.apiClient.LoginApi.GetLoginFlowRequest(ctx).
		Id(flowId).
		Cookie(cookie).
		Execute()
	if err != nil {
		s.gotoExecptionWithError(c, err)
		return
	}

	// debug
	s.debugPrint("login ui", flow.Ui)
	froms := groupLoginUi(flow.Ui)
	c.HTML(http.StatusOK, "login.html", gin.H{
		"ui":              froms,
		"aal":             flow.GetRequestedAal(),
		"logoutUrl":       "/logout",
		"recoveryUrl":     "/recovery",
		"registrationUrl": "/registration",
	})
}

func (s *Server) session(c *gin.Context) {
	ctx := c.Request.Context()
	cookie := c.Request.Header.Get("cookie")

	session, _, err := s.apiClient.SessionApi.ToSessionRequest(ctx).
		//XSessionToken().
		Cookie(cookie).
		Execute()
	if err != nil {
		s.gotoExecptionWithError(c, err)
		return
	}

	sessPretty, err := json.MarshalIndent(session, "", "  ")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "session.html", gin.H{
		"authed":  true,
		"session": string(sessPretty),
		"nav": navTmplValues{
			SessionUrl:  "#",
			SettingsUrl: "/settings",
			LogoutUrl:   "/logout",
		},
	})
}

func (s *Server) logout(c *gin.Context) {
	ctx := c.Request.Context()
	cookie := c.Request.Header.Get("cookie")

	logoutFlow, _, err := s.apiClient.LogoutApi.InitBrowserLogoutFlowRequest(ctx).
		Cookie(cookie).
		Execute()
	if err != nil {
		s.gotoExecptionWithError(c, err)
		return
	}

	var logoutUrl string
	if uri, err := url.Parse(logoutFlow.LogoutUrl); err != nil {
		s.gotoExecptionWithError(c, err)
		return
	} else {
		if s.logoutReturnTo != "" {
			// set return_to, it can also be set without this
			q := uri.Query()
			q.Set("return_to", s.logoutReturnTo)
			uri.RawQuery = q.Encode()
			logoutUrl = uri.String()
		} else {
			logoutUrl = logoutFlow.LogoutUrl
		}
	}

	c.Redirect(http.StatusSeeOther, logoutUrl)
}

func (s *Server) gotoLogin(c *gin.Context, vs url.Values) {
	redirectUrl := fmt.Sprintf("%s/self-service/login/browser?%s", s.identityEndpoint, vs.Encode())

	c.Redirect(http.StatusSeeOther, redirectUrl)
}

func (s *Server) gotoExecptionWithError(c *gin.Context, err error) {
	vs := url.Values{}
	if e := new(identityclient.GenericOpenAPIError); errors.As(err, &e) {
		if jer, ok := e.Model().(identityclient.JsonErrorResponse); ok {
			vs.Add("code", fmt.Sprintf("%v", jer.GetCode()))
			vs.Add("message", jer.GetMsg())
		}
	} else {
		vs.Add("code", fmt.Sprintf("%v", identityclient.CODE__5000))
		vs.Add("message", err.Error())
	}

	s.gotoExecption(c, vs)
}

func (s *Server) gotoExecption(c *gin.Context, vs url.Values) {
	execptionUrl := fmt.Sprintf("/execption?%s", vs.Encode())
	c.Redirect(http.StatusSeeOther, execptionUrl)
}

var (
	sessionCtxKey    = struct{}{}
	sessionErrCtxKey = struct{}{}
)

type authMiddleware struct {
	getApiClient func() *identityclient.APIClient
	unauthHandle func(c *gin.Context, uai UnAuthInfo)
}

func (m *authMiddleware) Handler(bypass ...bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		cookie := c.Request.Header.Get("cookie")

		session, uai := m.getAuthInfo(ctx, cookie)
		if session != nil {
			ctx = context.WithValue(ctx, sessionCtxKey, session)
			c.Request = c.Request.WithContext(ctx)
			c.Next()
			return
		}

		if len(bypass) > 0 && bypass[0] {
			ctx = context.WithValue(ctx, sessionErrCtxKey, &uai)
			c.Request = c.Request.WithContext(ctx)
			c.Next()
			return
		}

		// callback
		if m.unauthHandle != nil {
			m.unauthHandle(c, uai)
			c.Abort()
			return
		}

		c.AbortWithError(http.StatusUnauthorized, errors.New("not authenticated"))
	}
}

func (m *authMiddleware) getAuthInfo(ctx context.Context, cookie string) (session *identityclient.Session, uai UnAuthInfo) {
	if m.getApiClient == nil {
		panic("APIClient do not set")
	}

	sess, resp, err := m.getApiClient().SessionApi.
		ToSessionRequest(ctx).
		Cookie(cookie).
		Execute()
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusForbidden {
			// // ask to aal2
			uai.AskAal2 = true
		}

		uai.Message = err.Error()
	} else {
		session = sess
	}

	return
}

type UnAuthInfo struct {
	AskAal2 bool
	Message string
}

func AuthSession(ctx context.Context) *identityclient.Session {
	session, ok := ctx.Value(sessionCtxKey).(*identityclient.Session)
	if !ok {
		return nil
	}

	return session
}

func UnAuth(ctx context.Context) *UnAuthInfo {
	uai, ok := ctx.Value(sessionErrCtxKey).(*UnAuthInfo)
	if !ok {
		return nil
	}

	return uai
}

func (s *Server) debugPrint(title string, v any) {
	if !s.dev {
		return
	}

	buffer, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Printf("debug[%s]: unable to parse -> %v\n", title, err)
		return
	}

	fmt.Printf("debug[%s]: \n%s\n\n", title, string(buffer))
}
