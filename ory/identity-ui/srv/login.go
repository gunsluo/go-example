package srv

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/gunsluo/go-example/ory/identity-ui/swagger/identityclient"
)

func (s *Server) gotoLogin(c *gin.Context, vs url.Values) {
	redirectUrl := fmt.Sprintf("%s/self-service/login/browser?%s", s.identityEndpoint, vs.Encode())

	c.Redirect(http.StatusSeeOther, redirectUrl)
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
