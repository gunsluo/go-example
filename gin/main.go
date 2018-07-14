package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.Static("/views", "./views")
	router.LoadHTMLGlob("templates/*")

	router.GET("/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	api := router.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
