package srv

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gunsluo/go-example/ory/identity-ui/swagger/identityclient"
)

type execptionMessage struct {
	Code    int    `json:"code" form:"code"`
	Message string `json:"message" form:"message"`
}

func (s *Server) execption(c *gin.Context) {
	var param execptionMessage
	if err := c.BindQuery(&param); err != nil {
		param.Code = int(identityclient.CODE__4001)
		param.Message = err.Error()
	}

	c.HTML(http.StatusOK, "execption.html", gin.H{
		"code":     param.Code,
		"message":  param.Message,
		"indexUrl": "/",
	})
}
