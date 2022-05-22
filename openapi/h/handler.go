package h

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gunsluo/go-example/openapi/v"
)

// swagger:route POST /org/add-members organization addMembersRequest
//
// Add members to the organization
//
// This endpoint tells Configurator that some members is added to the organziation.
//
// The consent challenge is appended to the consent provider's URL to which the subject's user-agent (browser) is redirected to. The consent
// provider uses that challenge to fetch information on the OAuth2 request and then tells ORY Hydra if the subject accepted
// or rejected the request.
//
// The response contains information about the created user.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     security:
//     - Bearer:
//     	 -
//
//     Responses:
//       200: addMembersResponse
//       400: jsonErrorResponse
//       500: jsonErrorResponse
func AddMembers(c *gin.Context) {
	var param v.AddMembersRequest
	if err := c.ShouldBindJSON(&param.Body); err != nil {
		v.JSONError(c, v.BadRequest, "invalid request: %v", err)
		return
	}

	param.OrganizationId = getOrganizationId(c)
	fmt.Println("--->", param, c.Request.Header.Get("Authorization"))

	var resp v.AddMembersResponse
	resp.Data.Created = []string{"mock"}
	v.JSON(c, &resp)
}

func getOrganizationId(c *gin.Context) int64 {
	v := c.Request.Header.Get("organizationId")
	id, _ := strconv.ParseInt(v, 10, 64)
	return id
}
