package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gunsluo/go-example/openapi/h"
)

func main() {
	r := gin.Default()
	r.POST("/v2/org/add-members", h.AddMembers)
	r.Run(":9090")
}
