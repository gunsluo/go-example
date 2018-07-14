package main

import (
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	engine.Static("/views", "./views")
	engine.LoadHTMLGlob("templates/*")

	engine.GET("/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	api := engine.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	//engine.Run(":8080") // listen and serve on 0.0.0.0:8080
	address := ":8080"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("unable to address: %s", address)
	}

	// overwrite
	run(listener, engine)
}

func run(listener net.Listener, engine *gin.Engine) (err error) {
	err = http.Serve(listener, engine)
	return
}
