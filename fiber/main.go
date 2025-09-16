package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func main() {
	app := fiber.New()

	ui := app.Group("/")
	ui.Use(func(c fiber.Ctx) error {
		path := c.Path()
		if !strings.HasPrefix(path, "/api") {
			// 非 /api 请求，交给静态文件服务处理
			if strings.HasSuffix(path, "_app.config.js") {
				// no cache
				c.Set("Cache-Control", "no-cache, no-store, must-revalidate")
				c.Set("Pragma", "no-cache")
				c.Set("Expires", "0")
			} else {
				// cache
				d := time.Hour
				c.Set("Cache-Control", fmt.Sprintf("max-age=%d", int(d.Seconds())))
				c.Set("Expires", time.Now().Add(d).Format(http.TimeFormat))
				c.Set("Vary", "Accept-Encoding")
			}

			// filepath := "/Users/luoji/gopath/src/github.com/oxidnova/novadm/apps/web-antd/dist/" + path
			// fmt.Println("-------filepath", filepath)
			// static.New(filepath)(c)
			// return nil
		}

		return c.Next()
	})

	// Serve static files from the "./public" directory
	ui.Get("/*", static.New("/Users/luoji/gopath/src/github.com/oxidnova/novadm/apps/web-antd/dist/"))

	api := app.Group("/api")
	api.Get("/login", func(c fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s", c.Params("*"))
		return c.SendString(msg) // => ✋ register
	})

	log.Fatal(app.Listen(":3000"))
}
