package ginspnego

import (
	"github.com/gin-gonic/gin"
	"github.com/jcmturner/goidentity/v6"
	"github.com/jcmturner/gokrb5/v8/credentials"
	"github.com/jcmturner/gokrb5/v8/gssapi"
	"github.com/jcmturner/gokrb5/v8/keytab"
	"github.com/jcmturner/gokrb5/v8/service"
	"github.com/jcmturner/gokrb5/v8/types"
)

// SPNEGOKRB5AuthenticateHandler is a Kerberos SPNEGO authentication Gin handler wrapper.
func SPNEGOKRB5AuthenticateHandler(kt *keytab.Keytab, settings ...func(*service.Settings)) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		w := c.Writer

		// Set up the SPNEGO GSS-API mechanism
		var spnego *SPNEGO

		h, err := types.GetHostAddress(r.RemoteAddr)
		if err == nil {
			// put in this order so that if the user provides a ClientAddress it will override the one here.
			o := append([]func(*service.Settings){service.ClientAddress(h)}, settings...)
			spnego = SPNEGOService(kt, o...)
		} else {
			spnego = SPNEGOService(kt, settings...)
			spnego.Log("%s - SPNEGO could not parse client address: %v", r.RemoteAddr, err)
		}

		// Check if there is a session manager and if there is an already established session for this client
		id, err := getSessionCredentials(spnego, r)
		if err == nil && id.Authenticated() {
			// There is an established session so bypass auth and serve
			spnego.Log("%s - SPNEGO request served under session %s", r.RemoteAddr, id.SessionID())
			c.Request = goidentity.AddToHTTPRequestContext(&id, r)
			// inner.ServeHTTP(w, goidentity.AddToHTTPRequestContext(&id, r))
			c.Next()
			return
		}

		st, err := getAuthorizationNegotiationHeaderAsSPNEGOToken(spnego, r, w)
		if st == nil || err != nil {
			// response to client and logging handled in function above so just return
			c.Abort()
			return
		}

		// Validate the context token
		authed, ctx, status := spnego.AcceptSecContext(st)
		if status.Code != gssapi.StatusComplete && status.Code != gssapi.StatusContinueNeeded {
			spnegoResponseReject(spnego, w, "%s - SPNEGO validation error: %v", r.RemoteAddr, status)
			c.Abort()
			return
		}
		if status.Code == gssapi.StatusContinueNeeded {
			spnegoNegotiateKRB5MechType(spnego, w, "%s - SPNEGO GSS-API continue needed", r.RemoteAddr)
			c.Abort()
			return
		}

		if authed {
			// Authentication successful; get user's credentials from the context
			id := ctx.Value(ctxCredentials).(*credentials.Credentials)
			// Create a new session if a session manager has been configured
			err = newSession(spnego, r, w, id)
			if err != nil {
				return
			}
			spnegoResponseAcceptCompleted(spnego, w, "%s %s@%s - SPNEGO authentication succeeded", r.RemoteAddr, id.UserName(), id.Domain())
			// Add the identity to the context and serve the inner/wrapped handler
			c.Request = goidentity.AddToHTTPRequestContext(id, r)
			// inner.ServeHTTP(w, goidentity.AddToHTTPRequestContext(id, r))
			c.Next()
			return
		}

		// If we get to here we have not authenticationed so just reject
		spnegoResponseReject(spnego, w, "%s - SPNEGO Kerberos authentication failed", r.RemoteAddr)
		c.Abort()
		return
	}
}
